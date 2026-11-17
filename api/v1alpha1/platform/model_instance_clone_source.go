// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The instance, template, or checkpoint that an instance was cloned from.
type InstanceCloneSource struct {
	// The UUID of the source.
	Uuid string `json:"uuid"`
	// The name of the source.
	Name *string `json:"name,omitempty"`
	// The name of the node the source is placed on.
	Node *string `json:"node,omitempty"`
}
