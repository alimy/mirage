// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mirage/dao"
	"github.com/alimy/mirage/internal"
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type dockerSrv struct {
	base
	broker dao.Broker
}

func (s *dockerSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{
		gin.Logger(),
		gin.Recovery(),
	}
}

func (s *dockerSrv) DockerInfo(c *gin.Context) {
	info, err := s.broker.DockerInfo()
	s.resp(c, info, err)
}

func (s *dockerSrv) VersionInfo(c *gin.Context) {
	s.ok(c, s.broker.VersionInfo())
}

func (s *dockerSrv) Ping(c *gin.Context) {
	pong, err := s.broker.Ping()
	s.resp(c, pong, err)

}

func newDockerSrv() api.Docker {
	return &dockerSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
