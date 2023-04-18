package siteTrafficController

import (
	"blog_backend/app/controller/siteTrafficController/siteTrafficVo"
	"blog_backend/app/result"
	"blog_backend/app/service/siteTrafficService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddSiteTraffic(c *gin.Context) {
	r := result.New()
	var reqVo siteTrafficVo.SiteTrafficAddReqVo
	if err := c.ShouldBindJSON(&reqVo); err != nil {
		c.JSON(http.StatusOK, r.FailErr(err))
		return
	}
	res, err := siteTrafficService.AddSiteTraffic(c, reqVo)
	if err != nil {
		c.JSON(http.StatusOK, r.FailErr(err))
		return
	}
	c.JSON(http.StatusOK, r.Success(res))
}

func GetSiteTrafficPv(c *gin.Context) {
	r := result.New()
	linkUrls := c.QueryArray("linkUrl")
	if len(linkUrls) == 0 {
		c.JSON(http.StatusOK, r.Fail("linkUrl is required"))
		return
	}
	res, err := siteTrafficService.GetManySiteTraffic(c, linkUrls)
	if err != nil {
		c.JSON(http.StatusOK, r.FailErr(err))
		return
	}
	c.JSON(http.StatusOK, r.Success(res))
}
