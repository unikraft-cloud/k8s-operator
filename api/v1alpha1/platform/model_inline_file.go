// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// An inline file entry represents a single file within an image.

type InlineFile struct {
	// The file path within the image.
	Path string `json:"path"`
	// (Optional).  The encoding of the data field.  Defaults to "text".
	Encoding *InlineDataEncoding `json:"encoding,omitempty"`
	// The file data, encoded according to the encoding field.
	Data string `json:"data"`
}
