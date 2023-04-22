package siteTrafficService

import (
	"blog_backend/app/controller/siteTrafficController/siteTrafficVo"
	"blog_backend/app/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddSiteTraffic 更新浏览量并返回最新值 后期要加锁，防止高并发
func AddSiteTraffic(c *gin.Context, reqVo siteTrafficVo.SiteTrafficAddReqVo) (*siteTrafficVo.SiteTrafficWholeRespVo, error) {
	db := c.Value("DB").(*gorm.DB)

	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	siteTraffic := model.SiteTraffic{
		LinkUrl: reqVo.LinkUrl,
		Ip:      ip,
		Browser: reqVo.Browser,
		Os:      reqVo.Os,
		Device:  reqVo.Device,
	}
	err := db.Create(&siteTraffic).Error
	if err != nil {
		return nil, err
	}
	return getPageAndSiteStat(c, reqVo.LinkUrl, reqVo.NoPv)
}

// 插入逻辑：同一ip一天内访问多次只算一次uv.
func getPageAndSiteStat(c *gin.Context, linkUrl string, noPv bool) (*siteTrafficVo.SiteTrafficWholeRespVo, error) {
	db := c.Value("DB").(*gorm.DB)
	var siteTraffic model.SiteTraffic

	var data siteTrafficVo.SiteTrafficWholeRespVo
	var pv int64
	var sitePv int64
	var siteUv int64

	if !noPv {
		db.Model(&siteTraffic).Where("link_url = ?", linkUrl).Count(&pv)
		data.Pv = int(pv)
	}

	db.Model(&siteTraffic).Count(&sitePv)
	data.SitePv = int(sitePv)

	uvSql := "select count(*) from (" +
		"select ip, DATE_FORMAT(created_at, '%Y-%m-%d') date_time, count(*) count from site_traffic where is_del = 0 GROUP BY ip, date_time" +
		") t"
	db.Raw(uvSql).Scan(&siteUv)
	data.SiteUv = int(siteUv)
	return &data, nil
}

// GetManySiteTrafficPv 获取多篇文章浏览量
func GetManySiteTrafficPv(c *gin.Context, linkUrls []string) ([]siteTrafficVo.SiteTrafficUrlRespVo, error) {
	db := c.Value("DB").(*gorm.DB)

	var results []siteTrafficVo.SiteTrafficUrlRespVo
	tx := db.Model(&model.SiteTraffic{}).
		Select("link_url, count(*) pv").
		Where("link_url in ?", linkUrls).
		Group("link_url").
		Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return results, nil
}
