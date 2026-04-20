// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The current state of the instance.
// +kubebuilder:validation:Enum=stopped;starting;running;draining;stopping;template;standby
type SuspendInstancesResponseSuspendedInstanceState string

const (
	SuspendInstancesResponseSuspendedInstanceStateStopped  SuspendInstancesResponseSuspendedInstanceState = "stopped"
	SuspendInstancesResponseSuspendedInstanceStateStarting SuspendInstancesResponseSuspendedInstanceState = "starting"
	SuspendInstancesResponseSuspendedInstanceStateRunning  SuspendInstancesResponseSuspendedInstanceState = "running"
	SuspendInstancesResponseSuspendedInstanceStateDraining SuspendInstancesResponseSuspendedInstanceState = "draining"
	SuspendInstancesResponseSuspendedInstanceStateStopping SuspendInstancesResponseSuspendedInstanceState = "stopping"
	SuspendInstancesResponseSuspendedInstanceStateTemplate SuspendInstancesResponseSuspendedInstanceState = "template"
	SuspendInstancesResponseSuspendedInstanceStateStandby  SuspendInstancesResponseSuspendedInstanceState = "standby"
)

// The previous state of the instance before the suspend operation was invoked.
// +kubebuilder:validation:Enum=stopped;starting;running;draining;stopping;template;standby
type SuspendInstancesResponseSuspendedInstancePreviousState string

const (
	SuspendInstancesResponseSuspendedInstancePreviousStateStopped  SuspendInstancesResponseSuspendedInstancePreviousState = "stopped"
	SuspendInstancesResponseSuspendedInstancePreviousStateStarting SuspendInstancesResponseSuspendedInstancePreviousState = "starting"
	SuspendInstancesResponseSuspendedInstancePreviousStateRunning  SuspendInstancesResponseSuspendedInstancePreviousState = "running"
	SuspendInstancesResponseSuspendedInstancePreviousStateDraining SuspendInstancesResponseSuspendedInstancePreviousState = "draining"
	SuspendInstancesResponseSuspendedInstancePreviousStateStopping SuspendInstancesResponseSuspendedInstancePreviousState = "stopping"
	SuspendInstancesResponseSuspendedInstancePreviousStateTemplate SuspendInstancesResponseSuspendedInstancePreviousState = "template"
	SuspendInstancesResponseSuspendedInstancePreviousStateStandby  SuspendInstancesResponseSuspendedInstancePreviousState = "standby"
)

type SuspendInstancesResponseSuspendedInstance struct {
	// The UUID of the instance.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the instance.
	Metro *string `json:"metro,omitempty"`
	// The current state of the instance.
	State *SuspendInstancesResponseSuspendedInstanceState `json:"state,omitempty"`
	// The previous state of the instance before the suspend operation was invoked.
	PreviousState *SuspendInstancesResponseSuspendedInstancePreviousState `json:"previous_state,omitempty"`
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
}
