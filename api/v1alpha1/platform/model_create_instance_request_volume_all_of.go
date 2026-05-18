// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

type CreateInstanceRequestVolumeAllOf struct {
	// The UUID of an existing volume.
	//
	// If this is the only specified field, then it will look up an existing
	// volume by this UUID.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume.
	//
	// If this is the only specified field, then it will look up an existing
	// volume by this name.  If the volume does not exist, the request will
	// fail.  If a new volume is intended to be created, then this field must be
	// specified along with the mount point in the instance and a provisioning
	// source (size_mb or host_path).
	Name *string `json:"name,omitempty"`
	// The mount point for the volume in the instance.
	At *string `json:"at,omitempty"`
	// Whether the volume is read-only.
	//
	// If this field is set to true, the volume will be mounted as read-only in
	// the instance.  This field is optional and defaults to false and is only
	// applicable when using an existing volume.
	Readonly *bool `json:"readonly,omitempty"`
	// Quota policy for the volume.
	QuotaPolicy *string `json:"quota_policy,omitempty"`
	// Filesystem type to format or configure.
	// Without custom configuration, this is either `ext4` or `virtiofs`.
	Filesystem *string `json:"filesystem,omitempty"`
	// Tags to assign to the new volume.
	Tags []string `json:"tags,omitempty"`
	// Guest UID for managed volumes (host_path mode only).
	Uid *uint32 `json:"uid,omitempty"`
	// Guest GID for managed volumes (host_path mode only).
	Gid *uint32 `json:"gid,omitempty"`
	// Script arguments passed to volume initialization scripts.
	Args map[string]string `json:"args,omitempty"`
}
