// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Why an instance stopped.
type AuditStop struct {
	// The origins that contributed to the stop.
	Reason []string `json:"reason,omitempty"`
	// The kernel stop code.
	Code *uint32 `json:"code,omitempty"`
	// The stop cause.
	Cause *string `json:"cause,omitempty"`
}
