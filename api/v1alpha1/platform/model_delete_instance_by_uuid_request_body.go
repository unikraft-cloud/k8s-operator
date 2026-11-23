// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Parameters for deleting the instance by UUID.
type DeleteInstanceByUUIDRequestBody struct {
	// Timeout in seconds for a graceful stop before the deletion. The platform
	// asks the instance to shut down and waits until it has stopped, so the
	// instance can write its data to its volumes, and the volumes are free when
	// the request returns. If the timeout expires first, the platform stops the
	// instance forcefully. A value of -1 selects the platform's maximum API
	// timeout (60 seconds by default); larger values are rejected.
	//
	// With 0, the default, the platform stops the instance forcefully at once
	// and does not wait. The instance can then lose data that it has not yet
	// written to its volumes, and the volumes stay in use for a short time
	// after the request returns.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
	// Delete immediately without retention. If the instance is already
	// being retained, this will force its deletion. Ignored if retention
	// for instances is not configured.
	DontRetain *bool `json:"dont_retain,omitempty"`
}
