// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Relay configuration for this interface.

type InstanceNetworkInterfaceRelay struct {
	// UUID of the relay interface.
	Uuid string `json:"uuid"`
	// Name of the relay interface.
	Name string `json:"name"`
	// Whether DNS traffic is relayed through this interface.
	// Defaults to true.
	RelayDns bool `json:"relay_dns"`
}
