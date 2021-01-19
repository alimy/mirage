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

type imageSrv struct {
	base
	broker dao.Broker
}

func (s *imageSrv) ListImage(c *gin.Context) {
	// TODO
}

func (s *imageSrv) ImageInfo(c *gin.Context) {
	// TODO
}

func (s *imageSrv) DeleteImage(c *gin.Context) {
	// TODO
}

func (s *imageSrv) TagImage(c *gin.Context) {
	// TODO
}

func (s *imageSrv) SaveImage(c *gin.Context) {
	// TODO
}

func (s *imageSrv) PullImage(c *gin.Context) {
	// TODO
}

func newImageSrv() api.Image {
	return &imageSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
