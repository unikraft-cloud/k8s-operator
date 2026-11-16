// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// A helper program attached to the instance and reachable over a direct,
// authenticated HTTP endpoint. A plugin runs inside the instance next to
// the main application, loads from its own ROM image, and answers requests
// that the Unikraft Cloud API forwards to it.
type InstancePlugin struct {
	// The plugin name. It becomes the `<plugin_name>` segment in the plugin
	// endpoint (`.../plugins/<plugin_name>/<path>`). A plugin name has a
	// maximum length of 63 characters and contains only letters (`a`-`z`,
	// `A`-`Z`), digits (`0`-`9`), hyphen (`-`), and underscore (`_`).
	Name string `json:"name"`
	// The plugin's image, given as an image reference string such as
	// `user/myplugin:latest`. The platform loads the image, mounts it at
	// `/uk/plugins/<plugin_name>`, and runs its `init` program when the
	// plugin starts.
	Image string `json:"image"`
	// Arbitrary JSON configuration that the platform passes to the plugin's
	// `init` program on `STDIN`. Any JSON value works, including a string, a
	// number, or an object.
	Config *runtime.RawExtension `json:"config,omitempty"`
	// The port on the instance's address on which the plugin's API listens.
	// The platform picks it when the plugin is applied, so it lets a caller
	// that can reach the instance directly talk to the plugin without going
	// through the platform API.
	ApiPort *uint32 `json:"api_port,omitempty"`
	// The plugin's authorization method. Only present when it is not the
	// default (`userdb`).
	Authorization *InstancePluginAuthorization `json:"authorization,omitempty"`
}
