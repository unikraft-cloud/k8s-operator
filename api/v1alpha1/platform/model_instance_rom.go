// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Read-Only Memory (ROM) blob to attach to the instance.

type InstanceRom struct {
	// The name of the ROM to use for the instance configuration.
	Name *string `json:"name,omitempty"`
	// The image of the ROM to use for the instance configuration.
	Image string `json:"image"`
	// (Optional).  The path at which the ROM should be automatically mounted
	// inside the instance.  When set, the platform mounts the ROM device at
	// the specified path so the guest does not need to mount it manually.
	// When omitted, the ROM is exposed as a raw block device and the guest is
	// responsible for mounting it.
	At *string `json:"at,omitempty"`
}
