package siteTrafficVo

// SiteTrafficWholeRespVo 当前页面和站点统计数
type SiteTrafficWholeRespVo struct {
	Pv     int `gorm:"size:100;comment:页面pv" json:"pv"`
	SitePv int `gorm:"comment:站点pv" json:"sitePv"`
	SiteUv int `gorm:"comment:站点uv" json:"siteUv"`
}
