// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The request message for creating a volume.
type CreateVolumeRequest struct {
	// The name of the resource. If not provided, a random name is generated.
	// The name must be unique within the context of your account.
	Name *string `json:"name,omitempty"`
	// The name of the node to place the resource on. If not provided, the
	// platform picks one. Only accepted on metros that span multiple nodes.
	Node        *string            `json:"node,omitempty"`
	SizeMb      *uint64            `json:"size_mb,omitempty"`
	HostPath    *string            `json:"host_path,omitempty"`
	Template    *NameOrUUID        `json:"template,omitempty"`
	QuotaPolicy *VolumeQuotaPolicy `json:"quota_policy,omitempty"`
	Filesystem  *string            `json:"filesystem,omitempty"`
	Tags        []string           `json:"tags,omitempty"`
	Uid         *uint32            `json:"uid,omitempty"`
	Gid         *uint32            `json:"gid,omitempty"`
	AccessMode  *VolumeAccessMode  `json:"access_mode,omitempty"`
	Args        map[string]string  `json:"args,omitempty"`
}
