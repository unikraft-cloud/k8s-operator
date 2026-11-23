// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// What an operation was performed on the object.
//
// Open: the set depends on the object type, and instances carry
// `start`, `stop`, `drain`, `suspend` and `restart`.
//
// The values below are the known ones; any other string is accepted too.
type AuditOperationKind string

const (
	AuditOperationKindStart   AuditOperationKind = "start"
	AuditOperationKindStop    AuditOperationKind = "stop"
	AuditOperationKindDrain   AuditOperationKind = "drain"
	AuditOperationKindSuspend AuditOperationKind = "suspend"
	AuditOperationKindRestart AuditOperationKind = "restart"
)
