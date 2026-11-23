// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Mutable service group properties.
// +kubebuilder:validation:Enum="services";"domains";"soft_limit";"hard_limit";"autokill"
type MutableServiceGroupProperty string

const (
	MutableServiceGroupPropertyServices  MutableServiceGroupProperty = "services"
	MutableServiceGroupPropertyDomains   MutableServiceGroupProperty = "domains"
	MutableServiceGroupPropertySoftLimit MutableServiceGroupProperty = "soft_limit"
	MutableServiceGroupPropertyHardLimit MutableServiceGroupProperty = "hard_limit"
	MutableServiceGroupPropertyAutokill  MutableServiceGroupProperty = "autokill"
)
