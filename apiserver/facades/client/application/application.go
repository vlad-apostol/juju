// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package application

import (
	"context"
	"fmt"
	"reflect"

	"github.com/juju/errors"
	"github.com/juju/names/v6"
	"github.com/juju/schema"
	"github.com/juju/version/v2"
	goyaml "gopkg.in/yaml.v2"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/common/storagecommon"
	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/apiserver/facade"
	apiservercharms "github.com/juju/juju/apiserver/internal/charms"
	"github.com/juju/juju/caas"
	k8sconstants "github.com/juju/juju/caas/kubernetes/provider/constants"
	corebase "github.com/juju/juju/core/base"
	corecharm "github.com/juju/juju/core/charm"
	"github.com/juju/juju/core/config"
	"github.com/juju/juju/core/constraints"
	corehttp "github.com/juju/juju/core/http"
	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/core/leadership"
	corelogger "github.com/juju/juju/core/logger"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/core/network/firewall"
	"github.com/juju/juju/core/objectstore"
	"github.com/juju/juju/core/permission"
	coreresource "github.com/juju/juju/core/resource"
	"github.com/juju/juju/core/secrets"
	coreunit "github.com/juju/juju/core/unit"
	jujuversion "github.com/juju/juju/core/version"
	applicationerrors "github.com/juju/juju/domain/application/errors"
	applicationservice "github.com/juju/juju/domain/application/service"
	"github.com/juju/juju/environs/bootstrap"
	environsconfig "github.com/juju/juju/environs/config"
	"github.com/juju/juju/internal/charm"
	"github.com/juju/juju/internal/charmhub"
	"github.com/juju/juju/internal/configschema"
	internalerrors "github.com/juju/juju/internal/errors"
	"github.com/juju/juju/internal/storage"
	"github.com/juju/juju/internal/tools"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
	"github.com/juju/juju/state/stateenvirons"
)

var ClassifyDetachedStorage = storagecommon.ClassifyDetachedStorage

// APIv20 provides the Application API facade for version 20.
type APIv20 struct {
	*APIBase
}

// APIv19 provides the Application API facade for version 19.
type APIv19 struct {
	*APIv20
}

// APIBase implements the shared application interface and is the concrete
// implementation of the api end point.
type APIBase struct {
	backend       Backend
	storageAccess StorageInterface
	store         objectstore.ObjectStore

	externalControllerService ExternalControllerService
	authorizer                facade.Authorizer
	check                     BlockChecker
	repoDeploy                DeployFromRepository

	model              Model
	modelInfo          model.ModelInfo
	modelConfigService ModelConfigService
	machineService     MachineService
	applicationService ApplicationService
	networkService     NetworkService
	portService        PortService
	relationService    RelationService
	resourceService    ResourceService
	stubService        StubService
	storageService     StorageService

	resources        facade.Resources
	leadershipReader leadership.Reader

	registry              storage.ProviderRegistry
	caasBroker            CaasBrokerInterface
	deployApplicationFunc DeployApplicationFunc

	logger corelogger.Logger
}

type CaasBrokerInterface interface {
	ValidateStorageClass(ctx context.Context, config map[string]interface{}) error
}

func newFacadeBase(stdCtx context.Context, ctx facade.ModelContext) (*APIBase, error) {
	m, err := ctx.State().Model()
	if err != nil {
		return nil, errors.Annotate(err, "getting model")
	}
	// modelShim wraps the AllPorts() API.
	model := &modelShim{Model: m}
	storageAccess, err := getStorageState(ctx.State())
	if err != nil {
		return nil, errors.Annotate(err, "getting state")
	}
	domainServices := ctx.DomainServices()
	blockChecker := common.NewBlockChecker(domainServices.BlockCommand())

	storageService := domainServices.Storage()

	registry, err := storageService.GetStorageRegistry(stdCtx)
	if err != nil {
		return nil, errors.Annotate(err, "getting storage registry")
	}

	var caasBroker caas.Broker
	if model.Type() == state.ModelTypeCAAS {
		caasBroker, err = stateenvirons.GetNewCAASBrokerFunc(caas.New)(model, domainServices.Cloud(), domainServices.Credential(), domainServices.Config())
		if err != nil {
			return nil, errors.Annotate(err, "getting caas client")
		}
	}
	resources := ctx.Resources()

	leadershipReader, err := ctx.LeadershipReader()
	if err != nil {
		return nil, errors.Trace(err)
	}

	state := &stateShim{
		State: ctx.State(),
	}

	charmhubHTTPClient, err := ctx.HTTPClient(corehttp.CharmhubPurpose)
	if err != nil {
		return nil, fmt.Errorf(
			"getting charm hub http client: %w",
			err,
		)
	}

	repoLogger := ctx.Logger().Child("deployfromrepo")
	modelInfo, err := domainServices.ModelInfo().GetModelInfo(stdCtx)
	if err != nil {
		return nil, fmt.Errorf("getting model info: %w", err)
	}

	applicationService := domainServices.Application()

	validatorCfg := validatorConfig{
		charmhubHTTPClient: charmhubHTTPClient,
		caasBroker:         caasBroker,
		model:              m,
		modelInfo:          modelInfo,
		modelConfigService: domainServices.Config(),
		machineService:     domainServices.Machine(),
		applicationService: applicationService,
		registry:           registry,
		state:              state,
		storageService:     storageService,
		logger:             repoLogger,
	}

	repoDeploy := NewDeployFromRepositoryAPI(
		state,
		applicationService,
		ctx.ObjectStore(),
		makeDeployFromRepositoryValidator(stdCtx, validatorCfg),
		repoLogger,
	)

	return NewAPIBase(
		state,
		Services{
			ExternalControllerService: domainServices.ExternalController(),
			NetworkService:            domainServices.Network(),
			ModelConfigService:        domainServices.Config(),
			MachineService:            domainServices.Machine(),
			ApplicationService:        applicationService,
			PortService:               domainServices.Port(),
			RelationService:           domainServices.Relation(),
			ResourceService:           domainServices.Resource(),
			StorageService:            storageService,
			StubService:               domainServices.Stub(),
		},
		storageAccess,
		ctx.Auth(),
		blockChecker,
		model,
		modelInfo,
		leadershipReader,
		repoDeploy,
		DeployApplication,
		registry,
		resources,
		caasBroker,
		ctx.ObjectStore(),
		ctx.Logger().Child("application"),
	)
}

// DeployApplicationFunc is a function that deploys an application.
type DeployApplicationFunc = func(
	context.Context,
	ApplicationDeployer,
	Model,
	model.ModelInfo,
	ApplicationService,
	objectstore.ObjectStore,
	DeployApplicationParams,
	corelogger.Logger,
) (Application, error)

// NewAPIBase returns a new application API facade.
func NewAPIBase(
	backend Backend,
	services Services,
	storageAccess StorageInterface,
	authorizer facade.Authorizer,
	blockChecker BlockChecker,
	model Model,
	modelInfo model.ModelInfo,
	leadershipReader Leadership,
	repoDeploy DeployFromRepository,
	deployApplication DeployApplicationFunc,
	registry storage.ProviderRegistry,
	resources facade.Resources,
	caasBroker CaasBrokerInterface,
	store objectstore.ObjectStore,
	logger corelogger.Logger,
) (*APIBase, error) {
	if !authorizer.AuthClient() {
		return nil, apiservererrors.ErrPerm
	}

	if err := services.Validate(); err != nil {
		return nil, errors.Trace(err)
	}

	return &APIBase{
		backend:               backend,
		storageAccess:         storageAccess,
		authorizer:            authorizer,
		repoDeploy:            repoDeploy,
		check:                 blockChecker,
		model:                 model,
		modelInfo:             modelInfo,
		leadershipReader:      leadershipReader,
		deployApplicationFunc: deployApplication,
		registry:              registry,
		resources:             resources,
		caasBroker:            caasBroker,
		store:                 store,

		applicationService:        services.ApplicationService,
		externalControllerService: services.ExternalControllerService,
		machineService:            services.MachineService,
		modelConfigService:        services.ModelConfigService,
		networkService:            services.NetworkService,
		portService:               services.PortService,
		relationService:           services.RelationService,
		resourceService:           services.ResourceService,
		storageService:            services.StorageService,
		stubService:               services.StubService,

		logger: logger,
	}, nil
}

// checkAccess checks if this API has the requested access level.
func (api *APIBase) checkAccess(ctx context.Context, access permission.Access) error {
	return api.authorizer.HasPermission(ctx, access, names.NewModelTag(api.modelInfo.UUID.String()))
}

func (api *APIBase) checkCanRead(ctx context.Context) error {
	return api.checkAccess(ctx, permission.ReadAccess)
}

func (api *APIBase) checkCanWrite(ctx context.Context) error {
	return api.checkAccess(ctx, permission.WriteAccess)
}

// Deploy fetches the charms from the charm store and deploys them
// using the specified placement directives.
func (api *APIBase) Deploy(ctx context.Context, args params.ApplicationsDeploy) (params.ErrorResults, error) {
	if err := api.checkCanWrite(ctx); err != nil {
		return params.ErrorResults{}, errors.Trace(err)
	}
	result := params.ErrorResults{
		Results: make([]params.ErrorResult, len(args.Applications)),
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return result, errors.Trace(err)
	}

	for i, arg := range args.Applications {
		if err := apiservercharms.ValidateCharmOrigin(arg.CharmOrigin); err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		// Fill in the charm origin revision from the charm url if it's absent
		if arg.CharmOrigin.Revision == nil {
			curl, err := charm.ParseURL(arg.CharmURL)
			if err != nil {
				result.Results[i].Error = apiservererrors.ServerError(err)
				continue
			}
			rev := curl.Revision
			arg.CharmOrigin.Revision = &rev
		}
		err := api.deployApplication(ctx, arg)
		if err == nil {
			// Deploy succeeded, no cleanup needed, move on to the next.
			continue
		}
		result.Results[i].Error = apiservererrors.ServerError(errors.Annotatef(err, "cannot deploy %q", arg.ApplicationName))

		api.cleanupResourcesAddedBeforeApp(ctx, arg.ApplicationName, arg.Resources)
	}
	return result, nil
}

// cleanupResourcesAddedBeforeApp deletes any resources added before the
// application. Errors will be logged but not reported to the user. These
// errors mask the real deployment failure.
func (api *APIBase) cleanupResourcesAddedBeforeApp(ctx context.Context, appName string, argResources map[string]string) {
	if len(argResources) == 0 {
		return
	}

	pendingIDs := make([]coreresource.UUID, 0, len(argResources))
	for _, resource := range argResources {
		resUUID, err := coreresource.ParseUUID(resource)
		if err != nil {
			api.logger.Warningf(ctx, "unable to parse resource UUID %q, while cleaning up pending"+
				" resources from a failed application deployment: %w", resource, err)
			continue
		}
		pendingIDs = append(pendingIDs, resUUID)
	}
	err := api.resourceService.DeleteResourcesAddedBeforeApplication(ctx, pendingIDs)
	if err != nil {
		api.logger.Errorf(ctx, "removing pending resources for %q: %w", appName, err)
	}
}

// ConfigSchema returns the config schema and defaults for an application.
func ConfigSchema() (configschema.Fields, schema.Defaults, error) {
	return trustFields, trustDefaults, nil
}

func splitApplicationAndCharmConfig(inConfig map[string]string) (
	appCfg map[string]interface{},
	charmCfg map[string]string,
	_ error,
) {

	providerSchema, _, err := ConfigSchema()
	if err != nil {
		return nil, nil, errors.Trace(err)
	}
	appConfigKeys := config.KnownConfigKeys(providerSchema)

	appConfigAttrs := make(map[string]interface{})
	charmConfig := make(map[string]string)
	for k, v := range inConfig {
		if appConfigKeys.Contains(k) {
			appConfigAttrs[k] = v
		} else {
			charmConfig[k] = v
		}
	}
	return appConfigAttrs, charmConfig, nil
}

// splitApplicationAndCharmConfigFromYAML extracts app specific settings from a charm config YAML
// and returns those app settings plus a YAML with just the charm settings left behind.
func splitApplicationAndCharmConfigFromYAML(inYaml, appName string) (
	appCfg map[string]interface{},
	outYaml string,
	_ error,
) {
	var allSettings map[string]interface{}
	if err := goyaml.Unmarshal([]byte(inYaml), &allSettings); err != nil {
		return nil, "", errors.Annotate(err, "cannot parse settings data")
	}
	settings, ok := allSettings[appName].(map[interface{}]interface{})
	if !ok {
		// Application key not present; it might be 'juju get' output.
		if _, err := charmConfigFromConfigValues(allSettings); err != nil {
			return nil, "", errors.Errorf("no settings found for %q", appName)
		}

		return nil, inYaml, nil
	}

	providerSchema, _, err := ConfigSchema()
	if err != nil {
		return nil, "", errors.Trace(err)
	}
	appConfigKeys := config.KnownConfigKeys(providerSchema)

	appConfigAttrs := make(map[string]interface{})
	for k, v := range settings {
		if key, ok := k.(string); ok && appConfigKeys.Contains(key) {
			appConfigAttrs[key] = v
			delete(settings, k)
		}
	}
	if len(settings) == 0 {
		return appConfigAttrs, "", nil
	}

	allSettings[appName] = settings
	charmConfig, err := goyaml.Marshal(allSettings)
	if err != nil {
		return nil, "", errors.Annotate(err, "cannot marshall charm settings")
	}
	return appConfigAttrs, string(charmConfig), nil
}

// caasDeployParams contains deploy configuration requiring prechecks
// specific to a caas.
type caasDeployParams struct {
	applicationName string
	attachStorage   []string
	charm           CharmMeta
	config          map[string]string
	placement       []*instance.Placement
	storage         map[string]storage.Directive
}

// precheck, checks the deploy config based on caas specific
// requirements.
func (c caasDeployParams) precheck(
	ctx context.Context,
	modelConfigService ModelConfigService,
	storageService StorageService,
	registry storage.ProviderRegistry,
	caasBroker CaasBrokerInterface,
) error {
	if len(c.attachStorage) > 0 {
		return errors.Errorf(
			"AttachStorage may not be specified for container models",
		)
	}
	if len(c.placement) > 1 {
		return errors.Errorf(
			"only 1 placement directive is supported for container models, got %d",
			len(c.placement),
		)
	}
	for _, s := range c.charm.Meta().Storage {
		if s.Type == charm.StorageBlock {
			return errors.Errorf("block storage %q is not supported for container charms", s.Name)
		}
	}
	cfg, err := modelConfigService.ModelConfig(ctx)
	if err != nil {
		return fmt.Errorf("getting model config: %w", err)
	}

	workloadStorageClass, _ := cfg.AllAttrs()[k8sconstants.WorkloadStorageKey].(string)
	for storageName, cons := range c.storage {
		if cons.Pool == "" && workloadStorageClass == "" {
			return errors.Errorf("storage pool for %q must be specified since there's no model default storage class", storageName)
		}
		sp, err := charmStorageParams(ctx, "", workloadStorageClass, cfg, cons.Pool, storageService, registry)
		if err != nil {
			return errors.Annotatef(err, "getting workload storage params for %q", c.applicationName)
		}
		if sp.Provider != string(k8sconstants.StorageProviderType) {
			poolName := cfg.AllAttrs()[k8sconstants.WorkloadStorageKey]
			return errors.Errorf(
				"the %q storage pool requires a provider type of %q, not %q", poolName, k8sconstants.StorageProviderType, sp.Provider)
		}
		if err := caasBroker.ValidateStorageClass(ctx, sp.Attributes); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

// deployApplication fetches the charm from the charm store and deploys it.
// The logic has been factored out into a common function which is called by
// both the legacy API on the client facade, as well as the new application facade.
func (api *APIBase) deployApplication(
	ctx context.Context,
	args params.ApplicationDeploy,
) error {
	curl, err := charm.ParseURL(args.CharmURL)
	if err != nil {
		return errors.Trace(err)
	}
	if curl.Revision < 0 {
		return errors.Errorf("charm url must include revision")
	}

	// This check is done early so that errors deeper in the call-stack do not
	// leave an application deployment in an unrecoverable error state.
	if err := checkMachinePlacement(api.backend, api.modelInfo.UUID, args.ApplicationName, args.Placement); err != nil {
		return errors.Trace(err)
	}

	locator, err := apiservercharms.CharmLocatorFromURL(args.CharmURL)
	if err != nil {
		return errors.Trace(err)
	}
	ch, err := api.getCharm(ctx, locator)
	if err != nil {
		return errors.Trace(err)
	}

	if err := jujuversion.CheckJujuMinVersion(ch.Meta().MinJujuVersion, jujuversion.Current); err != nil {
		return errors.Trace(err)
	}

	// Codify an implicit assumption for Deploy, that AddPendingResources
	// has been called first by the client. This validates that local charm
	// and bundle deployments by a client, have provided the needed resource
	// data, whether or not the user has made specific requests. This differs
	// from the DeployFromRepository expected code path where unknown resource
	// specific are filled in by the facade method.
	if len(ch.Meta().Resources) != len(args.Resources) {
		return errors.Errorf("not all pending resources for charm provided")
	}

	if api.modelInfo.Type == model.CAAS {
		caas := caasDeployParams{
			applicationName: args.ApplicationName,
			attachStorage:   args.AttachStorage,
			charm:           ch,
			config:          args.Config,
			placement:       args.Placement,
			storage:         args.Storage,
		}
		if err := caas.precheck(ctx, api.modelConfigService, api.storageService, api.registry, api.caasBroker); err != nil {
			return errors.Trace(err)
		}
	}

	appConfig, _, charmSettings, _, err := parseCharmSettings(ch.Config(), args.ApplicationName, args.Config, args.ConfigYAML, environsconfig.UseDefaults)
	if err != nil {
		return errors.Trace(err)
	}

	// Parse storage tags in AttachStorage.
	if len(args.AttachStorage) > 0 && args.NumUnits != 1 {
		return errors.Errorf("AttachStorage is non-empty, but NumUnits is %d", args.NumUnits)
	}
	attachStorage := make([]names.StorageTag, len(args.AttachStorage))
	for i, tagString := range args.AttachStorage {
		tag, err := names.ParseStorageTag(tagString)
		if err != nil {
			return errors.Trace(err)
		}
		attachStorage[i] = tag
	}

	bindingsWithSpaceIDs, err := api.convertSpacesToIDInBindings(ctx, args.EndpointBindings)
	if err != nil {
		return errors.Trace(err)
	}
	bindings, err := state.NewBindings(api.backend, bindingsWithSpaceIDs)
	if err != nil {
		return errors.Trace(err)
	}
	origin, err := convertCharmOrigin(args.CharmOrigin)
	if err != nil {
		return errors.Trace(err)
	}

	// TODO: replace model with model info/config services
	_, err = api.deployApplicationFunc(ctx, api.backend, api.model, api.modelInfo, api.applicationService, api.store, DeployApplicationParams{
		ApplicationName:   args.ApplicationName,
		Charm:             ch,
		CharmOrigin:       origin,
		NumUnits:          args.NumUnits,
		ApplicationConfig: appConfig,
		CharmConfig:       charmSettings,
		Constraints:       args.Constraints,
		Placement:         args.Placement,
		Storage:           args.Storage,
		Devices:           args.Devices,
		AttachStorage:     attachStorage,
		EndpointBindings:  bindings.Map(),
		Resources:         args.Resources,
		Force:             args.Force,
	}, api.logger)
	return errors.Trace(err)
}

// convertCharmOrigin converts a params CharmOrigin to a core charm
// Origin. If the input origin is nil, a core charm Origin is deduced
// from the provided data. It is used in both deploying and refreshing
// charms, including from old clients which aren't charm origin aware.
// MaybeSeries is a fallback if the origin is not provided.
func convertCharmOrigin(origin *params.CharmOrigin) (corecharm.Origin, error) {
	if origin == nil {
		return corecharm.Origin{}, errors.NotValidf("nil charm origin")
	}

	originType := origin.Type
	base, err := corebase.ParseBase(origin.Base.Name, origin.Base.Channel)
	if err != nil {
		return corecharm.Origin{}, errors.Trace(err)
	}
	platform := corecharm.Platform{
		Architecture: origin.Architecture,
		OS:           base.OS,
		Channel:      base.Channel.Track,
	}

	var track string
	if origin.Track != nil {
		track = *origin.Track
	}
	var branch string
	if origin.Branch != nil {
		branch = *origin.Branch
	}
	// We do guarantee that there will be a risk value.
	// Ignore the error, as only caused by risk as an
	// empty string.
	var channel *charm.Channel
	if ch, err := charm.MakeChannel(track, origin.Risk, branch); err == nil {
		channel = &ch
	}

	return corecharm.Origin{
		Type:     originType,
		Source:   corecharm.Source(origin.Source),
		ID:       origin.ID,
		Hash:     origin.Hash,
		Revision: origin.Revision,
		Channel:  channel,
		Platform: platform,
	}, nil
}

func validateSecretConfig(chCfg *charm.Config, cfg charm.Settings) error {
	for name, value := range cfg {
		option, ok := chCfg.Options[name]
		if !ok {
			// This should never happen.
			return errors.NotValidf("unknown option %q", name)
		}
		if option.Type == "secret" {
			uriStr, ok := value.(string)
			if !ok {
				return errors.NotValidf("secret value should be a string, got %T instead", name)
			}
			if uriStr == "" {
				return nil
			}
			_, err := secrets.ParseURI(uriStr)
			return errors.Annotatef(err, "invalid secret URI for option %q", name)
		}
	}
	return nil
}

// parseCharmSettings parses, verifies and combines the config settings for a
// charm as specified by the provided config map and config yaml payload. Any
// model-specific application settings will be automatically extracted and
// returned back as an *application.Config.
func parseCharmSettings(
	chCfg *charm.Config, appName string, cfg map[string]string, configYaml string, defaults environsconfig.Defaulting,
) (_ *config.Config, _ configschema.Fields, chOut charm.Settings, _ schema.Defaults, err error) {
	defer func() {
		if chOut != nil {
			err = validateSecretConfig(chCfg, chOut)
		}
	}()

	// Split out the app config from the charm config for any config
	// passed in as a map as opposed to YAML.
	var (
		applicationConfig map[string]interface{}
		charmConfig       map[string]string
	)
	if len(cfg) > 0 {
		if applicationConfig, charmConfig, err = splitApplicationAndCharmConfig(cfg); err != nil {
			return nil, nil, nil, nil, errors.Trace(err)
		}
	}

	// Split out the app config from the charm config for any config
	// passed in as YAML.
	var (
		charmYamlConfig string
		appSettings     = make(map[string]interface{})
	)
	if len(configYaml) != 0 {
		if appSettings, charmYamlConfig, err = splitApplicationAndCharmConfigFromYAML(configYaml, appName); err != nil {
			return nil, nil, nil, nil, errors.Trace(err)
		}
	}

	// Entries from the string-based config map always override any entries
	// provided via the YAML payload.
	for k, v := range applicationConfig {
		appSettings[k] = v
	}

	appCfgSchema, schemaDefaults, err := ConfigSchema()
	if err != nil {
		return nil, nil, nil, nil, errors.Trace(err)
	}
	getDefaults := func() schema.Defaults {
		// If defaults is UseDefaults, defaults are baked into the app config.
		if defaults == environsconfig.UseDefaults {
			return schemaDefaults
		}
		return nil
	}
	appConfig, err := config.NewConfig(appSettings, appCfgSchema, getDefaults())
	if err != nil {
		return nil, nil, nil, nil, errors.Trace(err)
	}

	// If there isn't a charm YAML, then we can just return the charmConfig as
	// the settings and no need to attempt to parse an empty yaml.
	if len(charmYamlConfig) == 0 {
		settings, err := chCfg.ParseSettingsStrings(charmConfig)
		if err != nil {
			return nil, nil, nil, nil, errors.Trace(err)
		}
		return appConfig, appCfgSchema, settings, schemaDefaults, nil
	}

	var charmSettings charm.Settings
	// Parse the charm YAML and check the yaml against the charm config.
	if charmSettings, err = chCfg.ParseSettingsYAML([]byte(charmYamlConfig), appName); err != nil {
		// Check if this is 'juju get' output and parse it as such
		jujuGetSettings, pErr := charmConfigFromYamlConfigValues(charmYamlConfig)
		if pErr != nil {
			// Not 'juju output' either; return original error
			return nil, nil, nil, nil, errors.Trace(err)
		}
		charmSettings = jujuGetSettings
	}

	// Entries from the string-based config map always override any entries
	// provided via the YAML payload.
	if len(charmConfig) != 0 {
		// Parse config in a compatible way (see function comment).
		overrideSettings, err := parseSettingsCompatible(chCfg, charmConfig)
		if err != nil {
			return nil, nil, nil, nil, errors.Trace(err)
		}
		for k, v := range overrideSettings {
			charmSettings[k] = v
		}
	}

	return appConfig, appCfgSchema, charmSettings, schemaDefaults, nil
}

type MachinePlacementBackend interface {
	Machine(string) (Machine, error)
}

// checkMachinePlacement does a non-exhaustive validation of any supplied
// placement directives.
// If the placement scope is for a machine, ensure that the machine exists.
// If the placement scope is model-uuid, replace it with the actual model uuid.
func checkMachinePlacement(backend MachinePlacementBackend, modelID model.UUID, app string, placement []*instance.Placement) error {
	for _, p := range placement {
		if p == nil {
			continue
		}
		// Substitute the placeholder with the actual model uuid.
		if p.Scope == "model-uuid" {
			p.Scope = modelID.String()
			continue
		}

		dir := p.Directive

		toProvisionedMachine := p.Scope == instance.MachineScope
		if !toProvisionedMachine && dir == "" {
			continue
		}
	}

	return nil
}

// parseSettingsCompatible parses setting strings in a way that is
// compatible with the behavior before this CL based on the issue
// http://pad.lv/1194945. Until then setting an option to an empty
// string caused it to reset to the default value. We now allow
// empty strings as actual values, but we want to preserve the API
// behavior.
func parseSettingsCompatible(charmConfig *charm.Config, settings map[string]string) (charm.Settings, error) {
	setSettings := map[string]string{}
	unsetSettings := charm.Settings{}
	// Split settings into those which set and those which unset a value.
	for name, value := range settings {
		if value == "" {
			unsetSettings[name] = nil
			continue
		}
		setSettings[name] = value
	}
	// Validate the settings.
	changes, err := charmConfig.ParseSettingsStrings(setSettings)
	if err != nil {
		return nil, errors.Trace(err)
	}
	// Validate the unsettings and merge them into the changes.
	unsetSettings, err = charmConfig.ValidateSettings(unsetSettings)
	if err != nil {
		return nil, errors.Trace(err)
	}
	for name := range unsetSettings {
		changes[name] = nil
	}
	return changes, nil
}

type setCharmParams struct {
	AppName               string
	Application           Application
	CharmOrigin           *params.CharmOrigin
	ConfigSettingsStrings map[string]string
	ConfigSettingsYAML    string
	ResourceIDs           map[string]string
	StorageDirectives     map[string]params.StorageDirectives
	EndpointBindings      map[string]string
	Force                 forceParams
}

type forceParams struct {
	ForceBase, ForceUnits, Force bool
}

func (api *APIBase) setConfig(
	ctx context.Context,
	app Application,
	settingsYAML string,
	settingsStrings map[string]string,
) error {
	locator, err := api.getCharmLocatorByApplicationName(ctx, app.Name())
	if err != nil {
		return errors.Trace(err)
	}

	charmConfig, err := api.getCharmConfig(ctx, locator)
	if err != nil {
		return errors.Trace(err)
	}

	// parseCharmSettings is passed false for useDefaults because setConfig
	// should not care about defaults.
	// If defaults are wanted, one should call unsetApplicationConfig.
	appConfig, appConfigSchema, charmSettings, defaults, err := parseCharmSettings(charmConfig, app.Name(), settingsStrings, settingsYAML, environsconfig.NoDefaults)
	if err != nil {
		return errors.Annotate(err, "parsing settings for application")
	}

	if len(charmSettings) != 0 {
		if err = app.UpdateCharmConfig(charmSettings); err != nil {
			return errors.Annotate(err, "updating charm config settings")
		}
	}
	if cfgAttrs := appConfig.Attributes(); len(cfgAttrs) > 0 {
		if err = app.UpdateApplicationConfig(cfgAttrs, nil, appConfigSchema, defaults); err != nil {
			return errors.Annotate(err, "updating application config settings")
		}
	}

	return nil
}

// SetCharm sets the charm for a given for the application.
// The v1 args use "storage-constraints" as the storage directive attr tag.
func (api *APIv19) SetCharm(ctx context.Context, argsV1 params.ApplicationSetCharmV1) error {
	args := argsV1.ApplicationSetCharmV2
	args.StorageDirectives = argsV1.StorageDirectives
	return api.APIBase.SetCharm(ctx, args)
}

// SetCharm sets the charm for a given for the application.
func (api *APIBase) SetCharm(ctx context.Context, args params.ApplicationSetCharmV2) error {
	if err := api.checkCanWrite(ctx); err != nil {
		return err
	}

	// when forced units in error, don't block
	if !args.ForceUnits {
		if err := api.check.ChangeAllowed(ctx); err != nil {
			return errors.Trace(err)
		}
	}

	if err := apiservercharms.ValidateCharmOrigin(args.CharmOrigin); err != nil {
		return err
	}

	oneApplication, err := api.backend.Application(args.ApplicationName)
	if err != nil {
		return errors.Trace(err)
	}
	bindingsWithSpaceIDs, err := api.convertSpacesToIDInBindings(ctx, args.EndpointBindings)
	if err != nil {
		return errors.Trace(err)
	}
	return api.setCharmWithAgentValidation(
		ctx,
		setCharmParams{
			AppName:               args.ApplicationName,
			Application:           oneApplication,
			CharmOrigin:           args.CharmOrigin,
			ConfigSettingsStrings: args.ConfigSettings,
			ConfigSettingsYAML:    args.ConfigSettingsYAML,
			ResourceIDs:           args.ResourceIDs,
			StorageDirectives:     args.StorageDirectives,
			EndpointBindings:      bindingsWithSpaceIDs,
			Force: forceParams{
				ForceBase:  args.ForceBase,
				ForceUnits: args.ForceUnits,
				Force:      args.Force,
			},
		},
		args.CharmURL,
	)
}

var (
	storageUpgradeMessage = `
Juju on containers does not support updating storage on a statefulset.
The new charm's metadata contains updated storage declarations.
You'll need to deploy a new charm rather than upgrading if you need this change.
`[1:]

	devicesUpgradeMessage = `
Juju on containers does not support updating node selectors (configured from charm devices).
The new charm's metadata contains updated device declarations.
You'll need to deploy a new charm rather than upgrading if you need this change.
`[1:]
)

// setCharmWithAgentValidation checks the agent versions of the application
// and unit before continuing on. These checks are important to prevent old
// code running at the same time as the new code. If you encounter the error,
// the correct and only work around is to upgrade the units to match the
// controller.
func (api *APIBase) setCharmWithAgentValidation(
	ctx context.Context,
	params setCharmParams,
	url string,
) error {
	newOrigin, err := convertCharmOrigin(params.CharmOrigin)
	if err != nil {
		return errors.Trace(err)
	}

	newLocator, err := apiservercharms.CharmLocatorFromURL(url)
	if err != nil {
		return errors.Trace(err)
	}
	newCharm, err := api.getCharm(ctx, newLocator)
	if err != nil {
		return errors.Trace(err)
	}

	if api.modelInfo.Type == model.CAAS {
		locator, err := api.getCharmLocatorByApplicationName(ctx, params.AppName)
		if err != nil {
			return errors.Trace(err)
		}
		currentMetadata, err := api.getCharmMetadata(ctx, locator)
		if err != nil {
			return errors.Trace(err)
		}

		// We need to disallow updates that k8s does not yet support,
		// eg changing the filesystem or device directives.
		// TODO(wallyworld) - support resizing of existing storage.
		var unsupportedReason string
		if !reflect.DeepEqual(currentMetadata.Storage, newCharm.Meta().Storage) {
			unsupportedReason = storageUpgradeMessage
		} else if !reflect.DeepEqual(currentMetadata.Devices, newCharm.Meta().Devices) {
			unsupportedReason = devicesUpgradeMessage
		}
		if unsupportedReason != "" {
			return errors.NotSupportedf(unsupportedReason)
		}
		origin, err := StateCharmOrigin(newOrigin)
		if err != nil {
			return errors.Trace(err)
		}
		return api.applicationSetCharm(ctx, params, newCharm, origin)
	}

	origin, err := StateCharmOrigin(newOrigin)
	if err != nil {
		return errors.Trace(err)
	}
	return api.applicationSetCharm(ctx, params, newCharm, origin)
}

// applicationSetCharm sets the charm and updated config
// for the given application.
func (api *APIBase) applicationSetCharm(
	ctx context.Context,
	params setCharmParams,
	newCharm state.CharmRefFull,
	newOrigin *state.CharmOrigin,
) error {
	appConfig, appSchema, charmSettings, appDefaults, err := parseCharmSettings(newCharm.Config(), params.AppName, params.ConfigSettingsStrings, params.ConfigSettingsYAML, environsconfig.NoDefaults)
	if err != nil {
		return errors.Annotate(err, "parsing config settings")
	}
	if err := appConfig.Validate(); err != nil {
		return errors.Annotate(err, "validating config settings")
	}
	var stateStorageConstraints map[string]state.StorageConstraints
	if len(params.StorageDirectives) > 0 {
		stateStorageConstraints = make(map[string]state.StorageConstraints)
		for name, cons := range params.StorageDirectives {
			stateCons := state.StorageConstraints{Pool: cons.Pool}
			if cons.Size != nil {
				stateCons.Size = *cons.Size
			}
			if cons.Count != nil {
				stateCons.Count = *cons.Count
			}
			stateStorageConstraints[name] = stateCons
		}
	}

	// Enforce "assumes" requirements if the feature flag is enabled.
	if err := assertCharmAssumptions(ctx, api.applicationService, newCharm.Meta().Assumes); err != nil {
		if !errors.Is(err, errors.NotSupported) || !params.Force.Force {
			return errors.Trace(err)
		}

		api.logger.Warningf(context.TODO(), "proceeding with upgrade of application %q even though the charm feature requirements could not be met as --force was specified", params.AppName)
	}

	force := params.Force
	cfg := state.SetCharmConfig{
		Charm:              newCharm,
		CharmOrigin:        newOrigin,
		ForceBase:          force.ForceBase,
		ForceUnits:         force.ForceUnits,
		Force:              force.Force,
		PendingResourceIDs: params.ResourceIDs,
		StorageConstraints: stateStorageConstraints,
		EndpointBindings:   params.EndpointBindings,
	}
	if len(charmSettings) > 0 {
		cfg.ConfigSettings = charmSettings
	}

	// Disallow downgrading from a v2 charm to a v1 charm.
	oldCharmLocator, err := api.getCharmLocatorByApplicationName(ctx, params.AppName)
	if err != nil {
		return errors.Trace(err)
	}

	oldCharm, err := api.getCharm(ctx, oldCharmLocator)
	if err != nil {
		return errors.Trace(err)
	}
	if charm.MetaFormat(oldCharm) >= charm.FormatV2 && charm.MetaFormat(newCharm) == charm.FormatV1 {
		return errors.New("cannot downgrade from v2 charm format to v1")
	}

	// TODO(dqlite) - remove SetCharm (replaced below with UpdateApplicationCharm).
	if err := params.Application.SetCharm(cfg, api.store); err != nil {
		return errors.Annotate(err, "updating charm config")
	}

	var storageDirectives map[string]storage.Directive
	if len(params.StorageDirectives) > 0 {
		storageDirectives = make(map[string]storage.Directive)
		for name, cons := range params.StorageDirectives {
			sc := storage.Directive{Pool: cons.Pool}
			if cons.Size != nil {
				sc.Size = *cons.Size
			}
			if cons.Count != nil {
				sc.Count = *cons.Count
			}
			storageDirectives[name] = sc
		}
	}
	if err := api.applicationService.UpdateApplicationCharm(ctx, params.AppName, applicationservice.UpdateCharmParams{
		Charm:   newCharm,
		Storage: storageDirectives,
	}); err != nil {
		return errors.Annotatef(err, "updating charm for application %q", params.AppName)
	}
	if attr := appConfig.Attributes(); len(attr) > 0 {
		return params.Application.UpdateApplicationConfig(attr, nil, appSchema, appDefaults)
	}
	return nil
}

// charmConfigFromYamlConfigValues will parse a yaml produced by juju get and
// generate charm.Settings from it that can then be sent to the application.
func charmConfigFromYamlConfigValues(yamlContents string) (charm.Settings, error) {
	var allSettings map[string]interface{}
	if err := goyaml.Unmarshal([]byte(yamlContents), &allSettings); err != nil {
		return nil, errors.Annotate(err, "cannot parse settings data")
	}
	return charmConfigFromConfigValues(allSettings)
}

// charmConfigFromConfigValues will parse a yaml produced by juju get and
// generate charm.Settings from it that can then be sent to the application.
func charmConfigFromConfigValues(yamlContents map[string]interface{}) (charm.Settings, error) {
	onlySettings := charm.Settings{}
	settingsMap, ok := yamlContents["settings"].(map[interface{}]interface{})
	if !ok {
		return nil, errors.New("unknown format for settings")
	}

	for setting, content := range settingsMap {
		s, ok := content.(map[interface{}]interface{})
		if !ok {
			return nil, errors.Errorf("unknown format for settings section %v", setting)
		}
		// some keys might not have a value, we don't care about those.
		v, ok := s["value"]
		if !ok {
			continue
		}
		stringSetting, ok := setting.(string)
		if !ok {
			return nil, errors.Errorf("unexpected setting key, expected string got %T", setting)
		}
		onlySettings[stringSetting] = v
	}
	return onlySettings, nil
}

// GetCharmURLOrigin returns the charm URL and charm origin the given
// application is running at present.
func (api *APIBase) GetCharmURLOrigin(ctx context.Context, args params.ApplicationGet) (params.CharmURLOriginResult, error) {
	if err := api.checkCanRead(ctx); err != nil {
		return params.CharmURLOriginResult{}, errors.Trace(err)
	}
	oneApplication, err := api.backend.Application(args.ApplicationName)
	if err != nil {
		return params.CharmURLOriginResult{Error: apiservererrors.ServerError(err)}, nil
	}
	charmURL, _ := oneApplication.CharmURL()
	result := params.CharmURLOriginResult{URL: *charmURL}
	chOrigin := oneApplication.CharmOrigin()
	if chOrigin == nil {
		result.Error = apiservererrors.ServerError(errors.NotFoundf("charm origin for %q", args.ApplicationName))
		return result, nil
	}
	if result.Origin, err = makeParamsCharmOrigin(chOrigin); err != nil {
		result.Error = apiservererrors.ServerError(errors.NotFoundf("charm origin for %q", args.ApplicationName))
		return result, nil
	}
	result.Origin.InstanceKey = charmhub.CreateInstanceKey(oneApplication.ApplicationTag(), names.NewModelTag(api.modelInfo.UUID.String()))
	return result, nil
}

func makeParamsCharmOrigin(origin *state.CharmOrigin) (params.CharmOrigin, error) {
	retOrigin := params.CharmOrigin{
		Source: origin.Source,
		ID:     origin.ID,
		Hash:   origin.Hash,
	}
	if origin.Revision != nil {
		retOrigin.Revision = origin.Revision
	}
	if origin.Channel != nil {
		retOrigin.Risk = origin.Channel.Risk
		if origin.Channel.Track != "" {
			retOrigin.Track = &origin.Channel.Track
		}
		if origin.Channel.Branch != "" {
			retOrigin.Branch = &origin.Channel.Branch
		}
	}
	if origin.Platform != nil {
		retOrigin.Architecture = origin.Platform.Architecture
		retOrigin.Base = params.Base{Name: origin.Platform.OS, Channel: origin.Platform.Channel}
	}
	return retOrigin, nil
}

// CharmRelations implements the server side of Application.CharmRelations.
func (api *APIBase) CharmRelations(ctx context.Context, p params.ApplicationCharmRelations) (params.ApplicationCharmRelationsResults, error) {
	var results params.ApplicationCharmRelationsResults
	if err := api.checkCanRead(ctx); err != nil {
		return results, errors.Trace(err)
	}

	app, err := api.backend.Application(p.ApplicationName)
	if err != nil {
		return results, errors.Trace(err)
	}
	endpoints, err := app.Endpoints()
	if err != nil {
		return results, errors.Trace(err)
	}
	results.CharmRelations = make([]string, len(endpoints))
	for i, endpoint := range endpoints {
		results.CharmRelations[i] = endpoint.Relation.Name
	}
	return results, nil
}

// Expose changes the juju-managed firewall to expose any ports that
// were also explicitly marked by units as open.
func (api *APIBase) Expose(ctx context.Context, args params.ApplicationExpose) error {
	if err := api.checkCanWrite(ctx); err != nil {
		return errors.Trace(err)
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return errors.Trace(err)
	}
	app, err := api.backend.Application(args.ApplicationName)
	if err != nil {
		return errors.Trace(err)
	}

	// Map space names to space IDs before calling SetExposed
	mappedExposeParams, err := api.mapExposedEndpointParams(ctx, args.ExposedEndpoints)
	if err != nil {
		return apiservererrors.ServerError(err)
	}

	// If an empty exposedEndpoints list is provided, all endpoints should
	// be exposed. This emulates the expose behavior of pre 2.9 controllers.
	if len(mappedExposeParams) == 0 {
		mappedExposeParams = map[string]state.ExposedEndpoint{
			"": {
				ExposeToCIDRs: []string{firewall.AllNetworksIPV4CIDR, firewall.AllNetworksIPV6CIDR},
			},
		}
	}

	if err = app.MergeExposeSettings(mappedExposeParams); err != nil {
		return apiservererrors.ServerError(err)
	}
	return nil
}

func (api *APIBase) mapExposedEndpointParams(ctx context.Context, params map[string]params.ExposedEndpoint) (map[string]state.ExposedEndpoint, error) {
	if len(params) == 0 {
		return nil, nil
	}

	var res = make(map[string]state.ExposedEndpoint, len(params))

	spaceInfos, err := api.networkService.GetAllSpaces(ctx)
	if err != nil {
		return nil, errors.Trace(err)
	}

	for endpointName, exposeDetails := range params {
		mappedParam := state.ExposedEndpoint{
			ExposeToCIDRs: exposeDetails.ExposeToCIDRs,
		}

		if len(exposeDetails.ExposeToSpaces) != 0 {
			spaceIDs := make([]string, len(exposeDetails.ExposeToSpaces))
			for i, spaceName := range exposeDetails.ExposeToSpaces {
				sp := spaceInfos.GetByName(spaceName)
				if sp == nil {
					return nil, errors.NotFoundf("space %q", spaceName)
				}

				spaceIDs[i] = sp.ID
			}
			mappedParam.ExposeToSpaceIDs = spaceIDs
		}

		res[endpointName] = mappedParam

	}

	return res, nil
}

// Unexpose changes the juju-managed firewall to unexpose any ports that
// were also explicitly marked by units as open.
func (api *APIBase) Unexpose(ctx context.Context, args params.ApplicationUnexpose) error {
	if err := api.checkCanWrite(ctx); err != nil {
		return err
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return errors.Trace(err)
	}
	app, err := api.backend.Application(args.ApplicationName)
	if err != nil {
		return err
	}

	// No endpoints specified; unexpose application
	if len(args.ExposedEndpoints) == 0 {
		return app.ClearExposed()
	}

	// Unset expose settings for the specified endpoints
	return app.UnsetExposeSettings(args.ExposedEndpoints)
}

// AddUnits adds a given number of units to an application.
func (api *APIBase) AddUnits(ctx context.Context, args params.AddApplicationUnits) (params.AddApplicationUnitsResults, error) {
	if api.modelInfo.Type == model.CAAS {
		return params.AddApplicationUnitsResults{}, errors.NotSupportedf("adding units to a container-based model")
	}

	if err := api.checkCanWrite(ctx); err != nil {
		return params.AddApplicationUnitsResults{}, errors.Trace(err)
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return params.AddApplicationUnitsResults{}, errors.Trace(err)
	}

	// TODO(wallyworld) - enable-ha is how we add new controllers at the moment
	// Remove this check before 3.0 when enable-ha is refactored.
	locator, err := api.getCharmLocatorByApplicationName(ctx, args.ApplicationName)
	if err != nil {
		return params.AddApplicationUnitsResults{}, errors.Trace(err)
	}
	charmName, err := api.getCharmName(ctx, locator)
	if err != nil {
		return params.AddApplicationUnitsResults{}, errors.Trace(err)
	} else if charmName == bootstrap.ControllerCharmName {
		return params.AddApplicationUnitsResults{}, errors.NotSupportedf("adding units to the controller application")
	}
	charm, err := api.getCharm(ctx, locator)
	if err != nil {
		return params.AddApplicationUnitsResults{}, errors.Trace(err)
	}

	units, err := api.addApplicationUnits(ctx, args, charm.Meta())
	if err != nil {
		return params.AddApplicationUnitsResults{}, errors.Trace(err)
	}
	unitNames := make([]string, len(units))
	for i, unit := range units {
		unitNames[i] = unit.UnitTag().Id()
	}
	return params.AddApplicationUnitsResults{Units: unitNames}, nil
}

// addApplicationUnits adds a given number of units to an application.
func (api *APIBase) addApplicationUnits(
	ctx context.Context, args params.AddApplicationUnits, charmMeta *charm.Meta,
) ([]Unit, error) {
	if args.NumUnits < 1 {
		return nil, errors.New("must add at least one unit")
	}

	assignUnits := true
	if api.modelInfo.Type != model.IAAS {
		// In a CAAS model, there are no machines for
		// units to be assigned to.
		assignUnits = false
		if len(args.AttachStorage) > 0 {
			return nil, errors.Errorf(
				"AttachStorage may not be specified for %s models",
				api.modelInfo.Type,
			)
		}
		if len(args.Placement) > 1 {
			return nil, errors.Errorf(
				"only 1 placement directive is supported for %s models, got %d",
				api.modelInfo.Type,
				len(args.Placement),
			)
		}
	}

	// Parse storage tags in AttachStorage.
	if len(args.AttachStorage) > 0 && args.NumUnits != 1 {
		return nil, errors.Errorf("AttachStorage is non-empty, but NumUnits is %d", args.NumUnits)
	}
	attachStorage := make([]names.StorageTag, len(args.AttachStorage))
	for i, tagString := range args.AttachStorage {
		tag, err := names.ParseStorageTag(tagString)
		if err != nil {
			return nil, errors.Trace(err)
		}
		attachStorage[i] = tag
	}
	oneApplication, err := api.backend.Application(args.ApplicationName)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return api.addUnits(
		ctx,
		oneApplication,
		args.ApplicationName,
		args.NumUnits,
		args.Placement,
		attachStorage,
		assignUnits,
		charmMeta,
	)
}

// DestroyUnit removes a given set of application units.
func (api *APIBase) DestroyUnit(ctx context.Context, args params.DestroyUnitsParams) (params.DestroyUnitResults, error) {
	if api.modelInfo.Type == model.CAAS {
		return params.DestroyUnitResults{}, errors.NotSupportedf("removing units on a non-container model")
	}
	if err := api.checkCanWrite(ctx); err != nil {
		return params.DestroyUnitResults{}, errors.Trace(err)
	}
	if err := api.check.RemoveAllowed(ctx); err != nil {
		return params.DestroyUnitResults{}, errors.Trace(err)
	}

	destroyUnit := func(arg params.DestroyUnitParams) (*params.DestroyUnitInfo, error) {
		unitTag, err := names.ParseUnitTag(arg.UnitTag)
		if err != nil {
			return nil, errors.Trace(err)
		}

		name := unitTag.Id()
		unit, err := api.backend.Unit(name)
		if errors.Is(err, errors.NotFound) {
			return nil, errors.Errorf("unit %q does not exist", name)
		} else if err != nil {
			return nil, errors.Trace(err)
		}
		if !unit.IsPrincipal() {
			return nil, errors.Errorf("unit %q is a subordinate, to remove use remove-relation. Note: this will remove all units of %q", name, unit.ApplicationName())
		}

		// TODO(wallyworld) - enable-ha is how we remove controllers at the
		// moment Remove this check before 3.0 when enable-ha is refactored.
		appName, _ := names.UnitApplication(unitTag.Id())

		locator, err := api.getCharmLocatorByApplicationName(ctx, appName)
		if err != nil {
			return nil, errors.Trace(err)
		}
		charmName, err := api.getCharmName(ctx, locator)
		if err != nil {
			return nil, errors.Trace(err)
		} else if charmName == bootstrap.ControllerCharmName {
			return nil, errors.NotSupportedf("removing units from the controller application")
		}

		var info params.DestroyUnitInfo
		unitStorage, err := storagecommon.UnitStorage(api.storageAccess, unit.UnitTag())
		if err != nil {
			return nil, errors.Trace(err)
		}

		if arg.DestroyStorage {
			for _, s := range unitStorage {
				info.DestroyedStorage = append(
					info.DestroyedStorage,
					params.Entity{Tag: s.StorageTag().String()},
				)
			}
		} else {
			info.DestroyedStorage, info.DetachedStorage, err = ClassifyDetachedStorage(
				api.storageAccess.VolumeAccess(), api.storageAccess.FilesystemAccess(), unitStorage,
			)
			if err != nil {
				return nil, errors.Trace(err)
			}
		}
		if arg.DryRun {
			return &info, nil
		}

		unitName, err := coreunit.NewName(name)
		if err != nil {
			return nil, internalerrors.Errorf("parsing unit name %q: %w", name, err)
		}
		if err := api.applicationService.DestroyUnit(ctx, unitName); err != nil {
			if !errors.Is(err, applicationerrors.UnitNotFound) {
				return nil, errors.Trace(err)
			}
		}

		// TODO(units) - remove dual write to state
		op := unit.DestroyOperation(api.store)
		op.DestroyStorage = arg.DestroyStorage
		op.Force = arg.Force
		if arg.Force {
			op.MaxWait = common.MaxWait(arg.MaxWait)
		}
		if err := api.backend.ApplyOperation(op); err != nil {
			return nil, errors.Trace(err)
		}
		if len(op.Errors) != 0 {
			api.logger.Warningf(context.TODO(), "operational errors destroying unit %v: %v", unit.Name(), op.Errors)
		}
		return &info, nil
	}
	results := make([]params.DestroyUnitResult, len(args.Units))
	for i, entity := range args.Units {
		info, err := destroyUnit(entity)
		if err != nil {
			results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		results[i].Info = info
	}
	return params.DestroyUnitResults{
		Results: results,
	}, nil
}

// DestroyApplication removes a given set of applications.
func (api *APIBase) DestroyApplication(ctx context.Context, args params.DestroyApplicationsParams) (params.DestroyApplicationResults, error) {
	if err := api.checkCanWrite(ctx); err != nil {
		return params.DestroyApplicationResults{}, err
	}
	if err := api.check.RemoveAllowed(ctx); err != nil {
		return params.DestroyApplicationResults{}, errors.Trace(err)
	}
	destroyApp := func(arg params.DestroyApplicationParams) (*params.DestroyApplicationInfo, error) {
		tag, err := names.ParseApplicationTag(arg.ApplicationTag)
		if err != nil {
			return nil, err
		}

		locator, err := api.getCharmLocatorByApplicationName(ctx, tag.Id())
		if err != nil {
			return nil, errors.Trace(err)
		}

		name, err := api.getCharmName(ctx, locator)
		if err != nil {
			return nil, errors.Trace(err)
		} else if name == bootstrap.ControllerCharmName {
			return nil, errors.NotSupportedf("removing the controller application")
		}

		var info params.DestroyApplicationInfo
		app, err := api.backend.Application(tag.Id())
		if err != nil {
			return nil, err
		}
		units, err := app.AllUnits()
		if err != nil {
			return nil, err
		}
		storageSeen := names.NewSet()
		for _, unit := range units {
			info.DestroyedUnits = append(
				info.DestroyedUnits,
				params.Entity{Tag: unit.UnitTag().String()},
			)
			unitStorage, err := storagecommon.UnitStorage(api.storageAccess, unit.UnitTag())
			if err != nil {
				return nil, err
			}

			// Filter out storage we've already seen. Shared
			// storage may be attached to multiple units.
			var unseen []state.StorageInstance
			for _, stor := range unitStorage {
				storageTag := stor.StorageTag()
				if storageSeen.Contains(storageTag) {
					continue
				}
				storageSeen.Add(storageTag)
				unseen = append(unseen, stor)
			}
			unitStorage = unseen

			if arg.DestroyStorage {
				for _, s := range unitStorage {
					info.DestroyedStorage = append(
						info.DestroyedStorage,
						params.Entity{Tag: s.StorageTag().String()},
					)
				}
			} else {
				destroyed, detached, err := ClassifyDetachedStorage(
					api.storageAccess.VolumeAccess(), api.storageAccess.FilesystemAccess(), unitStorage,
				)
				if err != nil {
					return nil, err
				}
				info.DestroyedStorage = append(info.DestroyedStorage, destroyed...)
				info.DetachedStorage = append(info.DetachedStorage, detached...)
			}
		}

		if arg.DryRun {
			return &info, nil
		}

		// Minimally initiate destroy in dqlite.
		// It's sufficient for now just to advance the life to dying.
		err = api.applicationService.DestroyApplication(ctx, tag.Id())
		if err != nil && !errors.Is(err, applicationerrors.ApplicationNotFound) {
			return nil, errors.Annotatef(err, "destroying application %q", tag.Id())
		}

		op := app.DestroyOperation(api.store)
		op.DestroyStorage = arg.DestroyStorage
		op.Force = arg.Force
		if arg.Force {
			op.MaxWait = common.MaxWait(arg.MaxWait)
		}
		if err := api.backend.ApplyOperation(op); err != nil {
			return nil, err
		}
		if len(op.Errors) != 0 {
			api.logger.Warningf(context.TODO(), "operational errors destroying application %v: %v", tag.Id(), op.Errors)
		}

		// TODO(units) - remove when destroy is fully implemented.
		if op.Removed {
			err = api.applicationService.DeleteApplication(ctx, tag.Id())
		}
		return &info, err
	}
	results := make([]params.DestroyApplicationResult, len(args.Applications))
	for i, arg := range args.Applications {
		info, err := destroyApp(arg)
		if err != nil {
			results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		results[i].Info = info
	}
	return params.DestroyApplicationResults{
		Results: results,
	}, nil
}

// DestroyConsumedApplications removes a given set of consumed (remote) applications.
func (api *APIBase) DestroyConsumedApplications(ctx context.Context, args params.DestroyConsumedApplicationsParams) (params.ErrorResults, error) {
	if err := api.checkCanWrite(ctx); err != nil {
		return params.ErrorResults{}, err
	}
	if err := api.check.RemoveAllowed(ctx); err != nil {
		return params.ErrorResults{}, errors.Trace(err)
	}
	results := make([]params.ErrorResult, len(args.Applications))
	for i, arg := range args.Applications {
		appTag, err := names.ParseApplicationTag(arg.ApplicationTag)
		if err != nil {
			results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		app, err := api.backend.RemoteApplication(appTag.Id())
		if err != nil {
			results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		force := false
		if arg.Force != nil {
			force = *arg.Force
		}
		op := app.DestroyOperation(force)
		if force {
			op.MaxWait = common.MaxWait(arg.MaxWait)
		}
		err = api.backend.ApplyOperation(op)
		if len(op.Errors) > 0 {
			api.logger.Warningf(context.TODO(), "operational error encountered destroying consumed application %v: %v", appTag.Id(), op.Errors)
		}
		if err != nil {
			results[i].Error = apiservererrors.ServerError(err)
			continue
		}
	}
	return params.ErrorResults{
		Results: results,
	}, nil
}

// ScaleApplications scales the specified application to the requested number of units.
func (api *APIBase) ScaleApplications(ctx context.Context, args params.ScaleApplicationsParams) (params.ScaleApplicationResults, error) {
	if api.modelInfo.Type != model.CAAS {
		return params.ScaleApplicationResults{}, errors.NotSupportedf("scaling applications on a non-container model")
	}
	if err := api.checkCanWrite(ctx); err != nil {
		return params.ScaleApplicationResults{}, errors.Trace(err)
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return params.ScaleApplicationResults{}, errors.Trace(err)
	}
	scaleApplication := func(arg params.ScaleApplicationParams) (*params.ScaleApplicationInfo, error) {
		if arg.Scale < 0 && arg.ScaleChange == 0 {
			return nil, errors.NotValidf("scale < 0")
		} else if arg.Scale != 0 && arg.ScaleChange != 0 {
			return nil, errors.NotValidf("requesting both scale and scale-change")
		}

		appTag, err := names.ParseApplicationTag(arg.ApplicationTag)
		if err != nil {
			return nil, errors.Trace(err)
		}
		name := appTag.Id()

		var info params.ScaleApplicationInfo
		if arg.ScaleChange != 0 {
			newScale, err := api.applicationService.ChangeApplicationScale(ctx, name, arg.ScaleChange)
			if err != nil {
				return nil, errors.Trace(err)
			}
			info.Scale = newScale
		} else {
			if err := api.applicationService.SetApplicationScale(ctx, name, arg.Scale); err != nil {
				return nil, errors.Trace(err)
			}
			info.Scale = arg.Scale
		}
		return &info, nil
	}
	results := make([]params.ScaleApplicationResult, len(args.Applications))
	for i, entity := range args.Applications {
		info, err := scaleApplication(entity)
		if err != nil {
			results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		results[i].Info = info
	}
	return params.ScaleApplicationResults{
		Results: results,
	}, nil
}

// GetConstraints returns the constraints for a given application.
func (api *APIBase) GetConstraints(ctx context.Context, args params.Entities) (params.ApplicationGetConstraintsResults, error) {
	if err := api.checkCanRead(ctx); err != nil {
		return params.ApplicationGetConstraintsResults{}, errors.Trace(err)
	}
	results := params.ApplicationGetConstraintsResults{
		Results: make([]params.ApplicationConstraint, len(args.Entities)),
	}
	for i, arg := range args.Entities {
		cons, err := api.getConstraints(ctx, arg.Tag)
		results.Results[i].Constraints = cons
		results.Results[i].Error = apiservererrors.ServerError(err)
	}
	return results, nil
}

func (api *APIBase) getConstraints(ctx context.Context, entity string) (constraints.Value, error) {
	tag, err := names.ParseTag(entity)
	if err != nil {
		return constraints.Value{}, err
	}
	switch kind := tag.Kind(); kind {
	case names.ApplicationTagKind:
		appID, err := api.applicationService.GetApplicationIDByName(ctx, tag.Id())
		if errors.Is(err, applicationerrors.ApplicationNotFound) {
			return constraints.Value{}, errors.NotFoundf("application %s", tag.Id())
		} else if err != nil {
			return constraints.Value{}, errors.Trace(err)
		}
		cons, err := api.applicationService.GetApplicationConstraints(ctx, appID)
		if errors.Is(err, applicationerrors.ApplicationNotFound) {
			return constraints.Value{}, errors.NotFoundf("application %s", tag.Id())
		} else if err != nil {
			return constraints.Value{}, errors.Trace(err)
		}
		return cons, nil
	default:
		return constraints.Value{}, errors.Errorf("unexpected tag type, expected application, got %s", kind)
	}
}

// SetConstraints sets the constraints for a given application.
func (api *APIBase) SetConstraints(ctx context.Context, args params.SetConstraints) error {
	if err := api.checkCanWrite(ctx); err != nil {
		return err
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return errors.Trace(err)
	}

	appID, err := api.applicationService.GetApplicationIDByName(ctx, args.ApplicationName)
	if errors.Is(err, applicationerrors.ApplicationNotFound) {
		return errors.NotFoundf("application %s", args.ApplicationName)
	} else if err != nil {
		return errors.Trace(err)
	}
	err = api.applicationService.SetApplicationConstraints(ctx, appID, args.Constraints)
	if errors.Is(err, applicationerrors.ApplicationNotFound) {
		return errors.NotFoundf("application %s", args.ApplicationName)
	} else if err != nil {
		return errors.Trace(err)
	}

	// TODO(nvinuesa): Remove the double-write to mongodb once machines
	// are fully migrated to dqlite domain. We need the application
	// constraints to be available for machines, which still read from
	// mongodb.
	app, err := api.backend.Application(args.ApplicationName)
	if err != nil {
		return err
	}
	return app.SetConstraints(args.Constraints)
}

// AddRelation adds a relation between the specified endpoints and returns the relation info.
func (api *APIBase) AddRelation(ctx context.Context, args params.AddRelation) (_ params.AddRelationResults, err error) {
	if err := api.checkCanWrite(ctx); err != nil {
		return params.AddRelationResults{}, internalerrors.Capture(err)
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return params.AddRelationResults{}, internalerrors.Capture(err)
	}

	if len(args.ViaCIDRs) > 0 {
		// Integration via subnets is only for cross model relations.
		return params.AddRelationResults{}, internalerrors.Errorf("cross model relations are disabled until "+
			"backend functionality is moved to domain: %w", errors.NotImplemented)
	}

	if len(args.Endpoints) != 2 {
		return params.AddRelationResults{}, errors.BadRequestf("a relation should have exactly two endpoints")
	}
	ep1, ep2, err := api.relationService.AddRelation(
		ctx, args.Endpoints[0], args.Endpoints[1],
	)
	if err != nil {
		return params.AddRelationResults{}, internalerrors.Errorf(
			"adding relation between endpoints %q and %q: %w",
			args.Endpoints[0], args.Endpoints[1], err,
		)
	}
	return params.AddRelationResults{Endpoints: map[string]params.CharmRelation{
		ep1.ApplicationID.String(): encodeRelation(ep1.Relation),
		ep2.ApplicationID.String(): encodeRelation(ep2.Relation),
	}}, nil
}

// encodeRelation encodes a relation for sending over the wire.
func encodeRelation(rel charm.Relation) params.CharmRelation {
	return params.CharmRelation{
		Name:      rel.Name,
		Role:      string(rel.Role),
		Interface: rel.Interface,
		Optional:  rel.Optional,
		Limit:     rel.Limit,
		Scope:     string(rel.Scope),
	}
}

// DestroyRelation removes the relation between the
// specified endpoints or an id.
func (api *APIBase) DestroyRelation(ctx context.Context, args params.DestroyRelation) (err error) {
	if err := api.checkCanWrite(ctx); err != nil {
		return err
	}
	if err := api.check.RemoveAllowed(ctx); err != nil {
		return internalerrors.Capture(err)
	}

	return internalerrors.Errorf("destroying relations is not yet supported%w", errors.NotImplemented)
}

// SetRelationsSuspended sets the suspended status of the specified relations.
func (api *APIBase) SetRelationsSuspended(ctx context.Context, args params.RelationSuspendedArgs) (params.ErrorResults, error) {
	// Suspending relation is only available for Cross Model Relations
	return params.ErrorResults{}, internalerrors.Errorf("cross model relations are disabled until "+
		"backend functionality is moved to domain: %w", errors.NotImplemented)
}

// Consume adds remote applications to the model without creating any
// relations.
func (api *APIBase) Consume(ctx context.Context, args params.ConsumeApplicationArgsV5) (params.ErrorResults, error) {
	return params.ErrorResults{}, internalerrors.Errorf("cross model relations are disabled until "+
		"backend functionality is moved to domain: %w", errors.NotImplemented)
}

// Get returns the charm configuration for an application.
func (api *APIBase) Get(ctx context.Context, args params.ApplicationGet) (params.ApplicationGetResults, error) {
	if err := api.checkCanRead(ctx); err != nil {
		return params.ApplicationGetResults{}, err
	}

	return api.getConfig(ctx, args, describe)
}

// CharmConfig returns charm config for the input list of applications.
func (api *APIBase) CharmConfig(ctx context.Context, args params.ApplicationGetArgs) (params.ApplicationGetConfigResults, error) {
	if err := api.checkCanRead(ctx); err != nil {
		return params.ApplicationGetConfigResults{}, err
	}
	results := params.ApplicationGetConfigResults{
		Results: make([]params.ConfigResult, len(args.Args)),
	}
	for i, arg := range args.Args {
		config, err := api.getMergedAppAndCharmConfig(ctx, arg.ApplicationName)
		results.Results[i].Config = config
		results.Results[i].Error = apiservererrors.ServerError(err)
	}
	return results, nil
}

// GetConfig returns the charm config for each of the input applications.
func (api *APIBase) GetConfig(ctx context.Context, args params.Entities) (params.ApplicationGetConfigResults, error) {
	if err := api.checkCanRead(ctx); err != nil {
		return params.ApplicationGetConfigResults{}, err
	}
	results := params.ApplicationGetConfigResults{
		Results: make([]params.ConfigResult, len(args.Entities)),
	}
	for i, arg := range args.Entities {
		tag, err := names.ParseTag(arg.Tag)
		if err != nil {
			results.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		if tag.Kind() != names.ApplicationTagKind {
			results.Results[i].Error = apiservererrors.ServerError(
				errors.Errorf("unexpected tag type, expected application, got %s", tag.Kind()))
			continue
		}

		// Always deal with the master branch version of config.
		config, err := api.getMergedAppAndCharmConfig(ctx, tag.Id())
		results.Results[i].Config = config
		results.Results[i].Error = apiservererrors.ServerError(err)
	}
	return results, nil
}

// SetConfigs implements the server side of Application.SetConfig.  Both
// application and charm config are set. It does not unset values in
// Config map that are set to an empty string. Unset should be used for that.
func (api *APIBase) SetConfigs(ctx context.Context, args params.ConfigSetArgs) (params.ErrorResults, error) {
	var result params.ErrorResults
	if err := api.checkCanWrite(ctx); err != nil {
		return result, errors.Trace(err)
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return result, errors.Trace(err)
	}
	result.Results = make([]params.ErrorResult, len(args.Args))
	for i, arg := range args.Args {
		app, err := api.backend.Application(arg.ApplicationName)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		err = api.setConfig(ctx, app, arg.ConfigYAML, arg.Config)
		result.Results[i].Error = apiservererrors.ServerError(err)
	}
	return result, nil
}

// UnsetApplicationsConfig implements the server side of Application.UnsetApplicationsConfig.
func (api *APIBase) UnsetApplicationsConfig(ctx context.Context, args params.ApplicationConfigUnsetArgs) (params.ErrorResults, error) {
	var result params.ErrorResults
	if err := api.checkCanWrite(ctx); err != nil {
		return result, errors.Trace(err)
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return result, errors.Trace(err)
	}
	result.Results = make([]params.ErrorResult, len(args.Args))
	for i, arg := range args.Args {
		err := api.unsetApplicationConfig(arg)
		result.Results[i].Error = apiservererrors.ServerError(err)
	}
	return result, nil
}

func (api *APIBase) unsetApplicationConfig(arg params.ApplicationUnset) error {
	app, err := api.backend.Application(arg.ApplicationName)
	if err != nil {
		return errors.Trace(err)
	}

	configSchema, defaults, err := ConfigSchema()
	if err != nil {
		return errors.Trace(err)
	}
	appConfigFields := config.KnownConfigKeys(configSchema)

	var appConfigKeys []string
	charmSettings := make(charm.Settings)
	for _, name := range arg.Options {
		if appConfigFields.Contains(name) {
			appConfigKeys = append(appConfigKeys, name)
		} else {
			charmSettings[name] = nil
		}
	}

	if len(appConfigKeys) > 0 {
		if err := app.UpdateApplicationConfig(nil, appConfigKeys, configSchema, defaults); err != nil {
			return errors.Annotate(err, "updating application config values")
		}
	}

	if len(charmSettings) > 0 {
		if err := app.UpdateCharmConfig(charmSettings); err != nil {
			return errors.Annotate(err, "updating application charm settings")
		}
	}
	return nil
}

// ResolveUnitErrors marks errors on the specified units as resolved.
func (api *APIBase) ResolveUnitErrors(ctx context.Context, p params.UnitsResolved) (params.ErrorResults, error) {
	var result params.ErrorResults
	if err := api.checkCanWrite(ctx); err != nil {
		return result, errors.Trace(err)
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return result, errors.Trace(err)
	}

	if p.All {
		unitsWithErrors, err := api.backend.UnitsInError()
		if err != nil {
			return params.ErrorResults{}, errors.Trace(err)
		}
		for _, u := range unitsWithErrors {
			if err := u.Resolve(p.Retry); err != nil {
				return params.ErrorResults{}, errors.Annotatef(err, "resolve error for unit %q", u.UnitTag().Id())
			}
		}
	}

	result.Results = make([]params.ErrorResult, len(p.Tags.Entities))
	for i, entity := range p.Tags.Entities {
		tag, err := names.ParseUnitTag(entity.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		unit, err := api.backend.Unit(tag.Id())
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		err = unit.Resolve(p.Retry)
		result.Results[i].Error = apiservererrors.ServerError(err)
	}
	return result, nil
}

// ApplicationsInfo returns applications information.
func (api *APIBase) ApplicationsInfo(ctx context.Context, in params.Entities) (params.ApplicationInfoResults, error) {
	var result params.ApplicationInfoResults
	if err := api.checkCanRead(ctx); err != nil {
		return result, errors.Trace(err)
	}

	// Get all the space infos before iterating over the application infos.
	allSpaceInfosLookup, err := api.networkService.GetAllSpaces(ctx)
	if err != nil {
		return result, apiservererrors.ServerError(err)
	}

	out := make([]params.ApplicationInfoResult, len(in.Entities))
	for i, one := range in.Entities {
		tag, err := names.ParseApplicationTag(one.Tag)
		if err != nil {
			out[i].Error = apiservererrors.ServerError(err)
			continue
		}
		appLife, err := api.applicationService.GetApplicationLife(ctx, tag.Name)
		if errors.Is(err, applicationerrors.ApplicationNotFound) {
			err = errors.NotFoundf("application %q", tag.Name)
		}
		if err != nil {
			out[i].Error = apiservererrors.ServerError(err)
			continue
		}

		app, err := api.backend.Application(tag.Name)
		if err != nil {
			out[i].Error = apiservererrors.ServerError(err)
			continue
		}

		details, err := api.getConfig(ctx, params.ApplicationGet{ApplicationName: tag.Name}, describe)
		if err != nil {
			out[i].Error = apiservererrors.ServerError(err)
			continue
		}

		bindings, err := app.EndpointBindings()
		if err != nil {
			out[i].Error = apiservererrors.ServerError(err)
			continue
		}

		bindingsMap, err := bindings.MapWithSpaceNames(allSpaceInfosLookup)
		if err != nil {
			out[i].Error = apiservererrors.ServerError(err)
			continue
		}

		exposedEndpoints, err := api.mapExposedEndpointsFromState(ctx, app.ExposedEndpoints())
		if err != nil {
			out[i].Error = apiservererrors.ServerError(err)
			continue
		}

		var channel string
		origin := app.CharmOrigin()
		if origin != nil && origin.Channel != nil {
			ch := origin.Channel
			channel = charm.MakePermissiveChannel(ch.Track, ch.Risk, ch.Branch).String()
		} else {
			channel = details.Channel
		}

		out[i].Result = &params.ApplicationResult{
			Tag:              tag.String(),
			Charm:            details.Charm,
			Base:             details.Base,
			Channel:          channel,
			Constraints:      details.Constraints,
			Principal:        app.IsPrincipal(),
			Exposed:          app.IsExposed(),
			Remote:           app.IsRemote(),
			Life:             string(appLife),
			EndpointBindings: bindingsMap,
			ExposedEndpoints: exposedEndpoints,
		}
	}
	return params.ApplicationInfoResults{
		Results: out,
	}, nil
}

func (api *APIBase) mapExposedEndpointsFromState(ctx context.Context, exposedEndpoints map[string]state.ExposedEndpoint) (map[string]params.ExposedEndpoint, error) {
	if len(exposedEndpoints) == 0 {
		return nil, nil
	}

	var (
		err error
		res = make(map[string]params.ExposedEndpoint, len(exposedEndpoints))
	)

	spaceInfos, err := api.networkService.GetAllSpaces(ctx)
	if err != nil {
		return nil, err
	}

	for endpointName, exposeDetails := range exposedEndpoints {
		mappedParam := params.ExposedEndpoint{
			ExposeToCIDRs: exposeDetails.ExposeToCIDRs,
		}

		if len(exposeDetails.ExposeToSpaceIDs) != 0 {

			spaceNames := make([]string, len(exposeDetails.ExposeToSpaceIDs))
			for i, spaceID := range exposeDetails.ExposeToSpaceIDs {
				sp := spaceInfos.GetByID(spaceID)
				if sp == nil {
					return nil, errors.NotFoundf("space with ID %q", spaceID)
				}

				spaceNames[i] = string(sp.Name)
			}
			mappedParam.ExposeToSpaces = spaceNames
		}

		res[endpointName] = mappedParam
	}

	return res, nil
}

// MergeBindings merges operator-defined bindings with the current bindings for
// one or more applications.
func (api *APIBase) MergeBindings(ctx context.Context, in params.ApplicationMergeBindingsArgs) (params.ErrorResults, error) {
	if err := api.checkCanWrite(ctx); err != nil {
		return params.ErrorResults{}, err
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return params.ErrorResults{}, errors.Trace(err)
	}

	res := make([]params.ErrorResult, len(in.Args))
	for i, arg := range in.Args {
		tag, err := names.ParseApplicationTag(arg.ApplicationTag)
		if err != nil {
			res[i].Error = apiservererrors.ServerError(err)
			continue
		}
		app, err := api.backend.Application(tag.Name)
		if err != nil {
			res[i].Error = apiservererrors.ServerError(err)
			continue
		}

		bindingsWithSpaceIDs, err := api.convertSpacesToIDInBindings(ctx, arg.Bindings)
		if err != nil {
			res[i].Error = apiservererrors.ServerError(err)
			continue
		}
		bindings, err := state.NewBindings(api.backend, bindingsWithSpaceIDs)
		if err != nil {
			res[i].Error = apiservererrors.ServerError(err)
			continue
		}

		if err := app.MergeBindings(bindings, arg.Force); err != nil {
			res[i].Error = apiservererrors.ServerError(err)
		}
	}
	return params.ErrorResults{Results: res}, nil
}

// convertSpacesToIDInBindings takes the input bindings (which contain space
// names) and converts them to spaceIDs.
// TODO(nvinuesa): this method should not be needed once we migrate endpoint
// bindings to dqlite.
func (api *APIBase) convertSpacesToIDInBindings(ctx context.Context, bindings map[string]string) (map[string]string, error) {
	if bindings == nil {
		return nil, nil
	}
	newMap := make(map[string]string)
	for endpoint, spaceName := range bindings {
		space, err := api.networkService.SpaceByName(ctx, spaceName)
		if errors.Is(err, errors.NotFound) {
			return nil, errors.Annotatef(err, "space with name %q not found for endpoint %q", spaceName, endpoint)
		}
		if err != nil {
			return nil, err
		}
		newMap[endpoint] = space.ID
	}

	return newMap, nil
}

// AgentTools is a point of use agent tools requester.
type AgentTools interface {
	AgentTools() (*tools.Tools, error)
}

// AgentVersioner is a point of use agent version object.
type AgentVersioner interface {
	AgentVersion() (version.Number, error)
}

var (
	// ErrInvalidAgentVersions is a sentinal error for when we can no longer
	// upgrade juju using 2.5.x agents with 2.6 or greater controllers.
	ErrInvalidAgentVersions = errors.Errorf(
		"Unable to upgrade LXDProfile charms with the current model version. " +
			"Please run juju upgrade-model to upgrade the current model to match your controller.")
)

// UnitsInfo returns unit information for the given entities (units or
// applications).
func (api *APIBase) UnitsInfo(ctx context.Context, in params.Entities) (params.UnitInfoResults, error) {
	if err := api.checkCanRead(ctx); err != nil {
		return params.UnitInfoResults{}, err
	}

	var results []params.UnitInfoResult
	leaders, err := api.leadershipReader.Leaders()
	if err != nil {
		return params.UnitInfoResults{}, errors.Trace(err)
	}
	for _, one := range in.Entities {
		units, err := api.unitsFromTag(one.Tag)
		if err != nil {
			results = append(results, params.UnitInfoResult{Error: apiservererrors.ServerError(err)})
			continue
		}
		for _, unit := range units {
			result, err := api.unitResultForUnit(ctx, unit)
			if err != nil {
				results = append(results, params.UnitInfoResult{Error: apiservererrors.ServerError(err)})
				continue
			}
			if leader := leaders[unit.ApplicationName()]; leader == unit.Name() {
				result.Leader = true
			}
			results = append(results, params.UnitInfoResult{Result: result})
		}
	}
	return params.UnitInfoResults{
		Results: results,
	}, nil
}

// Returns the units referred to by the tag argument.  If the tag refers to a
// unit, a slice with a single unit is returned.  If the tag refers to an
// application, all the units in the application are returned.
func (api *APIBase) unitsFromTag(tag string) ([]Unit, error) {
	unitTag, err := names.ParseUnitTag(tag)
	if err == nil {
		unit, err := api.backend.Unit(unitTag.Id())
		if err != nil {
			return nil, err
		}
		return []Unit{unit}, nil
	}
	appTag, err := names.ParseApplicationTag(tag)
	if err == nil {
		app, err := api.backend.Application(appTag.Id())
		if err != nil {
			return nil, err
		}
		return app.AllUnits()
	}
	return nil, fmt.Errorf("tag %q is neither unit nor application tag", tag)
}

// Builds a *params.UnitResult describing the unit argument.
func (api *APIBase) unitResultForUnit(ctx context.Context, unit Unit) (*params.UnitResult, error) {
	app, err := api.backend.Application(unit.ApplicationName())
	if err != nil {
		return nil, err
	}
	curl, _ := app.CharmURL()
	if curl == nil {
		return nil, errors.NotValidf("application charm url")
	}
	machineId, _ := unit.AssignedMachineId()
	workloadVersion, err := unit.WorkloadVersion()
	if err != nil {
		return nil, err
	}

	unitName, err := coreunit.NewName(unit.Name())
	if err != nil {
		return nil, err
	}
	unitUUID, err := api.applicationService.GetUnitUUID(ctx, unitName)
	if errors.Is(err, applicationerrors.UnitNotFound) {
		err = errors.NotFoundf("unit %s", unitName)
	}
	if err != nil {
		return nil, err
	}
	unitLife, err := api.applicationService.GetUnitLife(ctx, unitName)
	if err != nil {
		return nil, err
	}

	result := &params.UnitResult{
		Tag:             unit.Tag().String(),
		WorkloadVersion: workloadVersion,
		Machine:         machineId,
		Charm:           *curl,
		Life:            string(unitLife),
	}
	if machineId != "" {
		machine, err := api.backend.Machine(machineId)
		if err != nil {
			return nil, err
		}
		publicAddress, err := machine.PublicAddress()
		if err == nil {
			result.PublicAddress = publicAddress.Value
		}
		// NOTE(achilleasa): this call completely ignores
		// subnets and lumps all port ranges together in a
		// single group. This works fine for pre 2.9 agents
		// as ports where always opened across all subnets.
		openPorts, err := api.openPortsOnUnit(ctx, unitUUID)
		if err != nil {
			return nil, err
		}
		result.OpenedPorts = openPorts
	}
	container, err := unit.ContainerInfo()
	if err != nil && !errors.Is(err, errors.NotFound) {
		return nil, err
	}
	if err == nil {
		if addr := container.Address(); addr != nil {
			result.Address = addr.Value
		}
		result.ProviderId = container.ProviderId()
		if len(result.OpenedPorts) == 0 {
			result.OpenedPorts = container.Ports()
		}
	}
	result.RelationData, err = api.relationData(ctx, app)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// openPortsOnMachineForUnit returns the unique set of opened ports for the
// specified unit and machine arguments without distinguishing between port
// ranges across subnets. This method is provided for backwards compatibility
// with pre 2.9 agents which assume open-ports apply to all subnets.
func (api *APIBase) openPortsOnUnit(ctx context.Context, unitUUID coreunit.UUID) ([]string, error) {
	var result []string

	groupedPortRanges, err := api.portService.GetUnitOpenedPorts(ctx, unitUUID)
	if err != nil {
		return nil, internalerrors.Errorf("getting opened ports for unit %q: %w", unitUUID, err)
	}
	for _, portRange := range groupedPortRanges.UniquePortRanges() {
		result = append(result, portRange.String())
	}
	return result, nil
}

func (api *APIBase) relationData(ctx context.Context, app Application) ([]params.EndpointRelationData, error) {
	appID, err := api.applicationService.GetApplicationIDByName(ctx, app.Name())
	if err != nil {
		return nil, internalerrors.Errorf("getting application id for %q: %v", app.Name(), err)
	}
	endpointsData, err := api.relationService.ApplicationRelationsInfo(ctx, appID)
	if err != nil {
		return nil, internalerrors.Capture(err)
	}
	var result []params.EndpointRelationData
	for _, endpointData := range endpointsData {
		unitRelationData := make(map[string]params.RelationData)
		for k, v := range endpointData.UnitRelationData {
			unitRelationData[k] = params.RelationData{
				InScope:  v.InScope,
				UnitData: v.UnitData,
			}
		}
		result = append(result, params.EndpointRelationData{
			RelationId:       endpointData.RelationID,
			Endpoint:         endpointData.Endpoint,
			CrossModel:       false,
			RelatedEndpoint:  endpointData.RelatedEndpoint,
			ApplicationData:  endpointData.ApplicationData,
			UnitRelationData: unitRelationData,
		})
	}
	return result, nil
}

// Leader returns the unit name of the leader for the given application.
func (api *APIBase) Leader(ctx context.Context, entity params.Entity) (params.StringResult, error) {
	if err := api.checkCanRead(ctx); err != nil {
		return params.StringResult{}, errors.Trace(err)
	}

	result := params.StringResult{}
	application, err := names.ParseApplicationTag(entity.Tag)
	if err != nil {
		return result, err
	}
	leaders, err := api.leadershipReader.Leaders()
	if err != nil {
		return result, errors.Annotate(err, "querying leaders")
	}
	var ok bool
	result.Result, ok = leaders[application.Name]
	if !ok || result.Result == "" {
		result.Error = apiservererrors.ServerError(errors.NotFoundf("leader for %s", entity.Tag))
	}
	return result, nil
}

// DeployFromRepository is a one-stop deployment method for repository
// charms. Only a charm name is required to deploy. If argument validation
// fails, a list of all errors found in validation will be returned. If a
// local resource is provided, details required for uploading the validated
// resource will be returned.
func (api *APIBase) DeployFromRepository(ctx context.Context, args params.DeployFromRepositoryArgs) (params.DeployFromRepositoryResults, error) {
	if err := api.checkCanWrite(ctx); err != nil {
		return params.DeployFromRepositoryResults{}, errors.Trace(err)
	}
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return params.DeployFromRepositoryResults{}, errors.Trace(err)
	}

	results := make([]params.DeployFromRepositoryResult, len(args.Args))
	for i, entity := range args.Args {
		bindingsWithSpaceIDs, err := api.convertSpacesToIDInBindings(ctx, entity.EndpointBindings)
		if err != nil {
			results[i].Errors = []*params.Error{apiservererrors.ServerError(err)}
			continue
		}
		entity.EndpointBindings = bindingsWithSpaceIDs
		info, pending, errs := api.repoDeploy.DeployFromRepository(ctx, entity)
		if len(errs) > 0 {
			results[i].Errors = apiservererrors.ServerErrors(errs)
			continue
		}
		results[i].Info = info
		results[i].PendingResourceUploads = pending
	}
	return params.DeployFromRepositoryResults{
		Results: results,
	}, nil
}
