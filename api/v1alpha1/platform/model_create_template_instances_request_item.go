// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// A single template instance to be created.
type CreateTemplateInstancesRequestItem struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// Timeout in seconds to wait for the template instances to be created.
	// No wait is performed for a value of 0, which is the default. The wait is
	// capped at the platform's maximum API timeout (60 seconds by default);
	// larger values are rejected, and a value of -1 waits for that maximum. If
	// several instances are converted in one request, the shortest positive
	// timeout applies to all of them.
	//
	// Reaching the timeout is not an HTTP error: the response carries the
	// message `Operation timed out`, and each instance that has not become a
	// template yet is reported with its current `state` and an `error` status.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
	// Automatic delete-on-idle configuration for the new template.
	Autokill *ItemAutokill `json:"autokill,omitempty"`
}
