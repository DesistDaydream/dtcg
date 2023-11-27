package models

type ReqBodyErrorResp struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

// type CommonReqQuery struct {
// 	// 若想要使用 Gin 的 ShouldBindQuery() 方法绑定 url query，则必须使用 form 标签
// 	// 这是因为 Gin 就是这么规定的，追踪下去会发现 Gin 的代码中，设定了获取 Tag 为 form 的字段，以便将结构体转换为 map 后更好得处理数据。
// 	PageSize int `form:"page_size"`
// 	PageNum  int `form:"page_num"`
// }

// type CommonReqBody struct {
// 	PageSize int `json:"page_size"`
// 	PageNum  int `json:"page_num"`
// }
