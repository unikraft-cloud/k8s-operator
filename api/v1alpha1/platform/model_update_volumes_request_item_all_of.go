// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import "k8s.io/apimachinery/pkg/runtime"

// The property to modify.
// +kubebuilder:validation:Enum=size_mb;tags;quota_policy;delete_lock
type UpdateVolumesRequestItemAllOfProp string

const (
	UpdateVolumesRequestItemAllOfPropSize_mb      UpdateVolumesRequestItemAllOfProp = "size_mb"
	UpdateVolumesRequestItemAllOfPropTags         UpdateVolumesRequestItemAllOfProp = "tags"
	UpdateVolumesRequestItemAllOfPropQuota_policy UpdateVolumesRequestItemAllOfProp = "quota_policy"
	UpdateVolumesRequestItemAllOfPropDelete_lock  UpdateVolumesRequestItemAllOfProp = "delete_lock"
)

// The operation to perform.
// +kubebuilder:validation:Enum=set;add;del
type UpdateVolumesRequestItemAllOfOp string

const (
	UpdateVolumesRequestItemAllOfOpSet UpdateVolumesRequestItemAllOfOp = "set"
	UpdateVolumesRequestItemAllOfOpAdd UpdateVolumesRequestItemAllOfOp = "add"
	UpdateVolumesRequestItemAllOfOpDel UpdateVolumesRequestItemAllOfOp = "del"
)

type UpdateVolumesRequestItemAllOf struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop *UpdateVolumesRequestItemAllOfProp `json:"prop,omitempty"`
	// The operation to perform.
	Op *UpdateVolumesRequestItemAllOfOp `json:"op,omitempty"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "size_mb": unsigned integer
	// - For "quota_policy": "static" or "dynamic"
	// - For "tags": array of Strings
	// - For "delete_lock": boolean
	Value *runtime.RawExtension `json:"value,omitempty"`
}
