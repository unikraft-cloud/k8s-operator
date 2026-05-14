// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

//	The mutable properties of a template instance that can be updated.
//
// +kubebuilder:validation:Enum=tags;delete_lock;autokill
type MutableTemplateInstanceProperty string

const (
	MutableTemplateInstancePropertyTAGS        MutableTemplateInstanceProperty = "tags"
	MutableTemplateInstancePropertyDELETE_LOCK MutableTemplateInstanceProperty = "delete_lock"
	MutableTemplateInstancePropertyAUTOKILL    MutableTemplateInstanceProperty = "autokill"
)
