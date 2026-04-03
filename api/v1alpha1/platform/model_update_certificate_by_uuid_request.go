// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The request message for updating a certificate by its UUID.

type UpdateCertificateByUUIDRequest struct {
	// The UUID of the certificate to update.
	Uuid *string `json:"uuid,omitempty"`
	// The new certificate chain.
	//
	// This is the public chain of the certificate in PEM format. The chain
	// should include the certificate and any intermediate certificates.
	Chain string `json:"chain"`
	// The new private key.
	//
	// This is the private key of the certificate in PEM format. The private
	// key must match the public key in the certificate chain.
	Pkey string `json:"pkey"`
}
