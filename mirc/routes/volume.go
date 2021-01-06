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
