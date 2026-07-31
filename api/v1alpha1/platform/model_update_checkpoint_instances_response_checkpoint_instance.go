// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

type UpdateCheckpointInstancesResponseCheckpointInstance struct {
	// The UUID of the checkpoint instance that was updated.
	Uuid string `json:"uuid"`
	// The name of the checkpoint instance that was updated.
	Name string `json:"name"`
	// The status of this particular checkpoint instance update operation.
	Status ResponseStatus `json:"status"`
	// (Optional).  The client-provided ID from the request.
	Id *string `json:"id,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
}
