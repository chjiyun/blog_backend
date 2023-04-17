package articleService

import (
	"blog_backend/app/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateArticlePv 浏览量自动 +1 后期要加加锁，防止高并发
func UpdateArticlePv(c *gin.Context, linkUrl string) (int, error) {
	db := c.Value("DB").(*gorm.DB)

	var data model.ArticleStat
	var pv int
	tx := db.Where("link_url = ?", linkUrl).Find(&data)
	if tx.Error != nil {
		return pv, tx.Error
	}
	if tx.RowsAffected == 0 {
		pv = 1
		data = model.ArticleStat{
			LinkUrl: linkUrl,
			Pv:      1,
		}
		db.Create(&data)
	} else {
		pv = data.Pv + 1
		db.Model(&data).Update("pv", pv)
	}
	return pv, nil
}

func GetManyArticlePv(c *gin.Context, linkUrls []string) ([]model.ArticleStat, error) {
	db := c.Value("DB").(*gorm.DB)

	var articleStats []model.ArticleStat
	tx := db.Where("link_url in ?", linkUrls).Find(&articleStats)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return articleStats, nil
}
