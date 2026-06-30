// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

type User struct {
	// The UUID of the user.
	Uuid string `json:"uuid"`
	// The name of the user.
	Name string `json:"name"`
	// Authentication token(s) associated with the user.
	AuthToken []string `json:"auth_token"`
	// The permission level of the user.
	Permissions []UserPermission `json:"permissions,omitempty"`
	// The user ID (UID) on the host system.
	Uid *uint32 `json:"uid,omitempty"`
	// Whether the user account is disabled.
	Disabled  *bool          `json:"disabled,omitempty"`
	Vmdb      *UserVmdb      `json:"vmdb,omitempty"`
	Net       *UserNet       `json:"net,omitempty"`
	Vmm       *UserVmm       `json:"vmm,omitempty"`
	Stor      *UserStor      `json:"stor,omitempty"`
	Autoscale *UserAutoscale `json:"autoscale,omitempty"`
}
