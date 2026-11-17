// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Automatic delete-on-idle configuration for the checkpoint instance.
type CheckpointAutokill struct {
	// Time in milliseconds after the checkpoint was last used for restoring
	// before it is deleted. A value of 0 disables checkpoint autokill.
	TimeMs *uint64 `json:"time_ms,omitempty"`
}
