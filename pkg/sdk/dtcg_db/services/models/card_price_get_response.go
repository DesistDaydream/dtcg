package models

type CardPriceGetResponse struct {
	Data    CardPrice `json:"data"`
	Message string    `json:"message"`
	Success bool      `json:"success"`
}

type CardPrice struct {
	AvgPrice string    `json:"avg_price"`
	CardID   string    `json:"card_id"`
	Products []Product `json:"data"`
	Enabled  bool      `json:"enabled"`
	Total    int64     `json:"total"`
}

type Product struct {
	CardVersionID    int64  `json:"card_version_id"`
	MinPrice         string `json:"min_price"`
	Quantity         string `json:"quantity"`
	SellerCity       string `json:"seller_city"`
	SellerProvince   string `json:"seller_province"`
	SellerUserAvatar string `json:"seller_user_avatar"`
	SellerUserID     int64  `json:"seller_user_id"`
	SellerUsername   string `json:"seller_username"`
	VerifyStatus     *int64 `json:"verify_status"`
}
