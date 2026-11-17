// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The service group configuration for the instance. This is a reference to the
// service group that the instance is part of. The service group defines the
// services (e.g. ports, connection handling) that the instance exposes and how
// they are configured.
type InstanceServiceGroup struct {
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
	// The name of the node the resource is placed on. Only present on metros
	// that span multiple nodes.
	Node *string `json:"node,omitempty"`
	// The domain configuration for the service group.
	Domains []ServiceGroupInstanceDomain `json:"domains,omitempty"`
}
