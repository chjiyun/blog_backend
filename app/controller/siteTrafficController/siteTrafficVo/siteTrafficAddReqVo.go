package siteTrafficVo

type SiteTrafficAddReqVo struct {
	LinkUrl string `form:"linkUrl" json:"linkUrl" binding:"required"`
	Browser string `form:"browser" json:"browser" biding:"max=50"`
	Os      string `form:"os" json:"os" binding:"max=50"`
	Device  string `form:"device" json:"device" binding:"max=50"`
	NoPv    bool   `form:"noPv" json:"noPv"`
}
