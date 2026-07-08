// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

//	PullPolicy defines when an image should be pulled.
//
// +kubebuilder:validation:Enum=always;if_not_present;never
type PullPolicy string

const (
	PullPolicyALWAYS         PullPolicy = "always"
	PullPolicyIF_NOT_PRESENT PullPolicy = "if_not_present"
	PullPolicyNEVER          PullPolicy = "never"
)
