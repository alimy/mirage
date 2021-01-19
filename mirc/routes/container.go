// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Container))
}

type Container struct {
	ListContainer    Get `mir:"/api/container"`
	NewContainer     Get `mir:"/api/container/run"`
	StartContainer   Get `mir:"/api/container/start/:containerId"`
	RestartContainer Get `mir:"/api/container/restart/:containerId"`
	StopContainer    Get `mir:"/api/container/stop/:containerId"`
	RemoveContainer  Get `mir:"/api/container/delete/:containerId"`
	ContainerInfo    Get `mir:"/api/container/info/:containerId"`
	ContainerPartLog Get `mir:"/api/container/log/part/:containerId"`
	ContainerAllLog  Get `mir:"/api/container/log/all/:containerId"`
}
