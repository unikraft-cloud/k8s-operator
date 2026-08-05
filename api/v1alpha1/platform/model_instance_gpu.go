// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A GPU attached to the instance.

type InstanceGpu struct {
	// The UUID of the GPU.
	Uuid string `json:"uuid"`
	// The GPU model, given as its PCI vendor and device ID in the form
	// `<vendor>:<device>`.
	Model string `json:"model"`
}
