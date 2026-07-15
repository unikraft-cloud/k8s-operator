// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// (Optional).  Reference to an existing instance to branch from.
// The instance can be running, stopped, or a template.  If the source
// instance is running, a snapshot will be taken asynchronously and the
// new instance will wait for it to complete before starting.
// Mutually exclusive with `image` and `template`.

type CreateInstanceRequestBranchFrom struct {
	// Mutually exclusive with name.
	Uuid string `json:"uuid"`
	// Mutually exclusive with UUID.
	Name string `json:"name"`
}
