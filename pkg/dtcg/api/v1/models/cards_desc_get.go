package models

type GetCardsDescReqQuery struct {
	// 若想要使用 Gin 的 ShouldBindQuery() 方法绑定 url query，则必须使用 form 标签
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}
