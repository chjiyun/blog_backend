package result

import "blog_backend/app/common"

type WlResult struct {
	Code int         `json:"errno"`
	Msg  string      `json:"errmsg"`
	Data interface{} `json:"data"`
}

func NewWl() *WlResult {
	r := &WlResult{}
	r.Code = 0
	r.Msg = common.Success.String()
	return r
}

func (r *WlResult) Success(data interface{}) *WlResult {
	r.Code = 0
	r.Data = data
	return r
}

func (r *WlResult) FailErr(err error) *WlResult {
	r.Code = 1000
	r.Msg = err.Error()
	return r
}
