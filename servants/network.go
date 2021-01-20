// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"errors"
	"strings"

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
	networks, err := s.broker.ListNetwork()
	s.resp(c, networks, err)
}

func (s *networkSrv) NetworkInfo(c *gin.Context) {
	id := c.Param("networkId")
	info, err := s.broker.NetworkInfo(id)
	s.resp(c, info, err)
}

func (s *networkSrv) CreateNetwork(c *gin.Context) {
	name := strings.Trim(c.Query("name"), " ")
	driver := strings.Trim(c.Query("driver"), " ")
	if name == "" {
		s.failure(c, "network name is empty")
		return
	}
	network, err := s.broker.CreateNetwork(name, driver)
	s.resp(c, network, err)
}

func (s *networkSrv) RemoveNetwork(c *gin.Context) {
	id := strings.Trim(c.Param("networkId"), " ")
	if id == "" {
		s.failure(c, "network id is null")
		return
	}
	err := s.broker.RemoveNetwork(id)
	s.reply(c, err)
}

func (s *networkSrv) ConnectNetwork(c *gin.Context) {
	containerId := strings.Trim(c.Param("containerId"), " ")
	networkId := strings.Trim(c.Param("networkId"), " ")
	operator := c.Param("operator")
	if containerId == "" || networkId == "" {
		s.failure(c, "container id or network id is null")
		return
	}
	var err error
	switch operator {
	case "connect":
		err = s.broker.ConnectNetwork(containerId, networkId)
	case "disconnect":
		err = s.broker.DisconnectNetwork(containerId, networkId, true)
	default:
		err = errors.New("operator not support")
	}
	s.error(c, err)
}

func newNetworkSrv() api.Network {
	return &networkSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
