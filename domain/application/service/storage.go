// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	corestorage "github.com/juju/juju/core/storage"
	"github.com/juju/juju/core/unit"
	"github.com/juju/juju/internal/errors"
	"github.com/juju/juju/internal/storage"
)

// StorageState describes retrieval and persistence methods for
// storage related interactions.
type StorageState interface {
	// AttachStorage attaches the specified storage to the specified unit.
	// The following error types can be expected:
	// - [applicationerrors.UnitNotFound]: when the unit does not exist.
	// - [applicationerrors.StorageAttachmentAlreadyExistsError]: when the attachment already exists.
	// - [applicationerrors.WrongStorageOwnerError]: when the storage owner does not match the unit.
	// - [applicationerrors.UnitNotAlive]: when the unit is not alive.
	// - [applicationerrors.StorageNotAlive]: when the storage is not alive.
	// - [applicationerrors.StorageNameNotSupported]: when storage name is not defined in charm metadata.
	// - [applicationerrors.InvalidStorageCount]: when the allowed attachment count would be violated.
	// - [applicationerrors.InvalidStorageMountPoint]: when the filesystem being attached to the unit's machine has a mount point path conflict.
	AttachStorage(ctx context.Context, storageID corestorage.ID, unitUUID unit.UUID) error

	// AddStorageForUnit attaches the specified storage to the specified unit.
	// The following error types can be expected:
	// - [applicationerrors.UnitNotFound]: when the unit does not exist.
	// - [applicationerrors.UnitNotAlive]: when the unit is not alive.
	// - [applicationerrors.StorageNotAlive]: when the storage is not alive.
	// - [applicationerrors.StorageNameNotSupported]: when storage name is not defined in charm metadata.
	// - [applicationerrors.InvalidStorageCount]: when the allowed attachment count would be violated.
	// - [applicationerrors.InvalidStorageMountPoint]: when the filesystem being attached to the unit's machine has a mount point path conflict.
	AddStorageForUnit(
		ctx context.Context, storageName corestorage.Name, unitUUID unit.UUID, stor storage.Directive,
	) ([]corestorage.ID, error)

	// DetachStorageForUnit detaches the specified storage from the specified unit.
	// The following error types can be expected:
	// - [applicationerrors.UnitNotFound]: when the unit does not exist.
	// - [applicationerrors.StorageNotDetachable]: when the type of storage is not detachable.
	DetachStorageForUnit(ctx context.Context, storageID corestorage.ID, unitUUID unit.UUID) error

	// DetachStorage detaches the specified storage from whatever node it is attached to.
	// The following error types can be expected:
	// - [applicationerrors.StorageNotDetachable]: when the type of storage is not detachable.
	DetachStorage(ctx context.Context, storageID corestorage.ID) error
}

// AttachStorage attached the specified storage to the specified unit.
// The following error types can be expected:
// - [unit.InvalidUnitName]: when the unit name is not valid.
// - [corestorage.InvalidStorageID]: when the storage ID is not valid.
// - [applicationerrors.UnitNotFound]: when the unit does not exist.
// - [applicationerrors.StorageAttachmentAlreadyExistsError]: when the attachment already exists.
// - [applicationerrors.WrongStorageOwnerError]: when the storage owner does not match the unit.
// - [applicationerrors.UnitNotAlive]: when the unit is not alive.
// - [applicationerrors.StorageNotAlive]: when the storage is not alive.
// - [applicationerrors.StorageNameNotSupported]: when storage name is not defined in charm metadata.
// - [applicationerrors.InvalidStorageCount]: when the allowed attachment count would be violated.
// - [applicationerrors.InvalidStorageMountPoint]: when the filesystem being attached to the unit's machine has a mount point path conflict.
func (s *Service) AttachStorage(ctx context.Context, storageID corestorage.ID, unitName unit.Name) error {
	if err := unitName.Validate(); err != nil {
		return errors.Capture(err)
	}
	if err := storageID.Validate(); err != nil {
		return errors.Capture(err)
	}
	unitUUID, err := s.st.GetUnitUUIDByName(ctx, unitName)
	if err != nil {
		return errors.Capture(err)
	}
	return s.st.AttachStorage(ctx, storageID, unitUUID)
}

// AddStorageForUnit adds storage instances to given unit.
// Missing storage constraints are populated based on model defaults.
// The following error types can be expected:
// - [unit.InvalidUnitName]: when the unit name is not valid.
// - [corestorage.InvalidStorageName]: when the storage name is not valid.
// - [applicationerrors.UnitNotFound]: when the unit does not exist.
// - [applicationerrors.UnitNotFound]: when the unit does not exist.
// - [applicationerrors.UnitNotAlive]: when the unit is not alive.
// - [applicationerrors.StorageNotAlive]: when the storage is not alive.
// - [applicationerrors.StorageNameNotSupported]: when storage name is not defined in charm metadata.
// - [applicationerrors.InvalidStorageCount]: when the allowed attachment count would be violated.
// - [applicationerrors.InvalidStorageMountPoint]: when the filesystem being attached to the unit's machine has a mount point path conflict.
func (s *Service) AddStorageForUnit(
	ctx context.Context, storageName corestorage.Name, unitName unit.Name, stor storage.Directive,
) ([]corestorage.ID, error) {
	if err := unitName.Validate(); err != nil {
		return nil, errors.Capture(err)
	}
	if err := storageName.Validate(); err != nil {
		return nil, errors.Capture(err)
	}
	unitUUID, err := s.st.GetUnitUUIDByName(ctx, unitName)
	if err != nil {
		return nil, errors.Capture(err)
	}
	return s.st.AddStorageForUnit(ctx, storageName, unitUUID, stor)
}

// DetachStorageForUnit detaches the specified storage from the specified unit.
// The following error types can be expected:
// - [unit.InvalidUnitName]: when the unit name is not valid.
// - [corestorage.InvalidStorageID]: when the storage ID is not valid.
// - [applicationerrors.UnitNotFound]: when the unit does not exist.
// - [applicationerrors.StorageNotDetachable]: when the type of storage is not detachable.
func (s *Service) DetachStorageForUnit(ctx context.Context, storageID corestorage.ID, unitName unit.Name) error {
	if err := unitName.Validate(); err != nil {
		return errors.Capture(err)
	}
	if err := storageID.Validate(); err != nil {
		return errors.Capture(err)
	}
	unitUUID, err := s.st.GetUnitUUIDByName(ctx, unitName)
	if err != nil {
		return errors.Capture(err)
	}
	return s.st.DetachStorageForUnit(ctx, storageID, unitUUID)
}

// DetachStorage detaches the specified storage from whatever node it is attached to.
// The following error types can be expected:
// - [corestorage.InvalidStorageName]: when the storage name is not valid.
// - [applicationerrors.StorageNotDetachable]: when the type of storage is not detachable.
func (s *Service) DetachStorage(ctx context.Context, storageID corestorage.ID) error {
	if err := storageID.Validate(); err != nil {
		return errors.Capture(err)
	}
	return s.st.DetachStorage(ctx, storageID)
}
