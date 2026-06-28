// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// +kubebuilder:validation:Enum=root;override_edns_blacklist;developer;volume_manager;override_vm_priority
type UserPermission string

const (
	UserPermissionROOT                    UserPermission = "root"
	UserPermissionOVERRIDE_EDNS_BLACKLIST UserPermission = "override_edns_blacklist"
	UserPermissionDEVELOPER               UserPermission = "developer"
	UserPermissionVOLUME_MANAGER          UserPermission = "volume_manager"
	UserPermissionOVERRIDE_VM_PRIORITY    UserPermission = "override_vm_priority"
)
