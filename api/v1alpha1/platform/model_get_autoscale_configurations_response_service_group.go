// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Per-item result for a get autoscale configurations operation.
type GetAutoscaleConfigurationsResponseServiceGroup struct {
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
	// The name of the node the resource is placed on. Only present on metros
	// that span multiple nodes.
	Node *string `json:"node,omitempty"`
	// If the autoscale configuration is enabled.
	Enabled bool `json:"enabled"`
	// The minimum number of instances to keep running. Only if enabled is true.
	MinSize *int64 `json:"min_size,omitempty"`
	// The maximum number of instances to keep running. Only if enabled is true.
	MaxSize *int64 `json:"max_size,omitempty"`
	// The warmup time in seconds for new instances. Only if enabled is true.
	WarmupTimeMs *int64 `json:"warmup_time_ms,omitempty"`
	// The cooldown time in seconds for the autoscale configuration. Only if
	// enabled is true.
	CooldownTimeMs *int64 `json:"cooldown_time_ms,omitempty"`
	// The instance template used for the autoscale configuration. Only if
	// enabled is true.
	Template *ServiceGroupTemplate `json:"template,omitempty"`
	// The policies applied to the autoscale configuration.
	Policies []AutoscalePolicy `json:"policies,omitempty"`
	// An optional message providing additional information about the status.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	Error *int32 `json:"error,omitempty"`
	// The status of the response.
	Status ResponseStatus `json:"status"`
}
