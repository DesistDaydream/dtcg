package models

type ProductsListResp struct {
	CurrentPage int           `json:"current_page"`
	Data        []ProductData `json:"data"`
	From        int           `json:"from"`
	LastPage    int           `json:"last_page"`
	NextPageURL string        `json:"next_page_url"`
	Path        string        `json:"path"`
	PerPage     int           `json:"per_page"`
	PrevPageURL string        `json:"prev_page_url"`
	To          int           `json:"to"`
	Total       int           `json:"total"`
}

type ProductData struct {
	AuthenticatorID   string     `json:"authenticator_id"`    // 评级公司ID
	AuthenticatorName string     `json:"authenticator_name"`  // 评级公司名称
	AvgPrice          string     `json:"avg_price"`           // 集换价
	CardID            int        `json:"card_id"`             // ？
	CardNameCN        string     `json:"card_name_cn"`        // ？
	CardNameCNCnocg   string     `json:"card_name_cn_cnocg"`  // ？
	CardNames         []CardName `json:"card_names"`          // ？
	CardVersionID     int        `json:"card_version_id"`     // 卡牌ID。集换社中关于卡牌的唯一标识符
	CardVersionImage  string     `json:"card_version_image"`  // 商品图片
	CardVersionNumber string     `json:"card_version_number"` // 卡牌编号
	CardVersionRarity string     `json:"card_version_rarity"` // 卡牌稀有度。C、U、R、SR、SEC、异画。
	Condition         int        `json:"condition"`           // 商品的品相。1: 流通品相，2: 有瑕疵，3: 有损伤
	Grading           string     `json:"grading"`             // 评分。仅适用于评级卡牌
	MinPrice          string     `json:"min_price"`           // 最低价
	NumberAlias       string     `json:"number_alias"`        // 编号别名
	OnSale            int        `json:"on_sale"`             // 售卖状态。1: 在售，0: 下架
	Price             string     `json:"price"`               // 商品售卖价格
	ProductID         int        `json:"product_id"`          // 商品ID
	ProductLanguage   string     `json:"product_language"`    // 商品语言。简、日、等等
	Quantity          int        `json:"quantity"`            // 售卖数量。注意：评级卡商品每次只能上架一张
	Remark            string     `json:"remark"`              // 商品备注信息
}

type CardName struct {
	NameKey   string `json:"name_key"`
	NameValue string `json:"name_value"`
}
