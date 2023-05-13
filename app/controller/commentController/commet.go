package commentController

import (
	"blog_backend/app/controller/commentController/commentVo"
	"blog_backend/app/result"
	"blog_backend/app/service/commentService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchComments(c *gin.Context) {

}

func AddComment(c *gin.Context) {
	r := result.NewWl()
	var reqVo commentVo.CommentAddReqVo
	if err := c.ShouldBindJSON(&reqVo); err != nil {
		c.JSON(http.StatusOK, r.FailErr(err))
		return
	}
	res, err := commentService.AddComment(c, reqVo)
	if err != nil {
		c.JSON(http.StatusOK, r.FailErr(err))
		return
	}
	c.JSON(http.StatusOK, r.Success(res))
}
