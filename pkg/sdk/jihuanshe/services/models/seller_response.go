package models

type SellerOrders struct {
	Total       int           `json:"total"`
	PerPage     int           `json:"per_page"`
	CurrentPage int           `json:"current_page"`
	LastPage    int           `json:"last_page"`
	NextPageURL string        `json:"next_page_url"`
	PrevPageURL interface{}   `json:"prev_page_url"`
	From        int           `json:"from"`
	To          int           `json:"to"`
	Data        []SellerOrder `json:"data"`
}
type SellerOrder struct {
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
