// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Records the current restart attempt of an instance.

type InstanceRestartAttempt struct {
	// Current restart attempt number. This is incremented each time the instance
	// is restarted automatically by the platform.
	Attempt uint32 `json:"attempt"`
	// Timestamp of the next scheduled restart attempt.
	NextAt *metav1.Time `json:"next_at,omitempty"`
}
