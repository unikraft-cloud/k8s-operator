// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Parameters for accepting a migration on the destination node.
type AcceptInstanceMigrationsRequestItem struct {
	// The UUID of the instance the token is valid for. If omitted, the token
	// is valid for any instance.
	Uuid *string `json:"uuid,omitempty"`
	// The lifetime of the token until its first use in seconds, at most 3600.
	// Uses the default of the node if 0 or omitted.
	TtlS *int32 `json:"ttl_s,omitempty"`
}
