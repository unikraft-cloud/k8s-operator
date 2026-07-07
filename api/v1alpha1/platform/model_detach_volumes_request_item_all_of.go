// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

type DetachVolumesRequestItemAllOf struct {
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// (Optional).  UUID or name of the instance to detach the volume from.
	// If not specified, the volume is detached from all instances.
	From *NameOrUUID `json:"from,omitempty"`
}
