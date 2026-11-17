// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The object an audit event is about.
type AuditObject struct {
	// The kind of object.
	Type AuditObjectType `json:"type"`
	// The object's UUID.
	Uuid string `json:"uuid"`
	// The UUID of the user the object belongs to.
	Owner *string `json:"owner,omitempty"`
	// The tags set on the object.
	Tags []string `json:"tags,omitempty"`
}
