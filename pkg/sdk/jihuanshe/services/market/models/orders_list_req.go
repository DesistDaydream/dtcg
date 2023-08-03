package models

type OrderListReqQuery struct {
	Page   string `form:"page"`
	Status string `form:"status"`
	Token  string `form:"token"`
}
