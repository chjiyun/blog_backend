package router

import (
	"blog_backend/app/controller/syncCacheController"
	"github.com/gin-gonic/gin"
)

func (r Router) Sync(g *gin.RouterGroup) {
	rg := g.Group("/sync")
	{
		rg.POST("/site-traffic", syncCacheController.SyncSiteTraffic)
	}
}
