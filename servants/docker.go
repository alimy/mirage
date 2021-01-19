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

func (s *dockerSrv) DockerInfo(c *gin.Context) {
	if info, err := s.broker.DockerInfo(); err == nil {
		s.ok(c, info)
	} else {
		s.error(c, err)
	}
}

func (s *dockerSrv) VersionInfo(c *gin.Context) {
	s.ok(c, s.broker.VersionInfo())
}

func (s *dockerSrv) Ping(c *gin.Context) {
	if pong, err := s.broker.Ping(); err == nil {
		s.ok(c, pong)
	} else {
		s.error(c, err)
	}

}

func newDockerSrv() api.Docker {
	return &dockerSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
