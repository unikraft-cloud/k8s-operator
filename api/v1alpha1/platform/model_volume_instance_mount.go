// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// An instance mount of a volume.
type VolumeInstanceMount struct {
	// The UUID of the instance that the volume is mounted in.
	Uuid string `json:"uuid"`
	// The name of the instance that the volume is mounted in.
	Name string `json:"name"`
	// The name of the node the resource is placed on. Only present on metros
	// that span multiple nodes.
	Node *string `json:"node,omitempty"`
	// Whether the volume is mounted read-only or read-write.
	Readonly bool `json:"readonly"`
}
