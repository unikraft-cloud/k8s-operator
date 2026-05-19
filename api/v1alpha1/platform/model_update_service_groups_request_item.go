// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import "k8s.io/apimachinery/pkg/runtime"

// A single update operation to be applied to a service group

type UpdateServiceGroupsRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop MutableServiceGroupProperty `json:"prop"`
	// The operation to perform.
	Op MutableServiceGroupOperation `json:"op"`
	// The value for the update operation:
	// - For "services": array of Service objects (same as for creation)
	// - For "domains": array of Domain objects (same as for creation)
	// - For "soft_limit": integer (1–65535), must be <= "hard_limit"
	// - For "hard_limit": integer (1–65535), must be >= "soft_limit"
	// - For "autokill": object with time_ms field
	Value *runtime.RawExtension `json:"value,omitempty"`
	// The UUID of the service group to update.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the service group to update.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
}
