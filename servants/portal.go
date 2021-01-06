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

func NewPortalSrv() api.Portal {
	return &portalSrv{
		staticHandler: http.FileServer(ui.NewFileSystem()),
	}
}
