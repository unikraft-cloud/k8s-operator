// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Current state of the volume.
// +kubebuilder:validation:Enum=uninitialized;initializing;available;idle;mounted;busy;error;template
type CreateVolumeResponseAllOfState string

const (
	CreateVolumeResponseAllOfStateUninitialized CreateVolumeResponseAllOfState = "uninitialized"
	CreateVolumeResponseAllOfStateInitializing  CreateVolumeResponseAllOfState = "initializing"
	CreateVolumeResponseAllOfStateAvailable     CreateVolumeResponseAllOfState = "available"
	CreateVolumeResponseAllOfStateIdle          CreateVolumeResponseAllOfState = "idle"
	CreateVolumeResponseAllOfStateMounted       CreateVolumeResponseAllOfState = "mounted"
	CreateVolumeResponseAllOfStateBusy          CreateVolumeResponseAllOfState = "busy"
	CreateVolumeResponseAllOfStateError         CreateVolumeResponseAllOfState = "error"
	CreateVolumeResponseAllOfStateTemplate      CreateVolumeResponseAllOfState = "template"
)

type CreateVolumeResponseAllOf struct {
	// The status of the response.
	Status *ResponseStatus           `json:"status,omitempty"`
	Data   *CreateVolumeResponseData `json:"data,omitempty"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitempty"`
	// The operation time in microseconds.  This is the time it took to process
	// the request and generate the response.
	OpTimeUs *uint64 `json:"op_time_us,omitempty"`
	// Current state of the volume.
	State *CreateVolumeResponseAllOfState `json:"state,omitempty"`
}
