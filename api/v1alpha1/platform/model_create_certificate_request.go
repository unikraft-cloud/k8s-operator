// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The request message for creating/uploading a new certificate.
type CreateCertificateRequest struct {
	// The name of the resource. If not provided, a random name is generated.
	// The name must be unique within the context of your account.
	Name *string `json:"name,omitempty"`
	// The name of the node to place the resource on. If not provided, the
	// platform picks one. Only accepted on metros that span multiple nodes.
	Node *string `json:"node,omitempty"`
	// The common name (CN) of the certificate.
	//
	// Deprecated: Use `common_name` instead.
	Cn *string `json:"cn,omitempty"`
	// The common name (CN) of the certificate.
	//
	// This must be a fully-qualified domain name (FQDN) ending with a period
	// `.`, e.g. `funky-town.com.`. Exactly one of `cn` or `common_name` must be
	// specified.
	CommonName *string `json:"common_name,omitempty"`
	// The certificate chain in PEM format. Required for user-uploaded certificates.
	Chain string `json:"chain"`
	// The private key in PEM format. Required for user-uploaded certificates.
	Pkey string `json:"pkey"`
}
