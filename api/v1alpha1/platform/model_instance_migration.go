// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The migration of an instance.
type InstanceMigration struct {
	// Whether the instance is the copy of an incoming migration. Only present
	// on the destination node.
	Incoming *bool `json:"incoming,omitempty"`
	// The migration URL of the destination. Only present on the source node.
	Destination *string `json:"destination,omitempty"`
	// The phase of the migration.
	State InstanceMigrationState `json:"state"`
	// The size of the files the migration copies in bytes, including the holes
	// of sparse files. Only present once the migration knows the files.
	TotalBytes *uint64 `json:"total_bytes,omitempty"`
	// The bytes of these files that have been copied. Only present together
	// with total_bytes.
	CopiedBytes *uint64 `json:"copied_bytes,omitempty"`
}
