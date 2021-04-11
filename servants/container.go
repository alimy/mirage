// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/alimy/mirage/dao"
	"github.com/alimy/mirage/internal"
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
)

type containerSrv struct {
	api.UnimplementedContainerServant
	base

	broker dao.Broker
}

func (s *containerSrv) ListContainer(c *gin.Context) {
	cs, err := s.broker.ListContainer()
	s.resp(c, cs, err)
}

func (s *containerSrv) NewContainer(c *gin.Context) {
	imageName := c.Query("imageName")
	containerName := c.Query("containerName")
	port := c.Query("bindPort")
	env := c.Query("env")
	volume := c.Query("volume")

	var envs, volumes []string
	if env = strings.Trim(env, " "); env != "" {
		envs = strings.Split(env, ";")
	}
	portBindings := map[nat.Port][]nat.PortBinding{}
	if port = strings.Trim(port, " "); port != "" {
		portKeys := strings.Split(port, ";")
		for _, key := range portKeys {
			ports := strings.Split(key, ":")
			if len(ports) != 2 {
				continue
			}
			portBinding := nat.PortBinding{HostIP: "", HostPort: ports[0]}
			portBindings[nat.Port(ports[1]+"/tcp")] = []nat.PortBinding{portBinding}
		}
	}
	if volume = strings.Trim(volume, " "); volume != "" {
		volumes = strings.Split(volume, ";")
	}
	cid, err := s.broker.NewContainer(imageName, containerName, envs, portBindings, volumes)
	s.resp(c, cid, err)
}

func (s *containerSrv) StartContainer(c *gin.Context) {
	id := c.Param("containerId")
	err := s.broker.StartContainer(id)
	s.reply(c, err)
}

func (s *containerSrv) RestartContainer(c *gin.Context) {
	id := c.Param("containerId")
	err := s.broker.RestartContainer(id)
	s.reply(c, err)
}

func (s *containerSrv) StopContainer(c *gin.Context) {
	id := c.Param("containerId")
	err := s.broker.StopContainer(id)
	s.reply(c, err)
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
	err := s.broker.RemoveContainer(id, options)
	s.reply(c, err)
}

func (s *containerSrv) ContainerInfo(c *gin.Context) {
	id := c.Param("containerId")
	info, err := s.broker.ContainerInfo(id)
	s.resp(c, info, err)
}

func (s *containerSrv) ContainerPartLog(c *gin.Context) {
	id := c.Param("containerId")
	logs, err := s.broker.ContainerLogs(id, "200")
	s.resp(c, logs, err)
}

func (s *containerSrv) ContainerAllLog(c *gin.Context) {
	id := c.Param("containerId")
	if logs, err := s.broker.ContainerLogs(id, ""); err == nil {
		c.Header("Content-Type", "application/force-download")
		c.Header("Content-Disposition", fmt.Sprintf("attachment;filename=%s-all.log", id))
		c.Header("Content-Transfer-Encoding", "binary")
		c.Status(http.StatusOK)
		c.Writer.Write(logs)
	} else {
		s.error(c, err)
	}
}

func newContainerSrv() api.Container {
	return &containerSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
