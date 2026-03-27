// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The response message for creating of a volume.
// Current state of the volume.
// +kubebuilder:validation:Enum=uninitialized;initializing;available;idle;mounted;busy;error;template
type CreateVolumeResponseState string

const (
	CreateVolumeResponseStateUninitialized CreateVolumeResponseState = "uninitialized"
	CreateVolumeResponseStateInitializing  CreateVolumeResponseState = "initializing"
	CreateVolumeResponseStateAvailable     CreateVolumeResponseState = "available"
	CreateVolumeResponseStateIdle          CreateVolumeResponseState = "idle"
	CreateVolumeResponseStateMounted       CreateVolumeResponseState = "mounted"
	CreateVolumeResponseStateBusy          CreateVolumeResponseState = "busy"
	CreateVolumeResponseStateError         CreateVolumeResponseState = "error"
	CreateVolumeResponseStateTemplate      CreateVolumeResponseState = "template"
)

type CreateVolumeResponse struct {
	// The status of the response.
	Status *ResponseStatus           `json:"status,omitempty"`
	Data   *CreateVolumeResponseData `json:"data,omitempty"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitempty"`
	// The operation time in microseconds.  This is the time it took to process
	// the request and generate the response.
	OpTimeUs *uint64 `json:"op_time_us,omitempty"`
	// Current state of the volume.
	State *CreateVolumeResponseState `json:"state,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
}
