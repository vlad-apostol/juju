// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package modelconfig_test

import (
	"context"

	"github.com/juju/errors"
	"github.com/juju/names/v5"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/facades/client/modelconfig"
	"github.com/juju/juju/apiserver/facades/client/modelconfig/mocks"
	apiservertesting "github.com/juju/juju/apiserver/testing"
	"github.com/juju/juju/core/constraints"
	coresecrets "github.com/juju/juju/core/secrets"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/internal/feature"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
	coretesting "github.com/juju/juju/testing"
)

type modelconfigSuite struct {
	testing.IsolationSuite
	coretesting.JujuOSEnvSuite
	backend                  *mockBackend
	authorizer               apiservertesting.FakeAuthorizer
	mockSecretBackendService *mocks.MockSecretBackendService
}

var _ = gc.Suite(&modelconfigSuite{})

func (s *modelconfigSuite) SetUpTest(c *gc.C) {
	s.SetInitialFeatureFlags(feature.DeveloperMode)
	s.IsolationSuite.SetUpTest(c)
	s.JujuOSEnvSuite.SetUpTest(c)
	s.authorizer = apiservertesting.FakeAuthorizer{
		Tag:      names.NewUserTag("bruce@local"),
		AdminTag: names.NewUserTag("bruce@local"),
	}
	s.backend = &mockBackend{
		cfg: config.ConfigValues{
			"type":            {Value: "dummy", Source: "model"},
			"agent-version":   {Value: "1.2.3.4", Source: "model"},
			"ftp-proxy":       {Value: "http://proxy", Source: "model"},
			"authorized-keys": {Value: coretesting.FakeAuthKeys, Source: "model"},
			"charmhub-url":    {Value: "http://meshuggah.rocks", Source: "model"},
		},
		secretBackend: &coresecrets.SecretBackend{
			ID:          "backend-1",
			Name:        "backend-1",
			BackendType: "vault",
			Config: map[string]interface{}{
				"endpoint": "http://0.0.0.0:8200",
			},
		},
	}
}

func (s *modelconfigSuite) getAPI(c *gc.C) (*modelconfig.ModelConfigAPIV3, *gomock.Controller) {
	ctrl := gomock.NewController(c)
	s.mockSecretBackendService = mocks.NewMockSecretBackendService(ctrl)
	api, err := modelconfig.NewModelConfigAPI(s.backend, s.mockSecretBackendService, &s.authorizer)
	c.Assert(err, jc.ErrorIsNil)
	return api, ctrl
}

func (s *modelconfigSuite) TestAdminModelGet(c *gc.C) {
	api, _ := s.getAPI(c)

	result, err := api.ModelGet(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result.Config, jc.DeepEquals, map[string]params.ConfigValue{
		"type":          {Value: "dummy", Source: "model"},
		"ftp-proxy":     {Value: "http://proxy", Source: "model"},
		"agent-version": {Value: "1.2.3.4", Source: "model"},
		"charmhub-url":  {Value: "http://meshuggah.rocks", Source: "model"},
	})
}

func (s *modelconfigSuite) TestUserModelGet(c *gc.C) {
	api, _ := s.getAPI(c)
	s.authorizer = apiservertesting.FakeAuthorizer{
		Tag:         names.NewUserTag("bruce@local"),
		HasWriteTag: names.NewUserTag("bruce@local"),
		AdminTag:    names.NewUserTag("mary@local"),
	}
	result, err := api.ModelGet(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result.Config, jc.DeepEquals, map[string]params.ConfigValue{
		"type":          {Value: "dummy", Source: "model"},
		"ftp-proxy":     {Value: "http://proxy", Source: "model"},
		"agent-version": {Value: "1.2.3.4", Source: "model"},
		"charmhub-url":  {Value: "http://meshuggah.rocks", Source: "model"},
	})
}

func (s *modelconfigSuite) assertConfigValue(c *gc.C, key string, expected interface{}) {
	value, found := s.backend.cfg[key]
	c.Assert(found, jc.IsTrue)
	c.Assert(value.Value, gc.Equals, expected)
}

func (s *modelconfigSuite) assertConfigValueMissing(c *gc.C, key string) {
	_, found := s.backend.cfg[key]
	c.Assert(found, jc.IsFalse)
}

func (s *modelconfigSuite) TestAdminModelSet(c *gc.C) {
	api, _ := s.getAPI(c)
	params := params.ModelSet{
		Config: map[string]interface{}{
			"some-key":  "value",
			"other-key": "other value",
		},
	}
	err := api.ModelSet(context.Background(), params)
	c.Assert(err, jc.ErrorIsNil)
	s.assertConfigValue(c, "some-key", "value")
	s.assertConfigValue(c, "other-key", "other value")
}

func (s *modelconfigSuite) blockAllChanges(msg string) {
	s.backend.msg = msg
	s.backend.b = state.ChangeBlock
}

func (s *modelconfigSuite) assertBlocked(c *gc.C, err error, msg string) {
	c.Assert(params.IsCodeOperationBlocked(err), jc.IsTrue, gc.Commentf("error: %#v", err))
	c.Assert(errors.Cause(err), jc.DeepEquals, &params.Error{
		Message: msg,
		Code:    "operation is blocked",
	})
}

func (s *modelconfigSuite) assertModelSetBlocked(c *gc.C, args map[string]interface{}, msg string) {
	api, _ := s.getAPI(c)
	err := api.ModelSet(context.Background(), params.ModelSet{Config: args})
	s.assertBlocked(c, err, msg)
}

func (s *modelconfigSuite) TestBlockChangesModelSet(c *gc.C) {
	s.blockAllChanges("TestBlockChangesModelSet")
	args := map[string]interface{}{"some-key": "value"}
	s.assertModelSetBlocked(c, args, "TestBlockChangesModelSet")
}

func (s *modelconfigSuite) TestModelSetCannotChangeAgentVersion(c *gc.C) {
	api, _ := s.getAPI(c)
	old, err := config.New(config.UseDefaults, coretesting.FakeConfig().Merge(coretesting.Attrs{
		"agent-version": "1.2.3.4",
	}))
	c.Assert(err, jc.ErrorIsNil)
	s.backend.old = old
	args := params.ModelSet{
		Config: map[string]interface{}{"agent-version": "9.9.9"},
	}
	err = api.ModelSet(context.Background(), args)
	c.Assert(err, gc.ErrorMatches, "agent-version cannot be changed")

	// It's okay to pass config back with the same agent-version.
	result, err := api.ModelGet(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result.Config["agent-version"], gc.NotNil)
	args.Config["agent-version"] = result.Config["agent-version"].Value
	err = api.ModelSet(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *modelconfigSuite) TestModelSetCannotChangeCharmHubURL(c *gc.C) {
	api, _ := s.getAPI(c)
	old, err := config.New(config.UseDefaults, coretesting.FakeConfig().Merge(coretesting.Attrs{
		"charmhub-url": "http://meshuggah.rocks",
	}))
	c.Assert(err, jc.ErrorIsNil)
	s.backend.old = old
	args := params.ModelSet{
		Config: map[string]interface{}{"charmhub-url": "http://another-url.com"},
	}
	err = api.ModelSet(context.Background(), args)
	c.Assert(err, gc.ErrorMatches, "charmhub-url cannot be changed")

	// It's okay to pass config back with the same charmhub-url.
	result, err := api.ModelGet(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result.Config["charmhub-url"], gc.NotNil)
	args.Config["charmhub-url"] = result.Config["charmhub-url"].Value
	err = api.ModelSet(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *modelconfigSuite) TestModelSetCannotSetAuthorizedKeys(c *gc.C) {
	api, _ := s.getAPI(c)
	// Try to set the authorized-keys model config.
	args := params.ModelSet{
		Config: map[string]interface{}{"authorized-keys": "ssh-rsa new Juju:juju-client-key"},
	}
	err := api.ModelSet(context.Background(), args)
	c.Assert(err, gc.ErrorMatches, "authorized-keys cannot be set")
	// Make sure the authorized-keys still contains its original value.
	s.assertConfigValue(c, "authorized-keys", coretesting.FakeAuthKeys)
}

func (s *modelconfigSuite) TestAdminCanSetLogTrace(c *gc.C) {
	api, _ := s.getAPI(c)
	args := params.ModelSet{
		Config: map[string]interface{}{"logging-config": "<root>=DEBUG;somepackage=TRACE"},
	}
	err := api.ModelSet(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)

	result, err := api.ModelGet(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result.Config["logging-config"].Value, gc.Equals, "<root>=DEBUG;somepackage=TRACE")
}

func (s *modelconfigSuite) TestUserCanSetLogNoTrace(c *gc.C) {
	api, _ := s.getAPI(c)
	args := params.ModelSet{
		Config: map[string]interface{}{"logging-config": "<root>=DEBUG;somepackage=ERROR"},
	}
	apiUser := names.NewUserTag("fred")
	s.authorizer.Tag = apiUser
	s.authorizer.HasWriteTag = apiUser
	err := api.ModelSet(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)

	result, err := api.ModelGet(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result.Config["logging-config"].Value, gc.Equals, "<root>=DEBUG;somepackage=ERROR")
}

func (s *modelconfigSuite) TestUserReadAccess(c *gc.C) {
	api, _ := s.getAPI(c)
	apiUser := names.NewUserTag("read")
	s.authorizer.Tag = apiUser

	_, err := api.ModelGet(context.Background())
	c.Assert(err, jc.ErrorIsNil)

	err = api.ModelSet(context.Background(), params.ModelSet{})
	c.Assert(errors.Cause(err), gc.ErrorMatches, "permission denied")
}

func (s *modelconfigSuite) TestUserCannotSetLogTrace(c *gc.C) {
	api, _ := s.getAPI(c)
	args := params.ModelSet{
		Config: map[string]interface{}{"logging-config": "<root>=DEBUG;somepackage=TRACE"},
	}
	apiUser := names.NewUserTag("fred")
	s.authorizer.Tag = apiUser
	s.authorizer.HasWriteTag = apiUser
	err := api.ModelSet(context.Background(), args)
	c.Assert(err, gc.ErrorMatches, `only controller admins can set a model's logging level to TRACE`)
}

func (s *modelconfigSuite) TestSetSecretBackend(c *gc.C) {
	api, _ := s.getAPI(c)
	args := params.ModelSet{
		Config: map[string]interface{}{"secret-backend": 1},
	}
	err := api.ModelSet(context.Background(), args)
	c.Assert(err, gc.ErrorMatches, `"secret-backend" config value is not a string`)

	args.Config = map[string]interface{}{"secret-backend": ""}
	err = api.ModelSet(context.Background(), args)
	c.Assert(err, gc.ErrorMatches, `empty "secret-backend" config value not valid`)

	args.Config = map[string]interface{}{"secret-backend": "auto"}
	err = api.ModelSet(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)
	result, err := api.ModelGet(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result.Config["secret-backend"].Value, gc.Equals, "auto")
}

func (s *modelconfigSuite) TestSetSecretBackendExternal(c *gc.C) {
	api, _ := s.getAPI(c)
	ctrl := gomock.NewController(c)
	defer ctrl.Finish()

	s.mockSecretBackendService.EXPECT().PingSecretBackend(gomock.Any(), "backend-1").Return(nil)

	args := params.ModelSet{
		Config: map[string]interface{}{"secret-backend": "backend-1"},
	}
	err := api.ModelSet(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)
	result, err := api.ModelGet(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result.Config["secret-backend"].Value, gc.Equals, "backend-1")
}

func (s *modelconfigSuite) TestSetSecretBackendExternalValidationFailed(c *gc.C) {
	api, ctrl := s.getAPI(c)
	defer ctrl.Finish()

	s.mockSecretBackendService.EXPECT().PingSecretBackend(gomock.Any(), "backend-1").Return(errors.New("not reachable"))

	args := params.ModelSet{
		Config: map[string]interface{}{"secret-backend": "backend-1"},
	}
	err := api.ModelSet(context.Background(), args)
	c.Assert(err, gc.ErrorMatches, `cannot ping backend "backend-1": not reachable`)
}

func (s *modelconfigSuite) TestModelUnset(c *gc.C) {
	api, _ := s.getAPI(c)
	err := s.backend.UpdateModelConfig(map[string]interface{}{"abc": 123}, nil)
	c.Assert(err, jc.ErrorIsNil)

	args := params.ModelUnset{Keys: []string{"abc"}}
	err = api.ModelUnset(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)
	s.assertConfigValueMissing(c, "abc")
}

func (s *modelconfigSuite) TestBlockModelUnset(c *gc.C) {
	api, _ := s.getAPI(c)
	err := s.backend.UpdateModelConfig(map[string]interface{}{"abc": 123}, nil)
	c.Assert(err, jc.ErrorIsNil)
	s.blockAllChanges("TestBlockModelUnset")

	args := params.ModelUnset{Keys: []string{"abc"}}
	err = api.ModelUnset(context.Background(), args)
	s.assertBlocked(c, err, "TestBlockModelUnset")
}

func (s *modelconfigSuite) TestModelUnsetMissing(c *gc.C) {
	api, _ := s.getAPI(c)
	// It's okay to unset a non-existent attribute.
	args := params.ModelUnset{Keys: []string{"not_there"}}
	err := api.ModelUnset(context.Background(), args)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *modelconfigSuite) TestClientSetModelConstraints(c *gc.C) {
	api, _ := s.getAPI(c)
	// Set constraints for the model.
	cons, err := constraints.Parse("mem=4096", "cores=2")
	c.Assert(err, jc.ErrorIsNil)
	err = api.SetModelConstraints(context.Background(), params.SetConstraints{
		ApplicationName: "app",
		Constraints:     cons,
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(s.backend.cons, gc.DeepEquals, cons)
}

func (s *modelconfigSuite) assertSetModelConstraintsBlocked(c *gc.C, msg string) {
	api, _ := s.getAPI(c)
	// Set constraints for the model.
	cons, err := constraints.Parse("mem=4096", "cores=2")
	c.Assert(err, jc.ErrorIsNil)
	err = api.SetModelConstraints(context.Background(), params.SetConstraints{
		ApplicationName: "app",
		Constraints:     cons,
	})
	s.assertBlocked(c, err, msg)
}

func (s *modelconfigSuite) TestBlockChangesClientSetModelConstraints(c *gc.C) {
	s.blockAllChanges("TestBlockChangesClientSetModelConstraints")
	s.assertSetModelConstraintsBlocked(c, "TestBlockChangesClientSetModelConstraints")
}

func (s *modelconfigSuite) TestClientGetModelConstraints(c *gc.C) {
	api, _ := s.getAPI(c)
	// Set constraints for the model.
	cons, err := constraints.Parse("mem=4096", "cores=2")
	c.Assert(err, jc.ErrorIsNil)
	s.backend.cons = cons
	obtained, err := api.GetModelConstraints(context.Background())
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(obtained.Constraints, gc.DeepEquals, cons)
}

type mockBackend struct {
	cfg           config.ConfigValues
	old           *config.Config
	b             state.BlockType
	msg           string
	cons          constraints.Value
	secretBackend *coresecrets.SecretBackend
}

func (m *mockBackend) SetModelConstraints(value constraints.Value) error {
	m.cons = value
	return nil
}

func (m *mockBackend) ModelConstraints() (constraints.Value, error) {
	return m.cons, nil
}

func (m *mockBackend) ModelConfigValues() (config.ConfigValues, error) {
	return m.cfg, nil
}

func (m *mockBackend) Sequences() (map[string]int, error) {
	return nil, nil
}

func (m *mockBackend) UpdateModelConfig(update map[string]interface{}, remove []string,
	validate ...state.ValidateConfigFunc) error {
	for _, validateFunc := range validate {
		if err := validateFunc(update, remove, m.old); err != nil {
			return err
		}
	}
	for k, v := range update {
		m.cfg[k] = config.ConfigValue{Value: v, Source: "model"}
	}
	for _, n := range remove {
		delete(m.cfg, n)
	}
	return nil
}

func (m *mockBackend) GetBlockForType(t state.BlockType) (state.Block, bool, error) {
	if m.b == t {
		return &mockBlock{t: t, m: m.msg}, true, nil
	} else {
		return nil, false, nil
	}
}

func (m *mockBackend) ModelTag() names.ModelTag {
	return names.NewModelTag("deadbeef-2f18-4fd2-967d-db9663db7bea")
}

func (m *mockBackend) ControllerTag() names.ControllerTag {
	return names.NewControllerTag("deadbeef-babe-4fd2-967d-db9663db7bea")
}

func (m *mockBackend) SpaceByName(string) error {
	return nil
}

func (m *mockBackend) GetSecretBackend(name string) (*coresecrets.SecretBackend, error) {
	if name == "invalid" {
		return nil, errors.NotFoundf("invalid")
	}
	return m.secretBackend, nil
}

type mockBlock struct {
	state.Block
	t state.BlockType
	m string
}

func (m mockBlock) Id() string { return "" }

func (m mockBlock) Tag() (names.Tag, error) { return names.NewModelTag("mocktesting"), nil }

func (m mockBlock) Type() state.BlockType { return m.t }

func (m mockBlock) Message() string { return m.m }

func (m mockBlock) ModelUUID() string { return "" }
