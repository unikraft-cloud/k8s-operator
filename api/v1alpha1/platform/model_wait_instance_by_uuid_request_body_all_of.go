// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The desired state to wait for.  Default is `running`.
// +kubebuilder:validation:Enum=stopped;starting;running;draining;stopping;template;standby
type WaitInstanceByUUIDRequestBodyAllOfState string

const (
	WaitInstanceByUUIDRequestBodyAllOfStateStopped  WaitInstanceByUUIDRequestBodyAllOfState = "stopped"
	WaitInstanceByUUIDRequestBodyAllOfStateStarting WaitInstanceByUUIDRequestBodyAllOfState = "starting"
	WaitInstanceByUUIDRequestBodyAllOfStateRunning  WaitInstanceByUUIDRequestBodyAllOfState = "running"
	WaitInstanceByUUIDRequestBodyAllOfStateDraining WaitInstanceByUUIDRequestBodyAllOfState = "draining"
	WaitInstanceByUUIDRequestBodyAllOfStateStopping WaitInstanceByUUIDRequestBodyAllOfState = "stopping"
	WaitInstanceByUUIDRequestBodyAllOfStateTemplate WaitInstanceByUUIDRequestBodyAllOfState = "template"
	WaitInstanceByUUIDRequestBodyAllOfStateStandby  WaitInstanceByUUIDRequestBodyAllOfState = "standby"
)

type WaitInstanceByUUIDRequestBodyAllOf struct {
	// The desired state to wait for.  Default is `running`.
	State *WaitInstanceByUUIDRequestBodyAllOfState `json:"state,omitempty"`
}
