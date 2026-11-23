// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The service group configuration when creating an instance. If no existing
// (persistent) service group is specified via its identifier, a new
// (ephemeral) service group can be created by specifying the services it
// should expose. A service defines the configuration settings of an exposed
// port by the instance. A service is a combination of a public port, an
// internal port, and a set of handlers that define how the service will handle
// incoming connections.
type CreateInstanceRequestServiceGroup struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the resource.
	Name      *string                       `json:"name,omitempty"`
	Domains   []CreateInstanceRequestDomain `json:"domains,omitempty"`
	Services  []Service                     `json:"services,omitempty"`
	SoftLimit *uint32                       `json:"soft_limit,omitempty"`
	HardLimit *uint32                       `json:"hard_limit,omitempty"`
}
