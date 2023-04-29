package syncCacheController

import (
	"blog_backend/app/result"
	"blog_backend/app/service/siteTrafficService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SyncSiteTraffic(c *gin.Context) {
	r := result.New()
	resetSiteUvFilter := c.Query("resetSiteUvFilter")
	flag := false
	if resetSiteUvFilter != "" {
		flag = true
	}
	err := siteTrafficService.SyncSiteTraffic(flag)
	if err != nil {
		c.JSON(http.StatusOK, r.FailErr(err))
		return
	}
	c.JSON(http.StatusOK, r)
}
