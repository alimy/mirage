// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import "github.com/alimy/mirage/mirc/auto/api"

func NewContainerSrv() api.Container {
	return newContainerSrv()
}

func NewDockerSrv() api.Docker {
	return newDockerSrv()
}

func NewImageSrv() api.Image {
	return newImageSrv()
}

func NewNetworkSrv() api.Network {
	return newNetworkSrv()
}

func NewVolumeSrv() api.Volume {
	return newVolumeSrv()
}
