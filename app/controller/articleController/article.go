package articleController

import (
	"blog_backend/app/controller/articleController/articleStatVo"
	"blog_backend/app/result"
	"blog_backend/app/service/articleService"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

func UpdateArticlePv(c *gin.Context) {
	r := result.New()
	linkUrl := c.Param("linkUrl")
	if linkUrl == "" {
		c.JSON(http.StatusOK, r.Fail("linkUrl is required"))
		return
	}
	res, err := articleService.UpdateArticlePv(c, linkUrl)
	if err != nil {
		c.JSON(http.StatusOK, r.FailErr(err))
		return
	}
	c.JSON(http.StatusOK, r.Success(res))
}

func GetManyArticlePv(c *gin.Context) {
	r := result.New()
	linkUrls := c.QueryArray("linkUrl")
	if len(linkUrls) == 0 {
		c.JSON(http.StatusOK, r.Fail("linkUrl is required"))
		return
	}
	res, err := articleService.GetManyArticlePv(c, linkUrls)
	if err != nil {
		c.JSON(http.StatusOK, r.FailErr(err))
		return
	}
	var articleStats []articleStatVo.ArticleStatRespVo
	_ = copier.Copy(&articleStats, &res)
	c.JSON(http.StatusOK, r.Success(articleStats))
}
