package models

type CardPriceGetResp struct {
	Data    CardPriceData `json:"data"`
	Message string        `json:"message"`
	Success bool          `json:"success"`
}

type CardPriceData struct {
	AvgPrice      string        `json:"avg_price"`
	CardID        string        `json:"card_id"`
	CardVersionID int           `json:"cv_id"`
	Data          []ProductData `json:"data"`
	Enabled       bool          `json:"enabled"`
	Link          string        `json:"link"`
	Total         int64         `json:"total"`
}

type ProductData struct {
	CardVersionID    int    `json:"card_version_id"`
	MinPrice         string `json:"min_price"`
	Quantity         string `json:"quantity"`
	SellerCity       string `json:"seller_city"`
	SellerProvince   string `json:"seller_province"`
	SellerUserAvatar string `json:"seller_user_avatar"`
	SellerUserID     int64  `json:"seller_user_id"`
	SellerUsername   string `json:"seller_username"`
	VerifyStatus     int64  `json:"verify_status"`
}
