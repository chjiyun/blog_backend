package siteTrafficService

import (
	"blog_backend/app/controller/siteTrafficController/siteTrafficVo"
	"blog_backend/app/model"
	"blog_backend/config"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type IpDate struct {
	Ip   string `json:"ip"`
	Date string `json:"date"`
}

// AddSiteTraffic 更新浏览量并返回最新值 后期要加锁，防止高并发
func AddSiteTraffic(c *gin.Context, reqVo siteTrafficVo.SiteTrafficAddReqVo) (*siteTrafficVo.SiteTrafficWholeRespVo, error) {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
		return nil, errors.New(ip + "is not considered")
	}
	siteTraffic := model.SiteTraffic{
		LinkUrl: reqVo.LinkUrl,
		Ip:      ip,
		Browser: reqVo.Browser,
		Os:      reqVo.Os,
		Device:  reqVo.Device,
	}
	go func(data model.SiteTraffic) {
		config.DB.Create(&data)
	}(siteTraffic)

	return cacheSiteTraffic(c, reqVo.LinkUrl, ip)
}

func cacheSiteTraffic(c *gin.Context, linkUrl, ip string) (*siteTrafficVo.SiteTrafficWholeRespVo, error) {
	redisDb := config.RedisDb
	var resp siteTrafficVo.SiteTrafficWholeRespVo

	// sitePv自增1  =>  counter
	sitePv, err := redisDb.Incr(c, "blog:site_pv").Result()
	if err != nil {
		return nil, err
	}
	// post pv  =>  hash
	pv, err := redisDb.HIncrBy(c, "blog:post_pv", linkUrl, 1).Result()
	if err != nil {
		return nil, err
	}
	today := time.Now().Format(time.DateOnly)
	// site uv  => bloom filter
	result, err := redisDb.Do(c, "BF.ADD", "blog:site_uv_filter", today+ip).Int64()
	if err != nil {
		return nil, err
	}
	var siteUv int64
	// 插入成功返回1,可能存在返回0
	if result == 1 {
		siteUv, err = redisDb.Incr(c, "blog:site_uv").Result()
		if err != nil {
			return nil, err
		}
	} else {
		siteUv, err = redisDb.Get(c, "blog:site_uv").Int64()
		if err != nil {
			return nil, err
		}
	}
	resp.SitePv = int(sitePv)
	resp.Pv = int(pv)
	resp.SiteUv = int(siteUv)

	return &resp, nil
}

// GetSiteTraffic 实时查询 插入逻辑：同一ip一天内访问多次只算一次uv.
func GetSiteTraffic(c *gin.Context, linkUrl string) (*siteTrafficVo.SiteTrafficWholeRespVo, error) {
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

	uvSql := "select count(*) from (" +
		"select ip, DATE_FORMAT(created_at, '%Y-%m-%d') date_time from site_traffic where is_del = 0 GROUP BY ip, date_time" +
		") t"
	db.Raw(uvSql).Scan(&siteUv)
	data.SiteUv = int(siteUv)
	return &data, nil
}

// GetManySiteTrafficPv 获取多篇文章浏览量
func GetManySiteTrafficPv(c *gin.Context, linkUrls []string) ([]siteTrafficVo.SiteTrafficUrlRespVo, error) {
	redisDb := config.RedisDb
	resp := make([]siteTrafficVo.SiteTrafficUrlRespVo, 0, len(linkUrls))
	result, err := redisDb.HMGet(c, "blog:post_pv", linkUrls...).Result()
	if err != nil {
		return resp, nil
	}
	for i, val := range result {
		str, ok := val.(string)
		var pv int64
		if ok {
			pv, err = strconv.ParseInt(str, 10, 64)
		}
		item := siteTrafficVo.SiteTrafficUrlRespVo{
			LinkUrl: linkUrls[i],
			Pv:      int(pv),
		}
		resp = append(resp, item)
	}
	return resp, nil
}

// SyncSiteTraffic pv siteUv sitePv 同步至redis
func SyncSiteTraffic(resetSiteUvFilter bool) error {
	db := config.DB
	redisDb := config.RedisDb
	var siteTraffic model.SiteTraffic
	var results []siteTrafficVo.SiteTrafficUrlRespVo
	var sitePv int64
	var siteUv int64

	err := db.Model(&siteTraffic).Count(&sitePv).Error
	if err != nil {
		return err
	}
	// 重置redis的计数器
	redisDb.Set(context.Background(), "blog:site_pv", sitePv, 0)
	err = db.Model(&siteTraffic).Select("link_url, count(*) pv").Group("link_url").Find(&results).Error
	if err != nil {
		return err
	}

	// 值必须是interface{}类型
	postPv := make(map[string]interface{}, len(results))
	for _, item := range results {
		postPv[item.LinkUrl] = item.Pv
	}
	// hash 计数器
	redisDb.HSet(context.Background(), "blog:post_pv", postPv)

	// 数据写进布隆过滤器
	if resetSiteUvFilter {
		var ipDates []IpDate
		err = db.Model(&siteTraffic).Select("ip, DATE_FORMAT(created_at, '%Y-%m-%d') as date").
			Group("ip, date").Order("date").Find(&ipDates).Error
		if err != nil {
			return err
		}
		args := make([]interface{}, 0, len(ipDates)+2)
		args = append(args, "BF.MADD", "blog:site_uv_filter")
		for _, item := range ipDates {
			args = append(args, item.Date+item.Ip)
		}
		redisDb.Do(context.Background(), args...)
	}
	uvSql := "select count(*) from (" +
		"select ip, DATE_FORMAT(created_at, '%Y-%m-%d') date_time from site_traffic where is_del = 0 GROUP BY ip, date_time" +
		") t"
	db.Raw(uvSql).Scan(&siteUv)
	redisDb.Set(context.Background(), "blog:site_uv", siteUv, 0)

	return nil
}
