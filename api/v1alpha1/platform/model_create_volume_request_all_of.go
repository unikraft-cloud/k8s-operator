// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

type CreateVolumeRequestAllOf struct {
	// The name of the volume.
	//
	// This is a human-readable name that can be used to identify the volume.
	// The name must be unique within the context of your account.  If no name is
	// specified, a random name of the form `vol-X` is generated for you, where
	// `X` is a 5 character long random alphanumeric suffix..  The name can also
	// be used to identify the volume in API calls.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// Quota policy for the volume.
	QuotaPolicy *VolumeQuotaPolicy `json:"quota_policy,omitempty"`
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
	// The access mode of the volume, controlling volume sharing behavior.
	// Defaults to `rwo` if not specified.
	AccessMode *VolumeAccessMode `json:"access_mode,omitempty"`
}
