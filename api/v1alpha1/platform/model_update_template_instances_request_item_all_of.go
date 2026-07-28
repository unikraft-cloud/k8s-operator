// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import "k8s.io/apimachinery/pkg/runtime"

type UpdateTemplateInstancesRequestItemAllOf struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop *MutableTemplateInstanceProperty `json:"prop,omitempty"`
	// The operation to perform on the property.
	Op *MutableTemplateInstanceOperation `json:"op,omitempty"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "tags": array of strings
	// - For "delete_lock": boolean
	// - For "autokill": object with time_ms field
	Value *runtime.RawExtension `json:"value,omitempty"`
}
