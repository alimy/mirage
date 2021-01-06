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
