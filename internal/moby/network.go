// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package moby

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
)

func (m *moby) ListNetwork() ([]types.NetworkResource, error) {
	var options types.NetworkListOptions
	return m.client.NetworkList(context.Background(), options)
}

func (m *moby) NetworkInfo(networkId string) (types.NetworkResource, error) {
	var options types.NetworkInspectOptions
	return m.client.NetworkInspect(context.Background(), networkId, options)
}

func (m *moby) CreateNetwork(networkName string, driver string) (types.NetworkCreateResponse, error) {
	options := types.NetworkCreate{Driver: driver}
	return m.client.NetworkCreate(context.Background(), networkName, options)
}

func (m *moby) RemoveNetwork(networkName string) error {
	return m.client.NetworkRemove(context.Background(), networkName)
}

func (m *moby) ConnectNetwork(containerId string, networkId string) error {
	var options *network.EndpointSettings
	return m.client.NetworkConnect(context.Background(), networkId, containerId, options)
}

func (m *moby) DisconnectNetwork(containerId string, networkId string, force bool) error {
	return m.client.NetworkDisconnect(context.Background(), networkId, containerId, force)
}
