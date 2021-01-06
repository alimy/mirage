package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Portal))
}

// Portal web ui interface handler
type Portal struct {
	Index         Get  `mir:"/"`
	GetMainAssets Get  `mir:"/index.html"`
	GetCSSAssets  Get  `mir:"/css/*filepath"`
	HeadCSSAssets Head `mir:"/css/*filepath"`
	GetJSAssets   Get  `mir:"/js/*filepath"`
	HeadJSAssets  Head `mir:"/js/*filepath"`
}
