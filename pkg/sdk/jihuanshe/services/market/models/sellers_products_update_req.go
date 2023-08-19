package models

type ProductsUpdateReqBody struct {
	AuthenticatorID         string `json:"authenticator_id"`           // 评级公司ID
	Grading                 string `json:"grading"`                    // 评分
	Condition               string `json:"condition"`                  // 商品的品相。1: 流通品相，2: 有瑕疵，3: 有损伤，4: 评级卡
	Default                 string `json:"default"`                    // 是否设为默认商品。1: 是; 0: 否
	OnSale                  string `json:"on_sale"`                    // 售卖状态。1: 在售，0: 下架
	Price                   string `json:"price"`                      // 售卖价格
	ProductCardVersionImage string `json:"product_card_version_image"` // 商品图片
	Quantity                string `json:"quantity"`                   // 售卖数量。注意：评级卡商品每次只能上架一张
	Remark                  string `json:"remark"`                     // 商品备注信息
}
