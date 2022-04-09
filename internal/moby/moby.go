// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package moby

import (
	"github.com/alimy/mirage/internal/dao"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

func NewBroker() dao.Broker {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logrus.Fatal(err)
	}
	return &moby{
		client: c,
	}
}
