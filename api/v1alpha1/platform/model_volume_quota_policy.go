// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Quota policy for a volume.
// +kubebuilder:validation:Enum="static";"dynamic"
type VolumeQuotaPolicy string

const (
	VolumeQuotaPolicyStatic  VolumeQuotaPolicy = "static"
	VolumeQuotaPolicyDynamic VolumeQuotaPolicy = "dynamic"
)
