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

	//var count int64
	//startTime := time.Now().Format(time.DateOnly)
	//endTime := time.Now().AddDate(0, 0, 1).Format(time.DateOnly)
	//
	//err := db.Where("link_url = ? and ip = ?", reqVo.LinkUrl, reqVo.Ip).
	//	Where("create_at > ? and created_at < ?", startTime, endTime).
	//	Count(&count).Error
	//if err != nil {
	//	return nil, err
	//}
	siteTraffic := model.SiteTraffic{
		LinkUrl: reqVo.LinkUrl,
		Ip:      c.ClientIP(),
		Ua:      reqVo.Ua,
	}
	err := db.Create(&siteTraffic).Error
	if err != nil {
		return nil, err
	}
	return getPageAndSiteStat(c, reqVo.LinkUrl)
}

// 插入逻辑：同一ip一天内访问多次只算一次uv.
func getPageAndSiteStat(c *gin.Context, linkUrl string) (*siteTrafficVo.SiteTrafficWholeRespVo, error) {
	db := c.Value("DB").(*gorm.DB)
	var siteTraffic model.SiteTraffic

	var data siteTrafficVo.SiteTrafficWholeRespVo
	var pv int64
	var sitePv int64
	var siteUv int64

	db.Model(&siteTraffic).Where("link_url = ?", linkUrl).Count(&pv)
	data.Pv = int(pv)

	db.Model(&siteTraffic).Count(&sitePv)
	data.SitePv = int(sitePv)

	uvSql := "select count(*) from (select ip, DATE_FORMAT(created_at, '%Y-%m-%d') date_time, count(*) count from site_traffic where is_del = 0 GROUP BY ip, date_time) t"
	db.Raw(uvSql).Scan(&siteUv)
	data.SiteUv = int(siteUv)
	return &data, nil
}

// GetManySiteTraffic 获取多篇文章浏览量
func GetManySiteTraffic(c *gin.Context, linkUrls []string) ([]siteTrafficVo.SiteTrafficUrlRespVo, error) {
	db := c.Value("DB").(*gorm.DB)

	var siteTraffics []model.SiteTraffic
	tx := db.Where("link_url in ?", linkUrls).Find(&siteTraffics)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// 初始化
	list := make([]siteTrafficVo.SiteTrafficUrlRespVo, len(linkUrls))
	//for i, item := range siteTraffics {
	//	list[i].LinkUrl = item.LinkUrl
	//	list[i].Pv = item.LinkUrl
	//}
	return list, nil
}
