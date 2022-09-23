package models

type BuyerOrdersResponse struct {
	Total       int          `json:"total"`
	PerPage     int          `json:"per_page"`
	CurrentPage int          `json:"current_page"`
	LastPage    int          `json:"last_page"`
	NextPageURL string       `json:"next_page_url"`
	PrevPageURL interface{}  `json:"prev_page_url"`
	From        int          `json:"from"`
	To          int          `json:"to"`
	Data        []BuyerOrder `json:"data"`
}
type BuyerOrder struct {
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
