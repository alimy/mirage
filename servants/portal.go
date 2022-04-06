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
	api.UnimplementedPortalServant

	staticHandler http.Handler
}

func (p *portalSrv) Index(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func (p *portalSrv) IndexHtml(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func (p *portalSrv) CSSAssets(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func (p *portalSrv) JSAssets(c *gin.Context) {
	p.staticHandler.ServeHTTP(c.Writer, c.Request)
}

func newPortalSrv() api.Portal {
	return &portalSrv{
		staticHandler: http.FileServer(ui.NewFileSystem()),
	}
}
