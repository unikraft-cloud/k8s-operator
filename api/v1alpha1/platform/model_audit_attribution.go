// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The operation an event belongs to.
//
// Events sharing an `operation` describe one overarching action, which is what
// makes a state transition attributable to the thing that caused it rather than
// only observable after the fact.
type AuditAttribution struct {
	// UUID shared by every event belonging to the same operation.
	Operation *string `json:"operation,omitempty"`
	// What was performed on the object.
	Kind *AuditOperationKind `json:"kind,omitempty"`
	// Whether this event was raised by the operation or observed after it.
	Trigger AuditTrigger `json:"trigger"`
	// What caused the operation.
	Origin AuditOrigin `json:"origin"`
	// The user that caused the operation, when one did.
	User *string `json:"user,omitempty"`
}
