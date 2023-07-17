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
	AuthenticatorID   string     `json:"authenticator_id"`
	AuthenticatorName string     `json:"authenticator_name"`
	AvgPrice          string     `json:"avg_price"`
	CardID            int        `json:"card_id"`
	CardNameCN        string     `json:"card_name_cn"`
	CardNameCNCnocg   string     `json:"card_name_cn_cnocg"`
	CardNames         []CardName `json:"card_names"`
	CardVersionID     int        `json:"card_version_id"`
	CardVersionImage  string     `json:"card_version_image"`
	CardVersionNumber string     `json:"card_version_number"`
	CardVersionRarity string     `json:"card_version_rarity"`
	Condition         int        `json:"condition"` // 商品的品相。1: 流通品相，2: 有瑕疵，3: 有损伤
	Grading           string     `json:"grading"`
	MinPrice          string     `json:"min_price"`
	NumberAlias       string     `json:"number_alias"`
	OnSale            int        `json:"on_sale"` // 售卖状态。1: 在售，0: 下架
	Price             string     `json:"price"`
	ProductID         int        `json:"product_id"`
	ProductLanguage   string     `json:"product_language"`
	Quantity          int        `json:"quantity"`
	Remark            string     `json:"remark"`
}

type CardName struct {
	NameKey   string `json:"name_key"`
	NameValue string `json:"name_value"`
}
