// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A volume defines a storage volume that can be attached to the instance.
type CreateInstanceRequestVolume struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the resource.
	Name        *string           `json:"name,omitempty"`
	SizeMb      *uint64           `json:"size_mb,omitempty"`
	HostPath    *string           `json:"host_path,omitempty"`
	At          string            `json:"at"`
	Readonly    *bool             `json:"readonly,omitempty"`
	QuotaPolicy *string           `json:"quota_policy,omitempty"`
	Filesystem  *string           `json:"filesystem,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Uid         *uint32           `json:"uid,omitempty"`
	Gid         *uint32           `json:"gid,omitempty"`
	AccessMode  *VolumeAccessMode `json:"access_mode,omitempty"`
	Args        map[string]string `json:"args,omitempty"`
}
