// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Storage configuration limits for the user.

type UserStor struct {
	// Maximum number of volumes the user can have at one moment.
	MaxVolumes *int32 `json:"max_volumes,omitempty"`
	// Minimum size of a volume in MB.
	MinVolumeMb *int32 `json:"min_volume_mb,omitempty"`
	// Maximum size of a volume in MB.
	MaxVolumeMb *int32 `json:"max_volume_mb,omitempty"`
	// Maximum total size of all volumes in MB.
	MaxTotalVolumeMb *int32 `json:"max_total_volume_mb,omitempty"`
}
