// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"strconv"

	"github.com/alimy/mirage/dao"
	"github.com/alimy/mirage/internal"
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type volumeSrv struct {
	api.UnimplementedVolumeServant
	base

	broker dao.Broker
}

func (s *volumeSrv) ListVolume(c *gin.Context) {
	volumes, err := s.broker.ListVolume()
	s.resp(c, volumes, err)
}

func (s *volumeSrv) CreateVolume(c *gin.Context) {
	name := c.Query("name")
	driver := c.Query("driver")
	volume, err := s.broker.CreateVolume(name, driver, nil)
	s.resp(c, volume, err)
}

func (s *volumeSrv) VolumeInfo(c *gin.Context) {
	id := c.Param("volumeId")
	info, err := s.broker.VolumeInfo(id)
	s.resp(c, info, err)
}

func (s *volumeSrv) RemoveVolume(c *gin.Context) {
	force, _ := strconv.ParseBool(c.Param("force"))
	id := c.Param("volumeId")
	err := s.broker.RemoveVolume(id, force)
	s.reply(c, err)
}

func (s *volumeSrv) PruneVolume(c *gin.Context) {
	result, err := s.broker.PruneVolume()
	s.resp(c, result, err)
}

func newVolumeSrv() api.Volume {
	return &volumeSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
