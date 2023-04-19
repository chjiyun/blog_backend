package siteTrafficVo

type SiteTrafficAddReqVo struct {
	LinkUrl string `form:"linkUrl" json:"linkUrl" binding:"required"`
	Ua      string `form:"ua" json:"ua"`
	NoPv    bool   `form:"noPv" json:"noPv"`
}
