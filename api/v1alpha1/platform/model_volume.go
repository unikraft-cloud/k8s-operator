// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// A volume represents a storage device that can be attached to an instance.
type Volume struct {
	// Indicates whether the operation was successful for this item.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information.
	Message *string `json:"message,omitempty"`
	// An optional error code.
	Error *int32 `json:"error,omitempty"`
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
	// The name of the node the resource is placed on. Only present on metros
	// that span multiple nodes.
	Node *string `json:"node,omitempty"`
	// The time the volume was created.
	CreatedAt metav1.Time `json:"created_at"`
	// Current state of the volume.
	State VolumeState `json:"state"`
	// The size of the volume in megabytes.
	SizeMb uint64 `json:"size_mb"`
	// Indicates if the volume will stay alive when the last instance is deleted
	// that this volume is attached to.
	Persistent bool `json:"persistent"`
	// List of instances that this volume is attached to.
	AttachedTo []VolumeInstanceID `json:"attached_to,omitempty"`
	// List of instances that have this volume mounted.
	// This does not apply to template volumes.
	MountedBy []VolumeInstanceMount `json:"mounted_by,omitempty"`
	// The tags associated with the volume.
	// Maximum 16 tags are allowed, and each tag may not be longer than 256 characters.
	Tags []string `json:"tags,omitempty"`
	// Either static or dynamic reservation.
	QuotaPolicy VolumeQuotaPolicy `json:"quota_policy"`
	// If set to true, the volume cannot be deleted.
	DeleteLock *bool `json:"delete_lock,omitempty"`
	// The amount of free space in the volume in megabytes.
	FreeMb *uint32 `json:"free_mb,omitempty"`
	// The filesystem type of this volume.
	// Without custom configuration, this is either `ext4` or `virtiofs`.
	Filesystem *string `json:"filesystem,omitempty"`
	// Host path backing this managed volume.
	// This field is only available for managed volumes and users with
	// appropriate permissions.
	HostPath *string `json:"host_path,omitempty"`
	// The access mode of the volume, controlling volume sharing behavior.
	// Defaults to `rwo` if not specified.
	AccessMode *VolumeAccessMode `json:"access_mode,omitempty"`
	// Guest UID for managed volumes (host_path mode only).
	Uid *uint32 `json:"uid,omitempty"`
	// Guest GID for managed volumes (host_path mode only).
	Gid *uint32 `json:"gid,omitempty"`
	// Optional script arguments that were applied to the custom volume filesystem
	// initialization scripts.
	Args map[string]string `json:"args,omitempty"`
}
