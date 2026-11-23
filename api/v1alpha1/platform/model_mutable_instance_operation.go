// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Mutable instance operations.
// +kubebuilder:validation:Enum="set";"add";"del"
type MutableInstanceOperation string

const (
	MutableInstanceOperationSet MutableInstanceOperation = "set"
	MutableInstanceOperationAdd MutableInstanceOperation = "add"
	MutableInstanceOperationDel MutableInstanceOperation = "del"
)
