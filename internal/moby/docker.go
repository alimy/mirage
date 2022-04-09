// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package moby

import (
	"context"

	"github.com/alimy/mirage/internal/dao"
	"github.com/docker/docker/api/types"
)

func (m *moby) DockerInfo() (types.Info, error) {
	info, err := m.client.Info(context.Background())
	if err != nil {
		return types.Info{}, err
	}
	return info, nil
}

func (m *moby) VersionInfo() *dao.Version {
	clientVersion := m.client.ClientVersion()
	serverVersion, _ := m.client.ServerVersion(context.Background())
	return &dao.Version{
		ClientVer: clientVersion,
		ServerVer: serverVersion,
	}
}

func (m *moby) Ping() (types.Ping, error) {
	return m.client.Ping(context.Background())
}

func (m *moby) DiskUsage() (types.DiskUsage, error) {
	return m.client.DiskUsage(context.Background())
}
