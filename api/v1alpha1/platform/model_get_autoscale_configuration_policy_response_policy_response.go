// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Per-item result for a get autoscale configuration policy operation.
type GetAutoscaleConfigurationPolicyResponsePolicyResponse struct {
	// The policy which was retrieved by the request.
	Policy AutoscalePolicy `json:"policy"`
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information about the status.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	Error *int32 `json:"error,omitempty"`
}
