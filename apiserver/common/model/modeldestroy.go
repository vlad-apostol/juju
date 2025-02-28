// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package model

import (
	"context"
	"time"

	"github.com/juju/errors"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/core/status"
	"github.com/juju/juju/domain/blockcommand"
	modelerrors "github.com/juju/juju/domain/model/errors"
	internallogger "github.com/juju/juju/internal/logger"
	"github.com/juju/juju/state"
	stateerrors "github.com/juju/juju/state/errors"
)

var logger = internallogger.GetLogger("juju.apiserver.common")

// BlockCommandService defines methods for interacting with block commands.
type BlockCommandService interface {
	// GetBlockSwitchedOn returns the optional block message if it is switched
	// on for the given type.
	GetBlockSwitchedOn(ctx context.Context, t blockcommand.BlockType) (string, error)

	// GetBlocks returns all the blocks that are currently in place.
	GetBlocks(ctx context.Context) ([]blockcommand.Block, error)
}

// DestroyController sets the controller model to Dying and, if requested,
// schedules cleanups so that all of the hosted models are destroyed, or
// otherwise returns an error indicating that there are hosted models
// remaining.
func DestroyController(
	ctx context.Context,
	st ModelManagerBackend,
	blockCommandService BlockCommandService,
	modelInfoService ModelInfoService,
	blockCommandServiceGetter func(context.Context, model.UUID) (BlockCommandService, error),
	destroyHostedModels bool,
	destroyStorage *bool,
	force *bool,
	maxWait *time.Duration,
	modelTimeout *time.Duration,
) error {
	modelTag := st.ModelTag()
	controllerModelTag := st.ControllerModelTag()
	if modelTag != controllerModelTag {
		return errors.Errorf(
			"expected state for controller model UUID %v, got %v",
			controllerModelTag.Id(),
			modelTag.Id(),
		)
	}
	if destroyHostedModels {
		uuids, err := st.AllModelUUIDs()
		if err != nil {
			return errors.Trace(err)
		}
		for _, uuid := range uuids {
			svc, err := blockCommandServiceGetter(ctx, model.UUID(uuid))
			if err != nil {
				return errors.Trace(err)
			}

			check := common.NewBlockChecker(svc)
			if err = check.DestroyAllowed(ctx); errors.Is(err, modelerrors.NotFound) {
				logger.Errorf(ctx, "model %v not found, skipping", uuid)
				continue
			} else if err != nil {
				return errors.Trace(err)
			}
		}
	}
	return destroyModel(ctx, st, blockCommandService, modelInfoService, state.DestroyModelParams{
		DestroyHostedModels: destroyHostedModels,
		DestroyStorage:      destroyStorage,
		Force:               force,
		MaxWait:             common.MaxWait(maxWait),
		Timeout:             modelTimeout,
	})
}

// DestroyModel sets the model to Dying, such that the model's resources will
// be destroyed and the model removed from the controller.
func DestroyModel(
	ctx context.Context,
	st ModelManagerBackend,
	blockCommandService BlockCommandService,
	modelInfoService ModelInfoService,
	destroyStorage *bool,
	force *bool,
	maxWait *time.Duration,
	timeout *time.Duration,
) error {
	return destroyModel(ctx, st, blockCommandService, modelInfoService, state.DestroyModelParams{
		DestroyStorage: destroyStorage,
		Force:          force,
		MaxWait:        common.MaxWait(maxWait),
		Timeout:        timeout,
	})
}

func destroyModel(
	ctx context.Context, st ModelManagerBackend,
	blockCommandService BlockCommandService, modelInfoService ModelInfoService,
	args state.DestroyModelParams,
) error {
	check := common.NewBlockChecker(blockCommandService)
	if err := check.DestroyAllowed(ctx); err != nil {
		return errors.Trace(err)
	}

	model, err := st.Model()
	if err != nil {
		return errors.Trace(err)
	}
	notForcing := args.Force == nil || !*args.Force
	if notForcing {
		// If model status is suspended, then model's cloud credential is invalid.
		modelStatus, err := modelInfoService.GetStatus(ctx)
		if err != nil {
			return errors.Trace(err)
		}
		if modelStatus.Status == status.Suspended {
			return errors.Errorf("invalid cloud credential, use --force")
		}
	}
	if err := model.Destroy(args); err != nil {
		if notForcing {
			return errors.Trace(err)
		}
		logger.Warningf(context.TODO(), "failed destroying model %v: %v", model.UUID(), err)
		if err := filterNonCriticalErrorForForce(err); err != nil {
			return errors.Trace(err)
		}
	}

	// Return to the caller. If it's the CLI, it will finish up by calling the
	// provider's Destroy method, which will destroy the controllers, any
	// straggler instances, and other provider-specific resources. Once all
	// resources are torn down, the Undertaker worker handles the removal of
	// the model.
	return nil
}

func filterNonCriticalErrorForForce(err error) error {
	if errors.Is(err, stateerrors.PersistentStorageError) {
		return err
	}
	return nil
}
