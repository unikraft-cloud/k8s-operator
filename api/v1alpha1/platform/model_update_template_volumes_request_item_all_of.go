// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import "k8s.io/apimachinery/pkg/runtime"

// The property to modify.
// +kubebuilder:validation:Enum=tags;delete_lock
type UpdateTemplateVolumesRequestItemAllOfProp string

const (
	UpdateTemplateVolumesRequestItemAllOfPropTags        UpdateTemplateVolumesRequestItemAllOfProp = "tags"
	UpdateTemplateVolumesRequestItemAllOfPropDelete_lock UpdateTemplateVolumesRequestItemAllOfProp = "delete_lock"
)

// The operation to perform.
// +kubebuilder:validation:Enum=set;add;del
type UpdateTemplateVolumesRequestItemAllOfOp string

const (
	UpdateTemplateVolumesRequestItemAllOfOpSet UpdateTemplateVolumesRequestItemAllOfOp = "set"
	UpdateTemplateVolumesRequestItemAllOfOpAdd UpdateTemplateVolumesRequestItemAllOfOp = "add"
	UpdateTemplateVolumesRequestItemAllOfOpDel UpdateTemplateVolumesRequestItemAllOfOp = "del"
)

type UpdateTemplateVolumesRequestItemAllOf struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop *UpdateTemplateVolumesRequestItemAllOfProp `json:"prop,omitempty"`
	// The operation to perform.
	Op *UpdateTemplateVolumesRequestItemAllOfOp `json:"op,omitempty"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "tags": array of Strings
	// - For "delete_lock": boolean
	Value *runtime.RawExtension `json:"value,omitempty"`
}
