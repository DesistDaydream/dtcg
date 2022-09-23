package models

type BuyerOrderProductsRequest struct {
	Token string `query:"token"`
}

type SellerOrderProductsRequest struct {
	Token string `query:"token"`
}
