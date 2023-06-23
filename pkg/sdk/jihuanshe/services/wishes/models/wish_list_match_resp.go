package models

type WishListMatchResultsResp []WishListMatchResult

type WishListMatchResult struct {
	Highlight             int64         `json:"highlight"`
	MatchCards            []MatchCard   `json:"match_cards"`
	NoMatchCards          []NoMatchCard `json:"no_match_cards"`
	SellerAvatar          string        `json:"seller_avatar"`
	SellerCreditRank      string        `json:"seller_credit_rank"`
	SellerPrice           float64       `json:"seller_price"`
	SellerQuantity        int64         `json:"seller_quantity"`
	SellerSettingCity     string        `json:"seller_setting_city"`
	SellerSettingProvince string        `json:"seller_setting_province"`
	SellerUserID          int64         `json:"seller_user_id"`
	SellerUsername        string        `json:"seller_username"`
	ShippingPrice         int64         `json:"shipping_price"`
}

type MatchCard struct {
	CardName  string `json:"card_name"`
	Number    string `json:"number"`
	Price     string `json:"price"`
	ProductID int64  `json:"product_id"`
	Quantity  int64  `json:"quantity"`
	Rarity    string `json:"rarity"`
}

type NoMatchCard struct {
	CardName     string `json:"card_name"`
	Number       string `json:"number"`
	Rarity       string `json:"rarity"`
	WishQuantity int64  `json:"wish_quantity"`
}
