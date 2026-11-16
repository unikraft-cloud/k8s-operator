// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A single request item to wait for an instance's state.
type WaitInstancesRequestItem struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// The desired state to wait for. Default is `running`.
	State *InstanceState `json:"state,omitempty"`
	// Deprecated: Use `timeout_s` instead. Timeout in milliseconds to
	// wait for the instance to reach the desired state. If `timeout_s` is
	// not set, this value is converted by rounding up to the next full
	// second. A value of -1 selects the platform's maximum API timeout
	// (60000 ms by default); larger values are rejected.
	TimeoutMs *int64 `json:"timeout_ms,omitempty"`
	// Timeout in seconds to wait for the instance to reach the desired
	// state. Every wait is capped at the platform's maximum API timeout
	// (60 seconds by default); larger values are rejected. A value of 0 or
	// -1, or omitting the timeout, waits for that maximum; there is no way
	// to wait indefinitely or to skip the wait. If several instances are
	// waited for in one request, the shortest positive timeout applies to
	// all of them.
	//
	// Reaching the timeout is not an HTTP error: the response carries the
	// message `Operation timed out`, and each instance that has not reached
	// the desired state is reported with its current `state` and an `error`
	// status.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
}
