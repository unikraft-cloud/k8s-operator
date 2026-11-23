// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The mutable properties of an instance that can be updated.
// +kubebuilder:validation:Enum="image";"args";"env";"memory_mb";"vcpus";"scale_to_zero";"tags";"delete_lock";"schedules";"autokill";"hostname";"roms";"dependencies";"sched_priority";"plugins";"annotations"
type MutableInstanceProperty string

const (
	MutableInstancePropertyImage         MutableInstanceProperty = "image"
	MutableInstancePropertyArgs          MutableInstanceProperty = "args"
	MutableInstancePropertyEnv           MutableInstanceProperty = "env"
	MutableInstancePropertyMemoryMb      MutableInstanceProperty = "memory_mb"
	MutableInstancePropertyVcpus         MutableInstanceProperty = "vcpus"
	MutableInstancePropertyScaleToZero   MutableInstanceProperty = "scale_to_zero"
	MutableInstancePropertyTags          MutableInstanceProperty = "tags"
	MutableInstancePropertyDeleteLock    MutableInstanceProperty = "delete_lock"
	MutableInstancePropertySchedules     MutableInstanceProperty = "schedules"
	MutableInstancePropertyAutokill      MutableInstanceProperty = "autokill"
	MutableInstancePropertyHostname      MutableInstanceProperty = "hostname"
	MutableInstancePropertyRoms          MutableInstanceProperty = "roms"
	MutableInstancePropertyDependencies  MutableInstanceProperty = "dependencies"
	MutableInstancePropertySchedPriority MutableInstanceProperty = "sched_priority"
	MutableInstancePropertyPlugins       MutableInstanceProperty = "plugins"
	MutableInstancePropertyAnnotations   MutableInstanceProperty = "annotations"
)
