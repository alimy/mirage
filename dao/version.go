// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package dao

import "github.com/docker/docker/api/types"

type Version struct {
	ClientVer string        `json:"client"`
	ServerVer types.Version `json:"server"`
}
