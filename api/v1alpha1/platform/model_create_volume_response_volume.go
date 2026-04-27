// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The state of the volume.
// +kubebuilder:validation:Enum=uninitialized;initializing;available;idle;mounted;busy;error;template
type CreateVolumeResponseVolumeState string

const (
	CreateVolumeResponseVolumeStateUninitialized CreateVolumeResponseVolumeState = "uninitialized"
	CreateVolumeResponseVolumeStateInitializing  CreateVolumeResponseVolumeState = "initializing"
	CreateVolumeResponseVolumeStateAvailable     CreateVolumeResponseVolumeState = "available"
	CreateVolumeResponseVolumeStateIdle          CreateVolumeResponseVolumeState = "idle"
	CreateVolumeResponseVolumeStateMounted       CreateVolumeResponseVolumeState = "mounted"
	CreateVolumeResponseVolumeStateBusy          CreateVolumeResponseVolumeState = "busy"
	CreateVolumeResponseVolumeStateError         CreateVolumeResponseVolumeState = "error"
	CreateVolumeResponseVolumeStateTemplate      CreateVolumeResponseVolumeState = "template"
)

type CreateVolumeResponseVolume struct {
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// UUID of the newly created volume.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the newly created volume.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the volume.
	Metro *string `json:"metro,omitempty"`
	// The state of the volume.
	State *CreateVolumeResponseVolumeState `json:"state,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
}
