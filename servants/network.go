package servants

import (
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type networkSrv struct {
	// TODO
}

func (s *networkSrv) GetNetworkList(c *gin.Context) {
	// TODO
}

func (s *networkSrv) GetNetworkInfo(c *gin.Context) {
	// TODO
}

func (s *networkSrv) CreateNetworkList(c *gin.Context) {
	// TODO
}

func (s *networkSrv) RemoveNetwork(c *gin.Context) {
	// TODO
}

func (s *networkSrv) ConnectNetwork(c *gin.Context) {
	// TODO
}

func newNetworkSrv() api.Network {
	return &networkSrv{}
}
