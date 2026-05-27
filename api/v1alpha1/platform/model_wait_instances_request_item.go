// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A single wait operation to be applied to an instance.

type WaitInstancesRequestItem struct {
	// The desired state to wait for.  Default is `running`.
	State *InstanceState `json:"state,omitempty"`
	// Timeout in milliseconds to wait for the instance to reach the desired
	// state.  If the timeout is reached, the request will fail with an error.
	// A value of -1 means to wait indefinitely until the instance reaches the
	// desired state.
	TimeoutMs *int64 `json:"timeout_ms,omitempty"`
	// The UUID of the instance to wait for.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to wait for.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
}
