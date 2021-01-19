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

type volumeSrv struct {
	base
	broker dao.Broker
}

func (s *volumeSrv) ListVolume(c *gin.Context) {
	// TODO
}

func (s *volumeSrv) CreateVolume(c *gin.Context) {
	// TODO
}

func (s *volumeSrv) VolumeInfo(c *gin.Context) {
	// TODO
}

func (s *volumeSrv) RemoveVolume(c *gin.Context) {
	// TODO
}

func (s *volumeSrv) PruneVolume(c *gin.Context) {
	// TODO
}

func newVolumeSrv() api.Volume {
	return &volumeSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
