// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Relay configuration for this interface.

type CreateInstanceRequestNetworkInterfaceRelay struct {
	// Whether the relay forwards DNS requests. Set to false
	// to let the default DNS server handle them instead.
	// Defaults to true.
	RelayDns *bool `json:"relay_dns,omitempty"`
	// UUID of the existing interface to relay through.
	// Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// Name of the existing interface to relay through.
	// Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
}
