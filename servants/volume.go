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
