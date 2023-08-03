package models

type ReqBodyErrorResp struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

type CommonReqQuery struct {
	// 若想要使用 Gin 的 ShouldBindQuery() 方法绑定 url query，则必须使用 form 标签
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}

type CommonReqBody struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}
