// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

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
