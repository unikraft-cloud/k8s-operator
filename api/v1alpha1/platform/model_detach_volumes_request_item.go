// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A single request of detaching a volume.

type DetachVolumesRequestItem struct {
	// (Optional).  UUID or name of the instance to detach the volume from.
	// If not specified, the volume is detached from all instances.
	From *NameOrUUID `json:"from,omitempty"`
	// The UUID of the volume to detach. Mutually exclusive with name.
	// Exactly one of uuid or name must be provided.
	Uuid string `json:"uuid"`
	// The name of the volume to detach. Mutually exclusive with UUID.
	// Exactly one of uuid or name must be provided.
	Name string `json:"name"`
}
