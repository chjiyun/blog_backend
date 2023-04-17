package router

import (
	"blog_backend/app/controller/articleController"
	"github.com/gin-gonic/gin"
)

func (r *Router) Article(g *gin.RouterGroup) {
	rg := g.Group("/article")
	{
		rg.GET("/pvs", articleController.GetManyArticlePv)
		rg.PUT("/pv/:linkUrl", articleController.UpdateArticlePv)
	}
}
