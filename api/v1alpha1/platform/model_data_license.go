// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// License information (admin only).

type DataLicense struct {
	// The serial number of the license certificate, hex-encoded.
	Serial string `json:"serial"`
	// Whether the license is currently valid.
	Valid bool `json:"valid"`
	// List of enabled features.
	Features []string `json:"features,omitempty"`
}
