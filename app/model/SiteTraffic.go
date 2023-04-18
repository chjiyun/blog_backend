package model

import "blog_backend/app/common"

type SiteTraffic struct {
	LinkUrl string `gorm:"not null;size:100;comment:文章url" json:"linkUrl"`
	Ip      string `gorm:"comment:ip" json:"ip"`
	Ua      string `gorm:"comment:user agent" json:"ua"`
	common.BaseModel
}
