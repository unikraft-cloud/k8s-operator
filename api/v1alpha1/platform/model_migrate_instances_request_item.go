// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A single request item to migrate an instance.
type MigrateInstancesRequestItem struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// The destination, as returned by its accept request.
	Destination MigrateInstanceDestination `json:"destination"`
	// Timeout in seconds to wait for the migration to end. No wait if 0
	// (default). Use -1 for the largest possible value.
	//
	// A migration that waits in the queue or for the instance to start or
	// stop has not ended. If the timeout expires first, the request fails
	// with a timeout and the instance reports the state of its migration.
	// The migration goes on.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
}
