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
type UpdateInstancesRequestItemAllOfProp string

const (
	UpdateInstancesRequestItemAllOfPropImage          UpdateInstancesRequestItemAllOfProp = "image"
	UpdateInstancesRequestItemAllOfPropArgs           UpdateInstancesRequestItemAllOfProp = "args"
	UpdateInstancesRequestItemAllOfPropEnv            UpdateInstancesRequestItemAllOfProp = "env"
	UpdateInstancesRequestItemAllOfPropMemory_mb      UpdateInstancesRequestItemAllOfProp = "memory_mb"
	UpdateInstancesRequestItemAllOfPropVcpus          UpdateInstancesRequestItemAllOfProp = "vcpus"
	UpdateInstancesRequestItemAllOfPropScale_to_zero  UpdateInstancesRequestItemAllOfProp = "scale_to_zero"
	UpdateInstancesRequestItemAllOfPropTags           UpdateInstancesRequestItemAllOfProp = "tags"
	UpdateInstancesRequestItemAllOfPropDelete_lock    UpdateInstancesRequestItemAllOfProp = "delete_lock"
	UpdateInstancesRequestItemAllOfPropSchedules      UpdateInstancesRequestItemAllOfProp = "schedules"
	UpdateInstancesRequestItemAllOfPropAutokill       UpdateInstancesRequestItemAllOfProp = "autokill"
	UpdateInstancesRequestItemAllOfPropHostname       UpdateInstancesRequestItemAllOfProp = "hostname"
	UpdateInstancesRequestItemAllOfPropRoms           UpdateInstancesRequestItemAllOfProp = "roms"
	UpdateInstancesRequestItemAllOfPropDependencies   UpdateInstancesRequestItemAllOfProp = "dependencies"
	UpdateInstancesRequestItemAllOfPropSched_priority UpdateInstancesRequestItemAllOfProp = "sched_priority"
)

// The operation to perform on the property.
// +kubebuilder:validation:Enum=set;add;del
type UpdateInstancesRequestItemAllOfOp string

const (
	UpdateInstancesRequestItemAllOfOpSet UpdateInstancesRequestItemAllOfOp = "set"
	UpdateInstancesRequestItemAllOfOpAdd UpdateInstancesRequestItemAllOfOp = "add"
	UpdateInstancesRequestItemAllOfOpDel UpdateInstancesRequestItemAllOfOp = "del"
)

type UpdateInstancesRequestItemAllOf struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// The property to modify.
	Prop *UpdateInstancesRequestItemAllOfProp `json:"prop,omitempty"`
	// The operation to perform on the property.
	Op *UpdateInstancesRequestItemAllOfOp `json:"op,omitempty"`
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
