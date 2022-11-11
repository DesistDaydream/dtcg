package models

type CardPriceGetResp struct {
	Data    CardPriceGetData `json:"data"`
	Message string           `json:"message"`
	Success bool             `json:"success"`
}

type CardPriceGetData struct {
	AvgPrice string    `json:"avg_price"`
	CardID   string    `json:"card_id"`
	Products []Product `json:"data"`
	Enabled  bool      `json:"enabled"`
	Total    int64     `json:"total"`
	// 新增了两个字段
	CardVersionID int    `json:"cv_id"`
	Link          string `json:"link"`
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
	VerifyStatus     int64  `json:"verify_status"`
}
