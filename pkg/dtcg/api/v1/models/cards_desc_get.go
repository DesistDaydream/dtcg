package models

type GetCardsDescReqQuery struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}
