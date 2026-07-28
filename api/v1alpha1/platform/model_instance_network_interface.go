// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// An instance network interface.

type InstanceNetworkInterface struct {
	// The UUID of the network interface. This is a unique identifier for the
	// network interface that is generated when the instance is created.
	Uuid string `json:"uuid"`
	// The private IP address of the network interface. This is the internal IP
	// address that is used for communication between instances within the same
	// network.
	PrivateIp string `json:"private_ip"`
	// The MAC address of the network interface.
	Mac string `json:"mac"`
	// The interface name. If omitted, Unikraft Cloud generates one as
	// <instance-name>-ethX, falling back to eth-<suffix> when the
	// instance name is too long.
	Name *string `json:"name,omitempty"`
	// The TAP device to attach the interface. Provide it together with ip.
	TapName *string `json:"tap_name,omitempty"`
	// Whether the interface is automatically configured inside the guest
	// (IP address, routes, etc.).  When absent or true, autoconfiguration
	// is enabled.  Present and false when the guest is expected to
	// configure the interface manually.
	Autoconfig *bool                          `json:"autoconfig,omitempty"`
	Relay      *InstanceNetworkInterfaceRelay `json:"relay,omitempty"`
}
