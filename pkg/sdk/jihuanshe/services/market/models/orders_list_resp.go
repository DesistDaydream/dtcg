package models

type BuyerOrdersListResp struct {
	CurrentPage int              `json:"current_page"`
	Data        []BuyerOrderData `json:"data"`
	From        int              `json:"from"`
	LastPage    int              `json:"last_page"`
	NextPageURL string           `json:"next_page_url"`
	Path        string           `json:"path"`
	PerPage     int              `json:"per_page"`
	PrevPageURL string           `json:"prev_page_url"`
	To          int              `json:"to"`
	Total       int              `json:"total"`
}
type BuyerOrderData struct {
	OrderID           int     `json:"order_id"`
	OrderUUID         string  `json:"order_uuid"`
	OrderName         string  `json:"order_name"`
	OrderImage        string  `json:"order_image"`
	TotalPrice        float64 `json:"total_price"`
	Status            string  `json:"status"`
	ReturnGoodsStatus int     `json:"return_goods_status"`
	SellerUserID      int     `json:"seller_user_id"`
	SellerUsername    string  `json:"seller_username"`
	SellerUserAvatar  string  `json:"seller_user_avatar"`
	IsWarehouse       bool    `json:"is_warehouse"`
	CreatedAt         string  `json:"created_at"`
}

type SellerOrderListResp struct {
	CurrentPage int               `json:"current_page"`
	Data        []SellerOrderData `json:"data"`
	From        int               `json:"from"`
	LastPage    int               `json:"last_page"`
	NextPageURL string            `json:"next_page_url"`
	Path        string            `json:"path"`
	PerPage     int               `json:"per_page"`
	PrevPageURL string            `json:"prev_page_url"`
	To          int               `json:"to"`
	Total       int               `json:"total"`
}

type SellerOrderData struct {
	OrderID           int     `json:"order_id"`
	OrderUUID         string  `json:"order_uuid"`
	OrderName         string  `json:"order_name"`
	OrderImage        string  `json:"order_image"`
	TotalPrice        float64 `json:"total_price"`
	Status            string  `json:"status"`
	ReturnGoodsStatus int     `json:"return_goods_status"`
	BuyerUserID       int     `json:"buyer_user_id"`
	BuyerUsername     string  `json:"buyer_username"`
	BuyerUserAvatar   string  `json:"buyer_user_avatar"`
	CreatedAt         string  `json:"created_at"`
}
