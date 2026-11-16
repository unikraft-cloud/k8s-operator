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

// The result of pinning a single image. On success, `uuid` through `tags`
// are set; on failure, only `message` and `error` are set (the image
// being pulled is not otherwise identified in the response).
type PinImagesResponseImage struct {
	// Indicates whether the operation was successful for this item.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information.
	Message *string `json:"message,omitempty"`
	// An optional error code.
	Error *int32 `json:"error,omitempty"`
	// The UUID of the image. Only set on success.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the image. Only set on success.
	Name *string `json:"name,omitempty"`
	// The time the image was created. Only set on success.
	CreatedAt *metav1.Time `json:"created_at,omitempty"`
	// The current state of the image (e.g. `ready`). Only set on success.
	State *string `json:"state,omitempty"`
	// The image URL. Only set on success.
	Url *string `json:"url,omitempty"`
	// Whether the image is pinned and exempt from cache eviction. Only set
	// on success, where it is always `true`.
	Persistent *bool `json:"persistent,omitempty"`
	// The tags associated with the image. Only set on success.
	Tags []string `json:"tags,omitempty"`
}
