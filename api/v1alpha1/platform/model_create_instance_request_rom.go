// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A ROM to attach when creating an instance.
type CreateInstanceRequestRom struct {
	// The name of the ROM to use for the instance configuration.
	Name string `json:"name"`
	// The image of the ROM to use for the instance configuration. Mutually
	// exclusive with `files`. Accepts either a plain image reference string
	// (`"nginx:latest"`) or an object carrying additional pull configuration
	// (`{"url": "nginx:latest", "pull_policy": "always"}`).
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	Image *ImageSource `json:"image,omitempty"`
	// The path at which the ROM should be automatically mounted inside the
	// instance. When set, the platform mounts the ROM device at the specified
	// path so the guest does not need to mount it manually. When omitted, the
	// ROM is exposed as a raw block device and the guest is responsible for
	// mounting it.
	At *string `json:"at,omitempty"`
	// Inline files to use as the ROM content. When specified, the platform
	// creates an EROFS image from the provided files. Mutually exclusive with
	// `image`.
	Files []InlineFile `json:"files,omitempty"`
}
