// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// Parameters for suspending a user account.
type SuspendUserByUUIDRequestBody struct {
	// Whether to also bring all instances of the user to rest. Autostart is
	// disabled, running instances are suspended (or stopped if they cannot be
	// suspended, or if they only return to standby and are not stateful), and
	// pending restarts are cancelled. Instances that are draining, stopping, or
	// in a start-stop cycle finish that transition first. Booting instances are
	// brought to rest once they are running. Without this flag only the account
	// is suspended and its instances are left as they are.
	SuspendInstances *bool `json:"suspend_instances,omitempty"`
}
