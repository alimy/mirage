// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Volume))
}

type Volume struct {
	GetVolumeList Get `mir:"/api/volume"`
	NewVolume     Get `mir:"/api/volume/new"`
	GetVolumeInfo Get `mir:"/api/volume/info/:volumeId"`
	RemoveVolume  Get `mir:"/api/volume/delete/:volumeId/:force"`
	PruneVolume   Get `mir:"/api/volume/prune"`
}
