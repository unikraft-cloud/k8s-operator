// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

type GetInstancesLogsResponseLoggedInstance struct {
	// The UUID of the instance.
	Uuid string `json:"uuid"`
	// The name of the instance.
	Name string `json:"name"`
	// (Only applies when using global control plane).
	// The metro of the instance.
	Metro *string `json:"metro,omitempty"`
	// Base64 encoded log output of the instance.
	Output    string                                          `json:"output"`
	Available GetInstancesLogsResponseLoggedInstanceAvailable `json:"available"`
	Range     GetInstancesLogsResponseLoggedInstanceRange     `json:"range"`
	// State of the instance when the logs were retrieved.
	State InstanceState `json:"state"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
}
