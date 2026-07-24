// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

//	The mutable properties of an instance that can be updated.
//
// +kubebuilder:validation:Enum=image;args;env;memory_mb;vcpus;scale_to_zero;tags;delete_lock;schedules;autokill;hostname;roms;dependencies;sched_priority;plugins
type MutableInstanceProperty string

const (
	MutableInstancePropertyIMAGE          MutableInstanceProperty = "image"
	MutableInstancePropertyARGS           MutableInstanceProperty = "args"
	MutableInstancePropertyENV            MutableInstanceProperty = "env"
	MutableInstancePropertyMEMORY_MB      MutableInstanceProperty = "memory_mb"
	MutableInstancePropertyVCPUS          MutableInstanceProperty = "vcpus"
	MutableInstancePropertySCALE_TO_ZERO  MutableInstanceProperty = "scale_to_zero"
	MutableInstancePropertyTAGS           MutableInstanceProperty = "tags"
	MutableInstancePropertyDELETE_LOCK    MutableInstanceProperty = "delete_lock"
	MutableInstancePropertySCHEDULES      MutableInstanceProperty = "schedules"
	MutableInstancePropertyAUTOKILL       MutableInstanceProperty = "autokill"
	MutableInstancePropertyHOSTNAME       MutableInstanceProperty = "hostname"
	MutableInstancePropertyROMS           MutableInstanceProperty = "roms"
	MutableInstancePropertyDEPENDENCIES   MutableInstanceProperty = "dependencies"
	MutableInstancePropertySCHED_PRIORITY MutableInstanceProperty = "sched_priority"
	MutableInstancePropertyPLUGINS        MutableInstanceProperty = "plugins"
)
