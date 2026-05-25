// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

//	The operations available on a template instance's properties.
//
// +kubebuilder:validation:Enum=set;add;del
type MutableTemplateInstanceOperation string

const (
	MutableTemplateInstanceOperationSET MutableTemplateInstanceOperation = "set"
	MutableTemplateInstanceOperationADD MutableTemplateInstanceOperation = "add"
	MutableTemplateInstanceOperationDEL MutableTemplateInstanceOperation = "del"
)
