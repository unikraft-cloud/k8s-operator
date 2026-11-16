// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Automatic unpin-on-idle configuration of a pinned image. Only populated
// for callers with image manager permissions.
type ImageAutokill struct {
	// The image is unpinned after this many milliseconds of inactivity.
	TimeMs uint64 `json:"time_ms"`
}
