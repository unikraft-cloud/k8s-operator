// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

//	AdjustmentType defines the type of adjustment to be made in an autoscaling
//	step policy.
//
// +kubebuilder:validation:Enum=change;exact;percentage
type AdjustmentType string

const (
	AdjustmentTypeCHANGE     AdjustmentType = "change"
	AdjustmentTypeEXACT      AdjustmentType = "exact"
	AdjustmentTypePERCENTAGE AdjustmentType = "percentage"
)
