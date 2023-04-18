package siteTrafficVo

// SiteTrafficUrlRespVo 页面pv统计
type SiteTrafficUrlRespVo struct {
	LinkUrl string `gorm:"size:100;comment:文章url" json:"linkUrl"`
	Pv      int    `gorm:"comment:访问量" json:"pv"`
	Uv      int    `gorm:"comment:访客量" json:"-"`
}
