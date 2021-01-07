// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

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
