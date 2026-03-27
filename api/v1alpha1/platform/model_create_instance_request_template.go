// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Template instances.
// An existing instance can be saved as a template. This template is then
// used to create new instances that inherit the exact configuration and
// state the original instance had when the template was created.

type CreateInstanceRequestTemplate struct {
	// (Optional).  Whether the instance needs to run in order to reach template state
	Prepare    *bool                                    `json:"prepare,omitempty"`
	NameOrUUID *CreateInstanceRequestTemplateNameOrUUID `json:"nameOrUUID,omitempty"`
	CreateArgs *CreateInstanceRequestTemplateCreateArgs `json:"create_args,omitempty"`
}
