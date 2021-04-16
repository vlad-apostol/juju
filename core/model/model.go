// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package model

import (
	"github.com/juju/charm/v8"
	"github.com/juju/collections/set"
	"github.com/juju/errors"

	"github.com/juju/juju/core/series"
)

// ModelType indicates a model type.
type ModelType string

const (
	// IAAS is the type for IAAS models.
	IAAS ModelType = "iaas"

	// CAAS is the type for CAAS models.
	CAAS ModelType = "caas"
)

// String returns m as a string.
func (m ModelType) String() string {
	return string(m)
}

// Model represents the state of a model.
type Model struct {
	// Name returns the human friendly name of the model.
	Name string

	// UUID is the universally unique identifier of the model.
	UUID string

	// ModelType is the type of model.
	ModelType ModelType
}

// ValidateModelTarget ensures the charm is valid for the model target type.
func ValidateModelTarget(modelType ModelType, meta *charm.Meta) error {
	isIAAS := true
	if set.NewStrings(meta.Series...).Contains(series.Kubernetes.String()) || len(meta.Containers) > 0 {
		isIAAS = false
	}

	switch modelType {
	case CAAS:
		if isIAAS {
			return errors.NotValidf("non container base charm for container based model type")
		}
	case IAAS:
		if !isIAAS {
			return errors.NotValidf("container base charm for non container based model type")
		}
	default:
		return errors.Errorf("invalid model type, expected CAAS or IAAS")
	}
	return nil
}
