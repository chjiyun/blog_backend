package articleStatVo

// ArticleStatRespVo 文章统计响应Vo
type ArticleStatRespVo struct {
	LinkUrl string `gorm:"size:100;comment:文章url" json:"linkUrl"`
	Pv      int    `gorm:"comment:浏览量" json:"pv"`
}
