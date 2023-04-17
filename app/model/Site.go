package model

import "blog_backend/app/common"

type Site struct {
	Pv int `gorm:"comment:站点浏览量" json:"pv"`
	Uv int `gorm:"comment:站点访客量" json:"uv"`
	common.BaseModel
}
