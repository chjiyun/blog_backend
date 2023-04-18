package siteTrafficVo

type SiteTrafficAddReqVo struct {
	LinkUrl string `form:"linkUrl" json:"linkUrl" binding:"required"`
	Ua      string `gorm:"comment:user agent" json:"ua"`
}
