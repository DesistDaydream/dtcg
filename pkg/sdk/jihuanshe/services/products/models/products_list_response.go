package models

type ProductsListResponse struct {
	Total       int           `json:"total"`
	PerPage     int           `json:"per_page"`
	CurrentPage int           `json:"current_page"`
	LastPage    int           `json:"last_page"`
	NextPageURL string        `json:"next_page_url"`
	PrevPageURL interface{}   `json:"prev_page_url"`
	From        int           `json:"from"`
	To          int           `json:"to"`
	Data        []ProductList `json:"data"`
}

type ProductList struct {
	ProductID         int         `json:"product_id"`
	Price             string      `json:"price"`
	Quantity          int         `json:"quantity"`
	OnSale            int         `json:"on_sale"`
	Condition         int         `json:"condition"`
	Remark            string      `json:"remark"`
	CardID            int         `json:"card_id"`
	CardVersionID     int         `json:"card_version_id"`
	CardNameCn        string      `json:"card_name_cn"`
	CardNameCnCnocg   interface{} `json:"card_name_cn_cnocg"`
	CardVersionNumber string      `json:"card_version_number"`
	NumberAlias       string      `json:"number_alias"`
	CardVersionRarity string      `json:"card_version_rarity"`
	CardVersionImage  string      `json:"card_version_image"`
	ProductLanguage   string      `json:"product_language"`
	MinPrice          string      `json:"min_price"`
	AvgPrice          string      `json:"avg_price"`
}
