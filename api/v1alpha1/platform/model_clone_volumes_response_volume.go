// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The state of the volume.
// +kubebuilder:validation:Enum=uninitialized;initializing;available;idle;mounted;busy;error;template
type CloneVolumesResponseVolumeState string

const (
	CloneVolumesResponseVolumeStateUninitialized CloneVolumesResponseVolumeState = "uninitialized"
	CloneVolumesResponseVolumeStateInitializing  CloneVolumesResponseVolumeState = "initializing"
	CloneVolumesResponseVolumeStateAvailable     CloneVolumesResponseVolumeState = "available"
	CloneVolumesResponseVolumeStateIdle          CloneVolumesResponseVolumeState = "idle"
	CloneVolumesResponseVolumeStateMounted       CloneVolumesResponseVolumeState = "mounted"
	CloneVolumesResponseVolumeStateBusy          CloneVolumesResponseVolumeState = "busy"
	CloneVolumesResponseVolumeStateError         CloneVolumesResponseVolumeState = "error"
	CloneVolumesResponseVolumeStateTemplate      CloneVolumesResponseVolumeState = "template"
)

type CloneVolumesResponseVolume struct {
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// The UUID of the newly cloned volume.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the newly cloned volume.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the volume.
	Metro *string `json:"metro,omitempty"`
	// The state of the volume.
	State *CloneVolumesResponseVolumeState `json:"state,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
}
