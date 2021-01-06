package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Docker))
}

type Docker struct {
	DockerInfo Get `mir:"/api/docker/info"`
	GetVersion Get `mir:"/api/docker/version"`
	Ping       Get `mir:"/api/docker/ping"`
}
