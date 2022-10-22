package models

type ProductsGetResp struct {
	CardID               int           `json:"card_id"`
	CardNameCn           string        `json:"card_name_cn"`
	CardSubName          string        `json:"card_sub_name"`
	CardVersionNumber    string        `json:"card_version_number"`
	CardVersionRarity    string        `json:"card_version_rarity"`
	CardVersionImage     string        `json:"card_version_image"`
	UserCardVersionImage string        `json:"user_card_version_image"`
	Language             interface{}   `json:"language"`
	ProductLanguage      string        `json:"product_language"`
	MinPrice             string        `json:"min_price"`
	AvgPrice             string        `json:"avg_price"`
	Warehouse            bool          `json:"warehouse"`
	Products             []ProductData `json:"products"`
}

type ProductData struct {
	ProductID       int         `json:"product_id"`
	CardNameCn      string      `json:"card_name_cn"`
	Price           float32     `json:"price"`
	Quantity        int         `json:"quantity"`
	Condition       int         `json:"condition"`
	Remark          string      `json:"remark"`
	PublishLocation interface{} `json:"publish_location"`
}
