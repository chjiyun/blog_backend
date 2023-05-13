package model

import (
	"blog_backend/app/common"
	"time"
)

// Comment 评论
type Comment struct {
	common.BaseModel
	UserId     uint      `gorm:"comment:ip" json:"userId"`
	Ip         string    `gorm:"comment:ip" json:"-"`
	Link       string    `gorm:"size:100;comment:用户输入网址" json:"link"`
	Mail       string    `gorm:"size:100;comment:邮箱" json:"-"`
	Nick       string    `gorm:"size:100;comment:昵称" json:"nick"`
	Comment    string    `gorm:"comment:评论" json:"comment"`
	InsertedAt time.Time `gorm:"comment:评论时间" json:"insertedAt"`
	Pid        uint      `gorm:"not null;default:0;comment:父节点" json:"pid"`
	Rid        uint      `gorm:"not null;default:0;comment:根节点" json:"rid"`
	Sticky     int       `gorm:"comment:根节点" json:"sticky"`
	Status     string    `gorm:"size:50;comment:状态" json:"status"`
	Url        string    `gorm:"size:100;comment:文章pathname" json:"-"`
	Like       int       `gorm:"default:0;comment:点赞数" json:"like"`
	Browser    string    `gorm:"size:50;comment:浏览器信息" json:"browser"`
	Os         string    `gorm:"size:50;comment:操作系统" json:"os"`
}
