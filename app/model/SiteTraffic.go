package model

import "blog_backend/app/common"

type SiteTraffic struct {
	LinkUrl string `gorm:"not null;size:100;comment:文章url" json:"linkUrl"`
	Ip      string `gorm:"comment:ip" json:"ip"`
	Browser string `gorm:"comment:浏览器信息" json:"browser"`
	Os      string `gorm:"comment:操作系统" json:"os"`
	Device  string `gorm:"comment:设备信息" json:"device"`
	common.BaseModel
}
