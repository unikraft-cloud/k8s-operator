// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A volume defines a storage which can be attached to the instance. Volumes
// can be used to store persistent data which should remain available even if
// the instance is stopped or restarted.
type InstanceVolume struct {
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
	// The name of the node the resource is placed on. Only present on metros
	// that span multiple nodes.
	Node *string `json:"node,omitempty"`
	// The mount point of the volume in the instance. This is the directory in
	// the instance where the volume will be mounted.
	At string `json:"at"`
	// Whether the volume is read-only or not.
	Readonly bool `json:"readonly"`
}
