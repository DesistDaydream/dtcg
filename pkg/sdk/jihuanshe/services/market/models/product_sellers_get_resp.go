package models

type ProductSellersGetResp struct {
	CurrentPage int64            `json:"current_page"`
	Data        []ProductSellers `json:"data"`
	From        int64            `json:"from"`
	LastPage    int64            `json:"last_page"`
	NextPageURL string           `json:"next_page_url"`
	PerPage     int64            `json:"per_page"`
	PrevPageURL string           `json:"prev_page_url"`
	To          int64            `json:"to"`
	Total       int64            `json:"total"`
}

type ProductSellers struct {
	CardVersionImage string      `json:"card_version_image"`
	EcommerceVerify  bool        `json:"ecommerce_verify"`
	MinPrice         string      `json:"min_price"`
	Quantity         string      `json:"quantity"`
	SellerCity       string      `json:"seller_city"`
	SellerCreditRank string      `json:"seller_credit_rank"`
	SellerProvince   string      `json:"seller_province"`
	SellerUserAvatar string      `json:"seller_user_avatar"`
	SellerUserID     int64       `json:"seller_user_id"`
	SellerUsername   string      `json:"seller_username"`
	VerifyStatus     interface{} `json:"verify_status"`
}
