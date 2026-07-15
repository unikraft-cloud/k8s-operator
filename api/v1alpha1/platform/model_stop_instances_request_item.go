// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A single request item to stop an instance.

type StopInstancesRequestItem struct {
	// Whether to immediately force stop the instance.
	Force *bool `json:"force,omitempty"`
	// Timeout for draining connections in milliseconds.  The instance does not
	// receive new connections in the draining phase.  The instance is stopped
	// when the last connection has been closed or the timeout expired.  The
	// maximum timeout may vary.  Use -1 for the largest possible value.
	//
	// Note: This endpoint does not block.  Use the wait endpoint for the
	// instance to reach the stopped state.
	DrainTimeoutMs *uint64 `json:"drain_timeout_ms,omitempty"`
	// Whether to perform a quick shutdown.  This flag is
	// overridden by force.
	Quick *bool `json:"quick,omitempty"`
	// Only stop the instance if it is in this state.
	Ifstate *string `json:"ifstate,omitempty"`
	// If set, forces the VMM to shutdown immediately and generate a coredump.
	// Can only be used in conjunction with force.
	Dump *bool `json:"dump,omitempty"`
	// The UUID of the instance to stop.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to stop.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
}
