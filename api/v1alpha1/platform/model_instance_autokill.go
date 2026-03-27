// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Automatic delete-on-idle/request-limit configuration.
// Not used for template instances.

type InstanceAutokill struct {
	// Time in milliseconds after the instance was stopped before it is deleted.
	// A value of 0 disables time-based autokill.
	TimeMs *uint64 `json:"time_ms,omitempty"`
	// Maximum number of requests/connections the instance serves before it is
	// deleted. A value of 0 disables request-based autokill.
	NumRequests *uint32 `json:"num_requests,omitempty"`
}
