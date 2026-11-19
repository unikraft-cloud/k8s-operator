// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A domain to attach when creating a service group.
type CreateServiceGroupRequestDomain struct {
	// Publicly accessible domain name. If this name ends in a period `.` it must
	// be a valid Full Qualified Domain Name (FQDN), otherwise it will become a
	// subdomain of the target metro.
	Name string `json:"name"`
	// Use an existing certificate for the domain. If this field is
	// specified, the domain must be associated with a valid certificate.
	Certificate *NameOrUUID `json:"certificate,omitempty"`
}
