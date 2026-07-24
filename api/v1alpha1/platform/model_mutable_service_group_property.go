// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

//	The mutable properties of a service group.
//
// +kubebuilder:validation:Enum=services;domains;soft_limit;hard_limit;autokill
type MutableServiceGroupProperty string

const (
	MutableServiceGroupPropertySERVICES   MutableServiceGroupProperty = "services"
	MutableServiceGroupPropertyDOMAINS    MutableServiceGroupProperty = "domains"
	MutableServiceGroupPropertySOFT_LIMIT MutableServiceGroupProperty = "soft_limit"
	MutableServiceGroupPropertyHARD_LIMIT MutableServiceGroupProperty = "hard_limit"
	MutableServiceGroupPropertyAUTOKILL   MutableServiceGroupProperty = "autokill"
)
