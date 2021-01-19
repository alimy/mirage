// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package moby

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
)

func (m *moby) ListImage() ([]types.ImageSummary, error) {
	images, err := m.client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (m *moby) PullImage(refStr string) (io.ReadCloser, error) {
	var options types.ImagePullOptions
	return m.client.ImagePull(context.Background(), refStr, options)
}

func (m *moby) ImageInfo(imageId string) (types.ImageInspect, error) {
	raw, _, err := m.client.ImageInspectWithRaw(context.Background(), imageId)
	if err != nil {
		return types.ImageInspect{}, err
	}
	return raw, nil
}

func (m *moby) TagImage(source string, target string) error {
	return m.client.ImageTag(context.Background(), source, target)
}

func (m *moby) DeleteImage(imageId string, forge bool) error {
	removeOption := types.ImageRemoveOptions{Force: forge, PruneChildren: true}
	_, err := m.client.ImageRemove(context.Background(), imageId, removeOption)
	if err != nil {
		return err
	}
	return nil
}

func (m *moby) SaveImage(imageId string) (io.ReadCloser, error) {
	return m.client.ImageSave(context.Background(), []string{imageId})
}
