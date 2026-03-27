// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Template-specific automatic delete-on-idle configuration.
// Not used for non-template instances.

type InstanceTemplateAutokill struct {
	// Time in milliseconds after the template was last used for cloning before
	// it is deleted. A value of 0 disables template autokill.
	TimeMs *uint64 `json:"time_ms,omitempty"`
}
