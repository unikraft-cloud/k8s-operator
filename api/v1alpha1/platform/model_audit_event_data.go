// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The payload of an audit event.
//
// Which fields are present depends on the event type:
//
// - `vm.state_change`
// - `vm.start_failed`
type AuditEventData struct {
	// The state before the transition.
	Prev *string `json:"prev,omitempty"`
	// The state after the transition.
	New *string `json:"new,omitempty"`
	// The state the instance was left in after a failed start.
	State *string `json:"state,omitempty"`
	// The error that failed the start, as an errno name such as `EDQUOT`.
	Error *string `json:"error,omitempty"`
	// Why the instance stopped.
	Stop *AuditStop `json:"stop,omitempty"`
}
