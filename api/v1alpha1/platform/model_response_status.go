// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The response status of an API request.
// +kubebuilder:validation:Enum="success";"error";"partial_success"
type ResponseStatus string

const (
	ResponseStatusSuccess        ResponseStatus = "success"
	ResponseStatusError          ResponseStatus = "error"
	ResponseStatusPartialSuccess ResponseStatus = "partial_success"
)
