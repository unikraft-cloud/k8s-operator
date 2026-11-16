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
type CreateInstanceRequestPlugin struct {
	// The plugin name. It becomes the `<plugin_name>` segment in the plugin
	// endpoint (`.../plugins/<plugin_name>/<path>`). A plugin name has a
	// maximum length of 63 characters and contains only letters (`a`-`z`,
	// `A`-`Z`), digits (`0`-`9`), hyphen (`-`), and underscore (`_`).
	Name string `json:"name"`
	// The plugin's image. The platform loads the image, mounts it at
	// `/uk/plugins/<plugin_name>`, and runs its `init` program when the plugin
	// starts. Accepts either a plain image reference string
	// (`"user/myplugin:latest"`) or an object carrying additional pull
	// configuration (`{"url": "user/myplugin:latest", "pull_policy": "always"}`).
	// Exactly one of `image` and `rom` must be set.
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	Image *ImageSource `json:"image,omitempty"`
	// The plugin's image, under its former name. The platform still accepts
	// it and adds a deprecation warning to the response. Exactly one of
	// `image` and `rom` must be set.
	//
	// Deprecated: Use `image` instead.
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	Rom *ImageSource `json:"rom,omitempty"`
	// Arbitrary JSON configuration that the platform passes to the plugin's
	// `init` program on `STDIN`. Any JSON value works, including a string, a
	// number, or an object.
	Config *runtime.RawExtension `json:"config,omitempty"`
	// How requests to the plugin endpoint are authenticated. Without it, the
	// platform API token of a user with access to the instance (`userdb`) is
	// required. The platform checks the token and strips it before forwarding
	// the request to the plugin.
	Authorization *PluginAuthorization `json:"authorization,omitempty"`
}
