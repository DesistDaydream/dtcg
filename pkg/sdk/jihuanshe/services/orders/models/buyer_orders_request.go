package models

type BuyerOrdersQuery struct {
	Page   string `query:"page"`
	Status string `query:"status"`
	Token  string `query:"token"`
}
