// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Autoscale configuration limits for the user.

type UserAutoscale struct {
	// Minimum size of an autoscale group.
	MinSize *int32 `json:"min_size,omitempty"`
	// Maximum size of an autoscale group.
	MaxSize *int32 `json:"max_size,omitempty"`
}
