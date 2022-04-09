// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package dao

import (
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/go-connections/nat"
)

type Broker interface {
	Whoami() string
	ListContainer() ([]types.Container, error)
	NewContainer(imageName string, containerName string, env []string, portBinding map[nat.Port][]nat.PortBinding, pathBind []string) (containerId string, err error)
	StopContainer(containerId string) error
	StartContainer(containerId string) error
	RestartContainer(containerId string) error
	RemoveContainer(containerId string, options types.ContainerRemoveOptions) error
	ContainerInfo(containerId string) (types.ContainerJSON, error)
	ContainerLogs(containerId string, tail string) ([]byte, error)
	DockerInfo() (types.Info, error)
	VersionInfo() *Version
	Ping() (types.Ping, error)
	DiskUsage() (types.DiskUsage, error)
	ListImage() ([]types.ImageSummary, error)
	PullImage(refStr string) (io.ReadCloser, error)
	ImageInfo(imageId string) (types.ImageInspect, error)
	TagImage(source string, target string) error
	DeleteImage(imageId string, forge bool) error
	SaveImage(imageId string) (io.ReadCloser, error)
	ListNetwork() ([]types.NetworkResource, error)
	NetworkInfo(networkId string) (types.NetworkResource, error)
	CreateNetwork(networkName string, driver string) (types.NetworkCreateResponse, error)
	RemoveNetwork(networkName string) error
	ConnectNetwork(containerId string, networkId string) error
	DisconnectNetwork(containerId string, networkId string, force bool) error
	ListVolume() (volume.VolumeListOKBody, error)
	CreateVolume(name string, driver string, labels map[string]string) (types.Volume, error)
	VolumeInfo(volumeId string) (types.Volume, error)
	RemoveVolume(volumeId string, force bool) error
	PruneVolume() (types.VolumesPruneReport, error)
}
