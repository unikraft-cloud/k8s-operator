// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Image struct {
	Url string `json:"url"`
	// (Only applies when using global control plane).
	// The metro of the image.
	Metro *string `json:"metro,omitempty"`
	// The time the volume was created.
	CreatedAt   metav1.Time       `json:"created_at"`
	InitrdOrRom bool              `json:"initrd_or_rom"`
	SizeInBytes int64             `json:"size_in_bytes"`
	Args        []string          `json:"args,omitempty"`
	Env         map[string]string `json:"env,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Users       []string          `json:"users,omitempty"`
}
