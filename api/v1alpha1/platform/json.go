// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "bytes"

// jsonKind reports the kind of the JSON value in data using the same bytes as
// encoding/json/jsontext's Kind, which the generated union decoders dispatch
// on: '"', '{', '[', 't', 'f', 'n', or '0' for any number.  It returns 0 for
// input that holds no value.
func jsonKind(data []byte) byte {
	data = bytes.TrimLeft(data, " \t\r\n")
	if len(data) == 0 {
		return 0
	}

	switch c := data[0]; {
	case c == '-' || (c >= '0' && c <= '9'):
		return '0'
	default:
		return c
	}
}
