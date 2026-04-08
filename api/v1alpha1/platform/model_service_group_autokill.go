// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Automatic delete-on-idle configuration.

type ServiceGroupAutokill struct {
	// Time in milliseconds after the service group becomes empty before it is
	// deleted. A value of 0 disables autokill.
	TimeMs *uint64 `json:"time_ms,omitempty"`
}
