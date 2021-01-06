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
