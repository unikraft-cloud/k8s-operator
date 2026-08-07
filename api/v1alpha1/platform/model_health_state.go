// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

//	The health state reported by a single health checker.
//
// +kubebuilder:validation:Enum=unknown;healthy;degraded
type HealthState string

const (
	HealthStateUNKNOWN  HealthState = "unknown"
	HealthStateHEALTHY  HealthState = "healthy"
	HealthStateDEGRADED HealthState = "degraded"
)
