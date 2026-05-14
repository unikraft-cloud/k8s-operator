// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Network configuration limits for the user.

type UserNet struct {
	// Maximum number of service groups the user can have at one moment.
	MaxServiceGroups *int32 `json:"max_service_groups,omitempty"`
	// Maximum number of services across all service groups the user can have
	// at one moment.
	MaxServices *int32 `json:"max_services,omitempty"`
	// Maximum number of TLS certificates the user can have at one moment.
	MaxCertificates *int32 `json:"max_certificates,omitempty"`
}
