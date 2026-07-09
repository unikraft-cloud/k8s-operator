// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

//	The operations available on a checkpoint instance's properties.
//
// +kubebuilder:validation:Enum=set;add;del
type MutableCheckpointInstanceOperation string

const (
	MutableCheckpointInstanceOperationSET MutableCheckpointInstanceOperation = "set"
	MutableCheckpointInstanceOperationADD MutableCheckpointInstanceOperation = "add"
	MutableCheckpointInstanceOperationDEL MutableCheckpointInstanceOperation = "del"
)
