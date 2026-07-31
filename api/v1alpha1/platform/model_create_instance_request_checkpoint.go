// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// (Optional).  Reference to an existing checkpoint to create the instance
// from.  The checkpoint must be in the `checkpoint` state.  The new instance
// will be created with the same configuration and state as the checkpoint.
// Mutually exclusive with `image`, `template`, and `branch_from`.

type CreateInstanceRequestCheckpoint struct {
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The name of the resource.
	Name string `json:"name"`
}
