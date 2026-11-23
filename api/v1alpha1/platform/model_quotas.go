// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Quotas with per-item response envelope fields merged in.
type Quotas struct {
	// Indicates whether the operation was successful for this item.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information.
	Message *string `json:"message,omitempty"`
	// An optional error code.
	Error *int32 `json:"error,omitempty"`
	// The UUID of the quota.
	Uuid string `json:"uuid"`
	// Used quota.
	Used QuotasStats `json:"used"`
	// Configured quota limits.
	Hard QuotasStats `json:"hard"`
	// Additional limits.
	Limits QuotasLimits `json:"limits"`
}
