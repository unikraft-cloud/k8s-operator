// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The request message for creating one or more template volumes.

type CreateTemplateVolumesRequest struct {
	// The list of IDs of the volumes that will be converted into templates.
	// Each ID can be either a UUID or a name.
	Ids []NameOrUUID `json:"ids"`
}
