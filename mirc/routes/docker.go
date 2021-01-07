// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Docker))
}

type Docker struct {
	DockerInfo Get `mir:"/api/docker/info"`
	GetVersion Get `mir:"/api/docker/version"`
	Ping       Get `mir:"/api/docker/ping"`
}
