package models

type PostCardsPriceReqQuery struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}
