// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Network interface configuration for a new instance.

type CreateInstanceRequestNetworkInterface struct {
	// The interface name. If omitted, Unikraft Cloud generates one as
	// <instance-name>-ethX, falling back to eth-<suffix> when the
	// instance name is too long.
	Name *string `json:"name,omitempty"`
	// The TAP device to attach the interface to. Provide it together
	// with ip.
	TapName *string `json:"tap_name,omitempty"`
	// The interface IP address in CIDR notation. Provide it together
	// with tap_name to bring your own interface.
	Ip *string `json:"ip,omitempty"`
	// Whether the guest configures the interface itself. Defaults to true.
	Autoconfig *bool                                       `json:"autoconfig,omitempty"`
	Relay      *CreateInstanceRequestNetworkInterfaceRelay `json:"relay,omitempty"`
}
