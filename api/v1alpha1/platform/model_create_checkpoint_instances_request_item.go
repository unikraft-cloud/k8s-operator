// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A single checkpoint creation request.

type CreateCheckpointInstancesRequestItem struct {
	From CreateCheckpointInstancesRequestItemFrom `json:"from"`
	// (Optional).  The name of the checkpoint.
	// If not provided, a name will be generated.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// Timeout in seconds to wait for the checkpoint to be created.
	// No wait performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
}
