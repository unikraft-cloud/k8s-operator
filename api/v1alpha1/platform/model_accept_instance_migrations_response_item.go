// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Per-item result for an accept migrations operation.
type AcceptInstanceMigrationsResponseItem struct {
	// Indicates whether the operation was successful for this item.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information.
	Message *string `json:"message,omitempty"`
	// An optional error code.
	Error *int32 `json:"error,omitempty"`
	// The UUID of the instance the token is valid for, if one was given.
	Uuid *string `json:"uuid,omitempty"`
	// The migration URL of this node, to pass to the migrate request.
	Url string `json:"url"`
	// The token for the migration, to pass to the migrate request.
	Token string `json:"token"`
	// The lifetime of the token until its first use in seconds.
	TtlS int32 `json:"ttl_s"`
}
