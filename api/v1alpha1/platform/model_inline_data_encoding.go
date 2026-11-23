// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Encoding type for inline file data.
// +kubebuilder:validation:Enum="text";"base64"
type InlineDataEncoding string

const (
	InlineDataEncodingText   InlineDataEncoding = "text"
	InlineDataEncodingBase64 InlineDataEncoding = "base64"
)
