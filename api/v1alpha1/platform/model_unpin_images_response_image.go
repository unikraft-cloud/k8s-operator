// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The result of unpinning a single image.
type UnpinImagesResponseImage struct {
	// Indicates whether the operation was successful for this item.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information.
	Message *string `json:"message,omitempty"`
	// An optional error code.
	Error *int32 `json:"error,omitempty"`
	// The UUID of the image.
	Uuid string `json:"uuid"`
	// The name of the image. Only set on success, and only if the image
	// has a name.
	Name *string `json:"name,omitempty"`
}
