package models

// 商品信息
type ProductsGetResp struct {
	AvgPrice                string         `json:"avg_price"`
	CardID                  int64          `json:"card_id"`
	CardNameCN              string         `json:"card_name_cn"`
	CardNames               []CardName     `json:"card_names"`
	CardSubName             string         `json:"card_sub_name"`
	CardVersionDefaultImage interface{}    `json:"card_version_default_image"`
	CardVersionImage        string         `json:"card_version_image"`
	CardVersionNumber       string         `json:"card_version_number"`
	CardVersionRarity       string         `json:"card_version_rarity"`
	DefaultProduct          DefaultProduct `json:"default_product"`
	Language                interface{}    `json:"language"`
	MinPrice                string         `json:"min_price"`
	ProductLanguage         string         `json:"product_language"`
	Products                []ProductData  `json:"products"`
	UserCardVersionImage    interface{}    `json:"user_card_version_image"`
	Warehouse               bool           `json:"warehouse"`
}

type CardName struct {
	NameKey   interface{} `json:"name_key"`
	NameValue string      `json:"name_value"`
}

type ProductData struct {
	AuthenticatorID   int         `json:"authenticator_id"`   // 评级公司ID
	AuthenticatorName string      `json:"authenticator_name"` // 评级公司名称
	CardNameCN        string      `json:"card_name_cn"`       // 卡牌中文名称
	CardVersionImage  string      `json:"card_version_image"` // 卡图
	Condition         int         `json:"condition"`          // 商品的品相。1: 流通品相，2: 有瑕疵，3: 有损伤
	Grading           string      `json:"grading"`            // 评分
	IsDefault         bool        `json:"is_default"`         // 是否为默认商品
	Price             float64     `json:"price"`              // 商品价格
	ProductID         int         `json:"product_id"`         // 商品 ID
	PublishLocation   interface{} `json:"publish_location"`   // TODO: 这是啥
	Quantity          int         `json:"quantity"`           // 商品数量
	Remark            string      `json:"remark"`             // 商品备注
}

type DefaultProduct struct {
	CardNameCN       string      `json:"card_name_cn"`       // 卡牌中文名称
	CardVersionImage string      `json:"card_version_image"` // 卡图
	Condition        int64       `json:"condition"`          // 商品的品相。1: 流通品相，2: 有瑕疵，3: 有损伤
	Price            float64     `json:"price"`              // 商品价格
	ProductID        int64       `json:"product_id"`         // 商品 ID
	PublishLocation  interface{} `json:"publish_location"`   // TODO: 这是啥
	PullOff          bool        `json:"pull_off"`           // 是否已下架
	Quantity         int64       `json:"quantity"`           // 商品数量
	Remark           string      `json:"remark"`             // 商品备注
}
