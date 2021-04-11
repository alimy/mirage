// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

//go:build !jsoniter
// +build !jsoniter

package json

import (
	stdJson "encoding/json"
)

var (
	// Marshal json marshal
	Marshal = stdJson.Marshal
	// Unmarshal json unmarshal
	Unmarshal = stdJson.Unmarshal
	// MarshalIndent json marshal indent
	MarshalIndent = stdJson.MarshalIndent
	// NewDecoder json new decoder function
	NewDecoder = stdJson.NewDecoder
	// NewEncoder json new encoder function
	NewEncoder = stdJson.NewEncoder
)

// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type RawMessage = stdJson.RawMessage
