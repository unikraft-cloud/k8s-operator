// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Wait parameters.

type WaitInstanceByUUIDRequestBody struct {
	// The desired state to wait for.  Default is `running`.
	State InstanceState `json:"state"`
	// Deprecated: Use `timeout_s` instead. Timeout in milliseconds to
	// wait for the instance to reach the desired state.  If `timeout_s` is
	// not set, this value is converted by rounding up to the next full
	// second. A value of -1 means to wait indefinitely.
	TimeoutMs *int64 `json:"timeout_ms,omitempty"`
	// Timeout in seconds to wait for the instance to reach the desired
	// state. If the timeout is reached, the request will fail with an
	// error. A value of -1 means to wait indefinitely until the instance
	// reaches the desired state.  No wait performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
}
