// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import "k8s.io/apimachinery/pkg/runtime"

// The property to modify.
// +kubebuilder:validation:Enum=image;args;env;memory_mb;vcpus;scale_to_zero;tags;delete_lock;schedules;autokill;hostname;roms;dependencies;sched_priority
type UpdateInstanceByUUIDRequestBodyProp string

const (
	UpdateInstanceByUUIDRequestBodyPropImage          UpdateInstanceByUUIDRequestBodyProp = "image"
	UpdateInstanceByUUIDRequestBodyPropArgs           UpdateInstanceByUUIDRequestBodyProp = "args"
	UpdateInstanceByUUIDRequestBodyPropEnv            UpdateInstanceByUUIDRequestBodyProp = "env"
	UpdateInstanceByUUIDRequestBodyPropMemory_mb      UpdateInstanceByUUIDRequestBodyProp = "memory_mb"
	UpdateInstanceByUUIDRequestBodyPropVcpus          UpdateInstanceByUUIDRequestBodyProp = "vcpus"
	UpdateInstanceByUUIDRequestBodyPropScale_to_zero  UpdateInstanceByUUIDRequestBodyProp = "scale_to_zero"
	UpdateInstanceByUUIDRequestBodyPropTags           UpdateInstanceByUUIDRequestBodyProp = "tags"
	UpdateInstanceByUUIDRequestBodyPropDelete_lock    UpdateInstanceByUUIDRequestBodyProp = "delete_lock"
	UpdateInstanceByUUIDRequestBodyPropSchedules      UpdateInstanceByUUIDRequestBodyProp = "schedules"
	UpdateInstanceByUUIDRequestBodyPropAutokill       UpdateInstanceByUUIDRequestBodyProp = "autokill"
	UpdateInstanceByUUIDRequestBodyPropHostname       UpdateInstanceByUUIDRequestBodyProp = "hostname"
	UpdateInstanceByUUIDRequestBodyPropRoms           UpdateInstanceByUUIDRequestBodyProp = "roms"
	UpdateInstanceByUUIDRequestBodyPropDependencies   UpdateInstanceByUUIDRequestBodyProp = "dependencies"
	UpdateInstanceByUUIDRequestBodyPropSched_priority UpdateInstanceByUUIDRequestBodyProp = "sched_priority"
)

// The operation to perform on the property.
// +kubebuilder:validation:Enum=set;add;del
type UpdateInstanceByUUIDRequestBodyOp string

const (
	UpdateInstanceByUUIDRequestBodyOpSet UpdateInstanceByUUIDRequestBodyOp = "set"
	UpdateInstanceByUUIDRequestBodyOpAdd UpdateInstanceByUUIDRequestBodyOp = "add"
	UpdateInstanceByUUIDRequestBodyOpDel UpdateInstanceByUUIDRequestBodyOp = "del"
)

type UpdateInstanceByUUIDRequestBody struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop UpdateInstanceByUUIDRequestBodyProp `json:"prop"`
	// The operation to perform on the property.
	Op UpdateInstanceByUUIDRequestBodyOp `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "image": string
	// - For "args": string or array of strings
	// - For "env": object (for SET/ADD) or string/array of strings (for DEL)
	// - For "memory_mb": integer
	// - For "vcpus": integer
	// - For "scale_to_zero": object with cooldown_time_ms, policy, and stateful fields
	// - For "tags": array of strings
	// - For "delete_lock": boolean
	// - For "schedules": array of schedule objects (with name, when, action, and optional args fields).
	//   Use action "exec" together with args to execute a command at the scheduled time.
	// - For "autokill": object with time_ms and num_requests fields
	// - For "hostname": string (valid DNS label)
	// - For "roms": array of ROM objects (with name and image fields) for SET/ADD, or array of ROM names for DEL
	// - For "dependencies": array of instance identifiers (name or UUID)
	// - For "sched_priority": integer (scheduling priority value)
	Value *runtime.RawExtension `json:"value,omitempty"`
}
