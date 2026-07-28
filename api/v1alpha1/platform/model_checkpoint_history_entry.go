// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// A checkpoint history entry, representing a single checkpoint in the
// history of an instance.

type CheckpointHistoryEntry struct {
	// The UUID of the checkpoint.
	Uuid string `json:"uuid"`
	// The name of the checkpoint.
	Name string `json:"name"`
	// The time the checkpoint was created.
	CreatedAt metav1.Time `json:"created_at"`
}
