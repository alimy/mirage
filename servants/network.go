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

type networkSrv struct {
	base
	broker dao.Broker
}

func (s *networkSrv) ListNetwork(c *gin.Context) {
	// TODO
}

func (s *networkSrv) NetworkInfo(c *gin.Context) {
	// TODO
}

func (s *networkSrv) CreateNetwork(c *gin.Context) {
	// TODO
}

func (s *networkSrv) RemoveNetwork(c *gin.Context) {
	// TODO
}

func (s *networkSrv) ConnectNetwork(c *gin.Context) {
	// TODO
}

func newNetworkSrv() api.Network {
	return &networkSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
