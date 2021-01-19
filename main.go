// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/alimy/mirage/servants"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	// register servants to gin
	registerServants(e)

	// start servant service
	if err := e.Run(":7878"); err != nil {
		log.Fatal(err)
	}
}

func registerServants(e *gin.Engine) {
	api.RegisterContainerServant(e, servants.NewContainerSrv())
	api.RegisterDockerServant(e, servants.NewDockerSrv())
	api.RegisterImageServant(e, servants.NewImageSrv())
	api.RegisterNetworkServant(e, servants.NewNetworkSrv())
	api.RegisterVolumeServant(e, servants.NewVolumeSrv())
	api.RegisterPortalServant(e, servants.NewPortalSrv())
}
