// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Automatically unpin the image after a period of inactivity.
type PinImageRequestItemAutokill struct {
	// Automatically unpin the image after this many milliseconds of
	// inactivity. `0` (the default) disables this.
	TimeMs uint64 `json:"time_ms"`
}
