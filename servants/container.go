// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"strconv"

	"github.com/alimy/mirage/dao"
	"github.com/alimy/mirage/internal"
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

type containerSrv struct {
	base
	broker dao.Broker
}

func (s *containerSrv) ListContainer(c *gin.Context) {
	// TODO
}

func (s *containerSrv) NewContainer(c *gin.Context) {
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
	id := c.Param("containerId")
	volume, _ := strconv.ParseBool(c.Query("volume"))
	link, _ := strconv.ParseBool(c.Query("link"))
	force, _ := strconv.ParseBool(c.Query("force"))
	options := types.ContainerRemoveOptions{
		RemoveVolumes: volume,
		RemoveLinks:   link,
		Force:         force,
	}
	if err := s.broker.RemoveContainer(id, options); err == nil {
		s.success(c)
	} else {
		s.error(c, err)
	}
}

func (s *containerSrv) ContainerInfo(c *gin.Context) {
	// TODO
}

func (s *containerSrv) ContainerPartLog(c *gin.Context) {
	// TODO
}

func (s *containerSrv) ContainerAllLog(c *gin.Context) {
	// TODO
}

func newContainerSrv() api.Container {
	return &containerSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
