// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

import (
	"encoding/json"
	"fmt"
)

// ImageSource holds exactly one of ImageReference, ImageSpec.
//
// It encodes to the JSON value of whichever member is set, and decodes into the
// member matching the kind of the JSON value.  A struct rather than an
// interface, so that it can be deep copied and embedded in a CRD.
type ImageSource struct {
	ImageReference *ImageReference `json:"-"`
	ImageSpec      *ImageSpec      `json:"-"`
}

type ImageReference string

// MarshalJSON encodes the member that is set, or null if none is.
func (u ImageSource) MarshalJSON() ([]byte, error) {
	switch {
	case u.ImageReference != nil:
		return json.Marshal(u.ImageReference)
	case u.ImageSpec != nil:
		return json.Marshal(u.ImageSpec)
	}
	return []byte("null"), nil
}

// UnmarshalJSON decodes data into the member matching its JSON kind.
func (u *ImageSource) UnmarshalJSON(data []byte) error {
	*u = ImageSource{}

	kind := jsonKind(data)
	if kind == 'n' {
		return nil
	}

	// The kind narrows the candidates but does not prove a match: a failed
	// decode falls through to the members that share the kind, and its error is
	// only reported once none of them matched either.
	var err error

	switch kind {
	case '"':
		var out ImageReference
		if err = json.Unmarshal(data, &out); err == nil {
			u.ImageReference = &out
			return nil
		}
	case '{':
		var out ImageSpec
		if err = json.Unmarshal(data, &out); err == nil {
			u.ImageSpec = &out
			return nil
		}
	}

	if err != nil {
		return err
	}

	return fmt.Errorf("cannot unmarshal JSON %q into ImageSource", kind)
}
