// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A single request item to detach a volume from an instance.
type DetachVolumesRequestItem struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// UUID or name of the instance to detach the volume from. If not specified,
	// the volume is detached from all instances.
	From *NameOrUUID `json:"from,omitempty"`
}
