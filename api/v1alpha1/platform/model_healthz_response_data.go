// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The response data for this request.

type HealthzResponseData struct {
	// The health state of each registered checker, keyed by checker name.
	// Valid keys are "images", "systemd", and "user-defined"; a checker's
	// key is only present if it is enabled. Checkers report only their
	// aggregate state; per-check detail (e.g. which default image is
	// missing, or which user-defined script failed) is not exposed here.
	Checks   map[string]HealthState `json:"checks,omitempty"`
	Versions map[string]string      `json:"versions,omitempty"`
	License  *DataLicense           `json:"license,omitempty"`
}
