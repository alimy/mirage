// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package moby

import (
	"bytes"
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

func (m *moby) ListContainer() []types.Container {
	containerList, err := m.client.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
		return []types.Container{}
	}
	return containerList
}

func (m *moby) NewContainer(imageName string, containerName string, env []string, portBinding map[nat.Port][]nat.PortBinding, pathBind []string) (containerId string, err error) {
	imageConfig := container.Config{Image: imageName, Env: env}

	hostConfig := container.HostConfig{PortBindings: portBinding, Binds: pathBind}
	resp, err := m.client.ContainerCreate(context.Background(), &imageConfig, &hostConfig, nil, nil, containerName)
	if err != nil {
		_ = m.RemoveContainer(resp.ID, types.ContainerRemoveOptions{Force: true})
		return "", err
	}

	startConfig := types.ContainerStartOptions{}
	if err := m.client.ContainerStart(context.Background(), resp.ID, startConfig); err != nil {
		_ = m.RemoveContainer(resp.ID, types.ContainerRemoveOptions{Force: true})
		return resp.ID, err
	}
	return resp.ID, nil
}

func (m *moby) StopContainer(containerId string) error {
	return m.client.ContainerStop(context.Background(), containerId, nil)
}

func (m *moby) StartContainer(containerId string) error {
	options := types.ContainerStartOptions{}
	return m.client.ContainerStart(context.Background(), containerId, options)
}

func (m *moby) RestartContainer(containerId string) error {
	duration := time.Second * 20
	return m.client.ContainerRestart(context.Background(), containerId, &duration)
}

func (m *moby) RemoveContainer(containerId string, options types.ContainerRemoveOptions) error {
	return m.client.ContainerRemove(context.Background(), containerId, options)
}

func (m *moby) ContainerInfo(containerId string) (types.ContainerJSON, error) {
	return m.client.ContainerInspect(context.Background(), containerId)
}

func (m *moby) ContainerLogs(containerId string, tail string) (string, error) {
	options := types.ContainerLogsOptions{ShowStdout: true}
	if tail != "" {
		options.Tail = tail
	}
	logs, err := m.client.ContainerLogs(context.Background(), containerId, options)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(logs)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
