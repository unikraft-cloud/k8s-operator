// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// UUID or name of the instance to attach the volume to.

type AttachVolumesRequestItemAttachTo struct {
	// Mutually exclusive with name.
	Uuid string `json:"uuid"`
	// Mutually exclusive with UUID.
	Name string `json:"name"`
}
