package commentVo

type CommentAddReqVo struct {
	Link    string `gorm:"size:100;comment:用户输入网址" json:"link"`
	Mail    string `gorm:"size:100;comment:邮箱" json:"mail"`
	Nick    string `gorm:"size:100;comment:昵称" json:"nick"`
	Comment string `gorm:"comment:评论内容" json:"comment"`
	Pid     uint   `gorm:"comment:父节点" json:"pid"`
	Rid     uint   `gorm:"comment:根节点" json:"rid"`
	Url     string `gorm:"size:255;comment:文章pathname" json:"url"`
	Ua      string `gorm:"size:50;comment:用户的user agent" json:"ua"`
	At      string `gorm:"comment:该条评论回复的评论的作者昵称 user agent" json:"at"`
}
