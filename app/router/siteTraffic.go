package router

import (
	"blog_backend/app/controller/siteTrafficController"
	"github.com/gin-gonic/gin"
)

func (r *Router) Article(g *gin.RouterGroup) {
	rg := g.Group("/site-traffic")
	{
		rg.POST("/", siteTrafficController.AddSiteTraffic)
		rg.GET("/pvs", siteTrafficController.GetSiteTrafficPv)
	}
}
