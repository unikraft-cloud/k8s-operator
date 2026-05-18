// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Global VM configuration limits for the user.

type UserVmm struct {
	// Maximum number of vCPUs the user can have assigned to live instances.
	MaxVcpus *int32 `json:"max_vcpus,omitempty"`
	// Maximum amount of memory in MB the user can have assigned to live
	// instances.
	MaxMemoryMb *int32 `json:"max_memory_mb,omitempty"`
}
