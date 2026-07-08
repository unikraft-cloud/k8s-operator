// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

type CloneVolumesRequestItemAllOf struct {
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// The name of the new cloned volume.  If not provided, a random name
	// of the form `vol-X` is generated for you, where `X` is a 5 character
	// long random alphanumeric suffix.
	VolName *string `json:"vol_name,omitempty"`
	// The quota policy for the new cloned volume.  If not provided, the quota
	// policy of the source volume is used.
	QuotaPolicy *VolumeQuotaPolicy `json:"quota_policy,omitempty"`
	// A list of tags to assign to the new cloned volume.
	Tags []string `json:"tags,omitempty"`
}
