// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

// +build jsoniter

package json

import (
	"github.com/json-iterator/go"

	stdJson "encoding/json"
)

var (
	js = jsoniter.ConfigCompatibleWithStandardLibrary

	// Marshal json marshal
	Marshal = js.Marshal
	// Unmarshal json unmarshal
	Unmarshal = js.Unmarshal
	// MarshalIndent json marshal indent
	MarshalIndent = js.MarshalIndent
	// NewDecoder json new decoder function
	NewDecoder = js.NewDecoder
	// NewEncoder json new encoder function
	NewEncoder = js.NewEncoder
)

// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type RawMessage = stdJson.RawMessage
