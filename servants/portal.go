// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/alimy/mirage-ui"
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type portalSrv struct {
	staticHandler http.Handler
}

func (p *portalSrv) Index(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func (p *portalSrv) GetMainAssets(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func (p *portalSrv) GetCSSAssets(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func (p *portalSrv) HeadCSSAssets(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func (p *portalSrv) GetJSAssets(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func (p *portalSrv) HeadJSAssets(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func newPortalSrv() api.Portal {
	return &portalSrv{
		staticHandler: http.FileServer(ui.NewFileSystem()),
	}
}
