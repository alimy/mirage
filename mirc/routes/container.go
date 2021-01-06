package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Container))
}

type Container struct {
	GetContainer       Get `mir:"/api/container"`
	CreateNewContainer Get `mir:"/api/container/run"`
	StartContainer     Get `mir:"/api/container/start/:containerId"`
	RestartContainer   Get `mir:"/api/container/restart/:containerId"`
	StopContainer      Get `mir:"/api/container/stop/:containerId"`
	RemoveContainer    Get `mir:"/api/container/delete/:containerId"`
	GetContainerInfo   Get `mir:"/api/container/info/:containerId"`
	GetContainerLog    Get `mir:"/api/container/log/part/:containerId"`
	GetContainerAllLog Get `mir:"/api/container/log/all/:containerId"`
}
