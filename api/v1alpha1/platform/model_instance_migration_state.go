// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The phase of the migration of an instance.
//
// | Phase        | Description |
// |--------------|-------------|
// | `waiting`    | The instance is starting, stopping or draining, and the migration waits until it has started or stopped. |
// | `queued`     | The migration waits for a transfer slot on the source node. The instance can still be started, stopped and changed. |
// | `copying`    | The instance and the files of its volumes are copied to the destination. |
// | `syncing`    | The running instance is transferred to the destination. |
// | `committing` | The migration is being committed. |
// | `aborting`   | The migration is being aborted. |
// +kubebuilder:validation:Enum="waiting";"queued";"copying";"syncing";"committing";"aborting"
type InstanceMigrationState string

const (
	InstanceMigrationStateWaiting    InstanceMigrationState = "waiting"
	InstanceMigrationStateQueued     InstanceMigrationState = "queued"
	InstanceMigrationStateCopying    InstanceMigrationState = "copying"
	InstanceMigrationStateSyncing    InstanceMigrationState = "syncing"
	InstanceMigrationStateCommitting InstanceMigrationState = "committing"
	InstanceMigrationStateAborting   InstanceMigrationState = "aborting"
)
