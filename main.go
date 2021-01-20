// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/alimy/mirage/servants"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	host        string
	port        uint
	inDebug     bool
	showVersion bool

	version = "v0.1.0"
)

func init() {
	flag.StringVar(&host, "host", "", "listening host")
	flag.UintVar(&port, "port", 7878, "listening port")
	flag.BoolVar(&inDebug, "debug", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version")
}

func main() {
	flag.Parse()
	setup()

	if showVersion {
		fmt.Println(version)
		return
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	if host == "" {
		host = "localhost"
	}
	fmt.Printf("listening in [%s]. Please open http://%s:%d in browser to enjoy yourself.\n", addr, host, port)

	e := gin.Default()
	registerServants(e)
	if err := e.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func setup() {
	ginRunMode, logrusLevel := gin.ReleaseMode, logrus.InfoLevel
	if inDebug {
		ginRunMode, logrusLevel = gin.DebugMode, logrus.DebugLevel
	}
	gin.SetMode(ginRunMode)
	logrus.SetLevel(logrusLevel)
}

func registerServants(e *gin.Engine) {
	api.RegisterContainerServant(e, servants.NewContainerSrv())
	api.RegisterDockerServant(e, servants.NewDockerSrv())
	api.RegisterImageServant(e, servants.NewImageSrv())
	api.RegisterNetworkServant(e, servants.NewNetworkSrv())
	api.RegisterVolumeServant(e, servants.NewVolumeSrv())
	api.RegisterPortalServant(e, servants.NewPortalSrv())
}
