// Code generated by go-mir. DO NOT EDIT.

package api

import (
	gin "github.com/gin-gonic/gin"
)

type Network interface {
	ListNetwork(*gin.Context)
	NetworkInfo(*gin.Context)
	CreateNetwork(*gin.Context)
	RemoveNetwork(*gin.Context)
	ConnectNetwork(*gin.Context)
}

// RegisterNetworkServant register Network servant to gin
func RegisterNetworkServant(e *gin.Engine, s Network) {
	router := e

	// register routes info to router
	router.Handle("GET", "/api/network", s.ListNetwork)
	router.Handle("GET", "/api/network/info/:networkId", s.NetworkInfo)
	router.Handle("GET", "/api/network/new", s.CreateNetwork)
	router.Handle("GET", "/api/network/delete/:networkId", s.RemoveNetwork)
	router.Handle("GET", "/api/network/container/:networkId/:containerId/:operator", s.ConnectNetwork)
}