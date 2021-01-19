// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package moby

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
)

func (m *moby) ListVolume() (volume.VolumeListOKBody, error) {
	filter := filters.Args{}
	return m.client.VolumeList(context.Background(), filter)
}

func (m *moby) CreateVolume(name string, driver string, labels map[string]string) (types.Volume, error) {
	body := volume.VolumeCreateBody{Name: name, Driver: driver, Labels: labels}
	return m.client.VolumeCreate(context.Background(), body)
}

func (m *moby) VolumeInfo(volumeId string) (types.Volume, error) {
	return m.client.VolumeInspect(context.Background(), volumeId)
}

func (m *moby) RemoveVolume(volumeId string, force bool) error {
	return m.client.VolumeRemove(context.Background(), volumeId, force)
}

func (m *moby) PruneVolume() (types.VolumesPruneReport, error) {
	var arg filters.Args
	return m.client.VolumesPrune(context.Background(), arg)
}
