// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import "k8s.io/apimachinery/pkg/runtime"

// A queued property change awaiting application (typically on next restart).
// The status of this update.
// +kubebuilder:validation:Enum=pending;failed
type InstancePendingUpdateStatus string

const (
	InstancePendingUpdateStatusPending InstancePendingUpdateStatus = "pending"
	InstancePendingUpdateStatusFailed  InstancePendingUpdateStatus = "failed"
)

type InstancePendingUpdate struct {
	// The property being updated.
	Prop MutableInstanceProperty `json:"prop"`
	// The patch operation type.
	Op MutableInstanceOperation `json:"op"`
	// The new value for the property.  Type depends on the property being
	// updated.
	Value runtime.RawExtension `json:"value"`
	// The status of this update.
	Status InstancePendingUpdateStatus `json:"status"`
	// Error message.  Only present when status is "failed".
	Error *string `json:"error,omitempty"`
}
