// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The request message for updating one or more certificate(s) by their
// UUID(s) or name(s).

type UpdateCertificatesRequest struct {
	// A list of update operations to apply to certificates.
	Body []UpdateCertificatesRequestItem `json:"body"`
}
