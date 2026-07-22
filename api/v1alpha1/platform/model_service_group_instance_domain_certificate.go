// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The certificate associated with the domain.
//
// The certificate is used to secure the domain with TLS/SSL.  If no
// certificate is specified, Unikraft Cloud will automatically generate a
// new certificate for the domain based on Let's Encrypt and seek to
// accomplish a DNS-01 challenge.

type ServiceGroupInstanceDomainCertificate struct {
	// Mutually exclusive with name.
	Uuid string `json:"uuid"`
	// Mutually exclusive with UUID.
	Name string `json:"name"`
}
