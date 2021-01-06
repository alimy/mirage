package servants

import (
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type imageSrv struct {
	// TODO
}

func (s *imageSrv) GetImageList(c *gin.Context) {
	// TODO
}

func (s *imageSrv) GetImageInfo(c *gin.Context) {
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
	return &imageSrv{}
}
