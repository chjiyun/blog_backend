package router

import (
	"blog_backend/app/controller/commentController"
	"github.com/gin-gonic/gin"
)

func (r Router) Comment(g *gin.RouterGroup) {
	rg := g.Group("/comment")
	{
		rg.GET("", commentController.SearchComments)
		rg.POST("", commentController.AddComment)
	}
}
