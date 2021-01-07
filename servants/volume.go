// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type volumeSrv struct {
	// TODO
}

func (s *volumeSrv) GetVolumeList(c *gin.Context) {
	// TODO
}

func (s *volumeSrv) NewVolume(c *gin.Context) {
	// TODO
}

func (s *volumeSrv) GetVolumeInfo(c *gin.Context) {
	// TODO
}

func (s *volumeSrv) RemoveVolume(c *gin.Context) {
	// TODO
}

func (s *volumeSrv) PruneVolume(c *gin.Context) {
	// TODO
}

func newVolumeSrv() api.Volume {
	return &volumeSrv{}
}
