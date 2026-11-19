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

// An image representing a VM which can be deployed on Unikraft Cloud.
type Image struct {
	// The UUID of the image. Not returned by the image store.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the image, if it has one.
	Name *string `json:"name,omitempty"`
	// The current state of the image (e.g. `ready`). Not returned by the image
	// store.
	State *string `json:"state,omitempty"`
	// The image URL.
	Url string `json:"url"`
	// The time the image was created.
	CreatedAt metav1.Time `json:"created_at"`
	// Whether the image is an initrd or ROM.
	InitrdOrRom bool `json:"initrd_or_rom"`
	// The size of the image in bytes.
	SizeInBytes int64 `json:"size_in_bytes"`
	// Command-line arguments for the image.
	Args []string `json:"args,omitempty"`
	// Environment variables for the image.
	Env map[string]string `json:"env,omitempty"`
	// Tags associated with the image.
	Tags []string `json:"tags,omitempty"`
	// Users associated with the image.
	Users []string `json:"users,omitempty"`
	// Whether the image is pinned and exempt from cache eviction. Only
	// populated (and only ever `true`) for callers with image manager
	// permissions; omitted otherwise, including when the image is not pinned.
	Persistent *bool `json:"persistent,omitempty"`
	// Automatic unpin-on-idle configuration of the image. Only populated for
	// callers with image manager permissions, on a pinned image that has an
	// autokill time.
	Autokill *ImageAutokill `json:"autokill,omitempty"`
}
