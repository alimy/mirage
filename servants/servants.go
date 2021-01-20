// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

// RegisterServants register servants to gin.Engine
func RegisterServants(e *gin.Engine) {
	api.RegisterContainerServant(e, newContainerSrv())
	api.RegisterDockerServant(e, newDockerSrv())
	api.RegisterImageServant(e, newImageSrv())
	api.RegisterNetworkServant(e, newNetworkSrv())
	api.RegisterVolumeServant(e, newVolumeSrv())
	api.RegisterPortalServant(e, newPortalSrv())
}
