// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Read-Only Memory (ROM) blob to attach to the instance.

type CreateInstanceRequestRom struct {
	// The name of the ROM to use for the instance configuration.
	Name string `json:"name"`
	// (Optional).  The image of the ROM to use for the instance configuration.
	// Mutually exclusive with `files`.
	Image *string `json:"image,omitempty"`
	// (Optional).  Inline files to use as the ROM content.  When specified,
	// the platform creates an EROFS image from the provided files.
	// Mutually exclusive with `image`.
	Files []InlineFile `json:"files,omitempty"`
}
