// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

//	The mutable properties of a volume.
//
// +kubebuilder:validation:Enum=size_mb;tags;quota_policy;delete_lock
type MutableVolumeProperty string

const (
	MutableVolumePropertySIZE_MB      MutableVolumeProperty = "size_mb"
	MutableVolumePropertyTAGS         MutableVolumeProperty = "tags"
	MutableVolumePropertyQUOTA_POLICY MutableVolumeProperty = "quota_policy"
	MutableVolumePropertyDELETE_LOCK  MutableVolumeProperty = "delete_lock"
)
