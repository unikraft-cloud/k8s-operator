// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// One audit event.
type AuditEvent struct {
	// The type of event.
	Type AuditEventType `json:"type"`
	// When the event was raised.
	Timestamp metav1.Time `json:"timestamp"`
	// The object the event is about. Absent on `gap`, which names no object.
	Object *AuditObject `json:"object,omitempty"`
	// The operation the event belongs to.
	Attribution *AuditAttribution `json:"attribution,omitempty"`
	// The event payload. Its fields depend on `type`.
	Data *AuditEventData `json:"data,omitempty"`
	// How many events were lost. Only on `gap`.
	Dropped *uint64 `json:"dropped,omitempty"`
}
