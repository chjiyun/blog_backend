package commentService

import (
	"blog_backend/app/controller/commentController/commentVo"
	"blog_backend/app/model"
	"blog_backend/app/service"
	"blog_backend/app/util"
	"blog_backend/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func AddComment(c *gin.Context, reqVo commentVo.CommentAddReqVo) (*commentVo.CommentRespVo, error) {
	db := c.Value("DB").(*gorm.DB)
	log := c.Value("Logger").(*logrus.Entry)

	var resp commentVo.CommentRespVo

	browser, os := uaParse(reqVo.Ua)
	status := "approved"

	comment := model.Comment{
		Ip:         c.ClientIP(),
		Link:       reqVo.Link,
		Mail:       reqVo.Mail,
		Nick:       reqVo.Nick,
		Comment:    reqVo.Comment,
		Url:        reqVo.Url,
		Pid:        reqVo.Pid,
		Rid:        reqVo.Rid,
		Browser:    browser,
		Os:         os,
		Status:     status,
		InsertedAt: time.Now(),
	}
	db.Create(&comment)
	_ = copier.Copy(&resp, &comment)
	if err := respHandler(&resp); err != nil {
		log.Error(err)
	}

	return &resp, nil
}

func uaParse(ua string) (string, string) {
	client := config.Parser.Parse(ua)
	var browser string
	var os string
	osArgs := make([]string, 0)
	if client.UserAgent.Family != "" {
		browser = util.WriteString(client.UserAgent.Family, client.UserAgent.Major)
	}
	if client.Os.Family != "" {
		osArgs = append(osArgs, client.Os.Family)
		if client.Os.Major != "" {
			osArgs = append(osArgs, " ", client.Os.Major)
		}
	}
	os = util.WriteString(osArgs...)
	return browser, os
}

// 处理IP addr comment
func respHandler(resp *commentVo.CommentRespVo) error {
	resp.ObjectId = resp.ID
	ipInfo, err := service.Ip2Region(resp.Ip)
	if err == nil {
		if ipInfo.City == "" {
			resp.Addr = ipInfo.Province
		} else {
			resp.Addr = ipInfo.City
		}
	}

	return nil
}
