package models

type ProductsListResp struct {
	CurrentPage int               `json:"current_page"`
	Data        []ProductListData `json:"data"`
	From        int               `json:"from"`
	LastPage    int               `json:"last_page"`
	NextPageURL string            `json:"next_page_url"`
	Path        string            `json:"path"`
	PerPage     int               `json:"per_page"`
	PrevPageURL string            `json:"prev_page_url"`
	To          int               `json:"to"`
	Total       int               `json:"total"`
}

type ProductListData struct {
	ProductID         int    `json:"product_id"`
	Price             string `json:"price"`
	Quantity          int    `json:"quantity"`
	OnSale            int    `json:"on_sale"`
	Condition         int    `json:"condition"`
	Remark            string `json:"remark"`
	CardID            int    `json:"card_id"`
	CardVersionID     int    `json:"card_version_id"`
	CardNameCn        string `json:"card_name_cn"`
	CardNameCnCnocg   string `json:"card_name_cn_cnocg"`
	CardVersionNumber string `json:"card_version_number"`
	NumberAlias       string `json:"number_alias"`
	CardVersionRarity string `json:"card_version_rarity"`
	CardVersionImage  string `json:"card_version_image"`
	ProductLanguage   string `json:"product_language"`
	MinPrice          string `json:"min_price"`
	AvgPrice          string `json:"avg_price"`
	AuthenticatorID   int    `json:"authenticator_id"`
	AuthenticatorName string `json:"authenticator_name"`
	Grading           string `json:"grading"`
}
