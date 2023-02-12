package models

type ProductsGetResp struct {
	CardID                  int            `json:"card_id"`
	CardNameCn              string         `json:"card_name_cn"`
	CardSubName             string         `json:"card_sub_name"`
	CardVersionNumber       string         `json:"card_version_number"`
	CardVersionRarity       string         `json:"card_version_rarity"`
	CardVersionDefaultImage string         `json:"card_version_default_image"`
	CardVersionImage        string         `json:"card_version_image"`
	UserCardVersionImage    string         `json:"user_card_version_image"`
	Language                string         `json:"language"`
	ProductLanguage         string         `json:"product_language"`
	MinPrice                string         `json:"min_price"`
	AvgPrice                string         `json:"avg_price"`
	Warehouse               bool           `json:"warehouse"`
	Products                []ProductData  `json:"products"`
	DefaultProduct          DefaultProduct `json:"default_product"`
}

type ProductData struct {
	ProductID         int         `json:"product_id"`
	CardNameCn        string      `json:"card_name_cn"`
	Price             float64     `json:"price"`
	Quantity          int         `json:"quantity"`
	Condition         int         `json:"condition"`          // 商品的品相。1: 流通品相，2: 有瑕疵，3: 有损伤
	Remark            string      `json:"remark"`             // 备注
	PublishLocation   interface{} `json:"publish_location"`   // TODO: 这是啥
	CardVersionImage  string      `json:"card_version_image"` // 卡图
	IsDefault         bool        `json:"is_default"`         // 是否为默认商品
	AuthenticatorID   int         `json:"authenticator_id"`   // 评级公司ID
	AuthenticatorName string      `json:"authenticator_name"` // 评级公司名称
	Grading           string      `json:"grading"`            // 评分
}

type DefaultProduct struct {
	ProductID        int         `json:"product_id"`
	CardNameCn       string      `json:"card_name_cn"`
	Price            float64     `json:"price"`
	Quantity         int         `json:"quantity"`
	Condition        int         `json:"condition"`
	Remark           string      `json:"remark"`
	PublishLocation  interface{} `json:"publish_location"`
	CardVersionImage string      `json:"card_version_image"`
	PullOff          bool        `json:"pull_off"`
}
