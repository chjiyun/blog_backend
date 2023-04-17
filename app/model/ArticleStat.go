package model

import "blog_backend/app/common"

type ArticleStat struct {
	LinkUrl string `gorm:"not null;size:100;comment:文章url" json:"linkUrl"`
	Pv      int    `gorm:"comment:浏览量" json:"pv"`
	common.BaseModel
}
