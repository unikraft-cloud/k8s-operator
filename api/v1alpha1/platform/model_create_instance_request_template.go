// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Template configuration when creating an instance.
type CreateInstanceRequestTemplate struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// Whether the instance needs to run in order to reach template state
	Prepare *bool `json:"prepare,omitempty"`
	// Configuration parameters to apply when building the new instance from the
	// source template.
	CreateArgs *Instance `json:"create_args,omitempty"`
	// Timeout in seconds for preparing the template before the preparation is
	// aborted. Only applies when `prepare` is set. A value of 0 means no
	// timeout.
	PrepareTimeoutS *int64 `json:"prepare_timeout_s,omitempty"`
	// Automatic delete-on-idle configuration for the template. Only applies
	// when `prepare` is set.
	Autokill *TemplateAutokill `json:"autokill,omitempty"`
}
