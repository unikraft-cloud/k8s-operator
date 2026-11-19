// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The request item for unpinning a single image. Select the image by exactly
// one of `uuid` or `url`.
type UnpinImageRequestItem struct {
	// The UUID of the image to unpin.
	Uuid *string `json:"uuid,omitempty"`
	// The URL of the image to unpin, by tag or by digest. A URL without a
	// digest selects the image the tag currently resolves to on the node.
	Url *string `json:"url,omitempty"`
}
