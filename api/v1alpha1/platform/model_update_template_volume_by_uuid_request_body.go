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
type UpdateTemplateVolumeByUUIDRequestBodyProp string

const (
	UpdateTemplateVolumeByUUIDRequestBodyPropTags        UpdateTemplateVolumeByUUIDRequestBodyProp = "tags"
	UpdateTemplateVolumeByUUIDRequestBodyPropDelete_lock UpdateTemplateVolumeByUUIDRequestBodyProp = "delete_lock"
)

// The operation to perform.
// +kubebuilder:validation:Enum=set;add;del
type UpdateTemplateVolumeByUUIDRequestBodyOp string

const (
	UpdateTemplateVolumeByUUIDRequestBodyOpSet UpdateTemplateVolumeByUUIDRequestBodyOp = "set"
	UpdateTemplateVolumeByUUIDRequestBodyOpAdd UpdateTemplateVolumeByUUIDRequestBodyOp = "add"
	UpdateTemplateVolumeByUUIDRequestBodyOpDel UpdateTemplateVolumeByUUIDRequestBodyOp = "del"
)

type UpdateTemplateVolumeByUUIDRequestBody struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop UpdateTemplateVolumeByUUIDRequestBodyProp `json:"prop"`
	// The operation to perform.
	Op UpdateTemplateVolumeByUUIDRequestBodyOp `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "tags": array of Strings
	// - For "delete_lock": boolean
	Value *runtime.RawExtension `json:"value,omitempty"`
}
