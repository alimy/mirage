package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Network))
}

type Network struct {
	GetNetworkList    Get `mir:"/api/network"`
	GetNetworkInfo    Get `mir:"/api/network/info/:networkId"`
	CreateNetworkList Get `mir:"/api/network/new"`
	RemoveNetwork     Get `mir:"/api/network/delete/:networkId"`
	ConnectNetwork    Get `mir:"/api/network/container/:networkId/:containerId/:operator"`
}
