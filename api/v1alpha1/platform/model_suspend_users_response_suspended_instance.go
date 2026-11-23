// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// An instance that a suspend users operation brought to rest.
type SuspendUsersResponseSuspendedInstance struct {
	// Indicates whether the operation was successful for this item.
	Status ResponseStatus `json:"status"`
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
	// The current state of the instance.
	State InstanceState `json:"state"`
	// The previous state of the instance before the suspend operation was
	// invoked.
	PreviousState InstanceState `json:"previous_state"`
	// What resuming the user will do with this instance. Absent if the
	// instance was already at rest and resuming the user leaves it unchanged.
	OnResume *UserResumeAction `json:"on_resume,omitempty"`
}
