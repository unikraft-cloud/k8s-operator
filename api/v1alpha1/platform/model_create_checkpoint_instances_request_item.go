// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A single checkpoint creation request.
type CreateCheckpointInstancesRequestItem struct {
	// The source instance to create a checkpoint from (by name or UUID).
	From NameOrUUID `json:"from"`
	// The name of the checkpoint. If not provided, a name will be generated.
	Name *string `json:"name,omitempty"`
	// The name of the node to place the checkpoint on. Must be the node of the
	// source instance. Only accepted on metros that span multiple nodes.
	Node *string `json:"node,omitempty"`
	// Timeout in seconds to wait for the checkpoint to be created.
	// No wait performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
	// Tags to associate with the checkpoint.
	Tags []string `json:"tags,omitempty"`
	// Automatic delete-on-idle configuration for the new checkpoint.
	Autokill *CheckpointAutokill `json:"autokill,omitempty"`
}
