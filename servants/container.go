// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type containerSrv struct {
	// TODO
}

func (s *containerSrv) GetContainer(c *gin.Context) {
	// TODO
}

func (s *containerSrv) CreateNewContainer(c *gin.Context) {
	// TODO
}

func (s *containerSrv) StartContainer(c *gin.Context) {
	// TODO
}

func (s *containerSrv) RestartContainer(c *gin.Context) {
	// TODO
}

func (s *containerSrv) StopContainer(c *gin.Context) {
	// TODO
}

func (s *containerSrv) RemoveContainer(c *gin.Context) {
	// TODO
}

func (s *containerSrv) GetContainerInfo(c *gin.Context) {
	// TODO
}

func (s *containerSrv) GetContainerLog(c *gin.Context) {
	// TODO
}

func (s *containerSrv) GetContainerAllLog(c *gin.Context) {
	// TODO
}

func newContainerSrv() api.Container {
	return &containerSrv{}
}
