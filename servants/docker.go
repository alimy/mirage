// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type dockerSrv struct {
	// TODO
}

func (s *dockerSrv) DockerInfo(c *gin.Context) {
	// TODO
}

func (s *dockerSrv) GetVersion(c *gin.Context) {
	// TODO
}

func (s *dockerSrv) Ping(c *gin.Context) {
	// TODO
}

func newDockerSrv() api.Docker {
	return &dockerSrv{}
}
