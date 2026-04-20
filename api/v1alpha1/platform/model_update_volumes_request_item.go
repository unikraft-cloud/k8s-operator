// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import "k8s.io/apimachinery/pkg/runtime"

// A single request item for updating a volume.
// The property to modify.
// +kubebuilder:validation:Enum=size_mb;tags;quota_policy;delete_lock
type UpdateVolumesRequestItemProp string

const (
	UpdateVolumesRequestItemPropSize_mb      UpdateVolumesRequestItemProp = "size_mb"
	UpdateVolumesRequestItemPropTags         UpdateVolumesRequestItemProp = "tags"
	UpdateVolumesRequestItemPropQuota_policy UpdateVolumesRequestItemProp = "quota_policy"
	UpdateVolumesRequestItemPropDelete_lock  UpdateVolumesRequestItemProp = "delete_lock"
)

// The operation to perform.
// +kubebuilder:validation:Enum=set;add;del
type UpdateVolumesRequestItemOp string

const (
	UpdateVolumesRequestItemOpSet UpdateVolumesRequestItemOp = "set"
	UpdateVolumesRequestItemOpAdd UpdateVolumesRequestItemOp = "add"
	UpdateVolumesRequestItemOpDel UpdateVolumesRequestItemOp = "del"
)

type UpdateVolumesRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop UpdateVolumesRequestItemProp `json:"prop"`
	// The operation to perform.
	Op UpdateVolumesRequestItemOp `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "size_mb": unsigned integer
	// - For "quota_policy": "static" or "dynamic"
	// - For "tags": array of Strings
	// - For "delete_lock": boolean
	Value *runtime.RawExtension `json:"value,omitempty"`
	// The UUID of the volume to update.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to update.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
}
