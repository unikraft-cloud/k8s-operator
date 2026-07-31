// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Per-VM Configuration limits for the user.

type UserVmdb struct {
	// Maximum number of VM instances the user can have at one moment.
	MaxInstances *int32 `json:"max_instances,omitempty"`
	// Minimum amount of memory assigned to a VM in MB.
	MinMemoryMb *int32 `json:"min_memory_mb,omitempty"`
	// Default amount of memory assigned to a VM in MB.
	DefMemoryMb *int32 `json:"def_memory_mb,omitempty"`
	// Maximum amount of memory assigned to a VM in MB.
	MaxMemoryMb *int32 `json:"max_memory_mb,omitempty"`
	// Minimum number of vCPUs assigned to a VM.
	MinVcpus *int32 `json:"min_vcpus,omitempty"`
	// Maximum number of vCPUs assigned to a VM.
	MaxVcpus *int32 `json:"max_vcpus,omitempty"`
}
