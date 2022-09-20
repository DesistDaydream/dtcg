package models

type SellerOrderProducts struct {
	OrderID                 int                  `json:"order_id"`
	OrderUUID               string               `json:"order_uuid"`
	OrderName               string               `json:"order_name"`
	OrderImage              string               `json:"order_image"`
	Status                  string               `json:"status"`
	ReturnGoodsStatus       int                  `json:"return_goods_status"`
	Remark                  string               `json:"remark"`
	ReceiveLeftSecond       int                  `json:"receive_left_second"`
	CanDeferReceiveDeadline bool                 `json:"can_defer_receive_deadline"`
	CreatedAt               string               `json:"created_at"`
	PayDeadlineLeft         interface{}          `json:"pay_deadline_left"`
	TotalPrice              float64              `json:"total_price"`
	ProductPrice            string               `json:"product_price"`
	ShippingPrice           string               `json:"shipping_price"`
	HandlingFee             float64              `json:"handling_fee"`
	TechFeeRate             string               `json:"tech_fee_rate"`
	TransFeeRate            string               `json:"trans_fee_rate"`
	ActualTotalPrice        float64              `json:"actual_total_price"`
	ProductPriceSnapshot    string               `json:"product_price_snapshot"`
	ShippingPriceSnapshot   string               `json:"shipping_price_snapshot"`
	Name                    string               `json:"name"`
	Phone                   string               `json:"phone"`
	Province                string               `json:"province"`
	City                    string               `json:"city"`
	District                string               `json:"district"`
	Address                 string               `json:"address"`
	Postcode                string               `json:"postcode"`
	ExpressType             int                  `json:"express_type"`
	ExpressName             string               `json:"express_name"`
	ExpressNumber           string               `json:"express_number"`
	ExpressOrderCode        interface{}          `json:"express_order_code"`
	BuyerUserID             int                  `json:"buyer_user_id"`
	BuyerUsername           string               `json:"buyer_username"`
	BuyerUserAvatar         string               `json:"buyer_user_avatar"`
	ProductQuantity         int                  `json:"product_quantity"`
	OrderProducts           []SellerOrderProduct `json:"order_products"`
}
type SellerOrderProduct struct {
	Price                  string      `json:"price"`
	Quantity               int         `json:"quantity"`
	ProductGameKey         string      `json:"product_game_key"`
	ProductGameSubKey      string      `json:"product_game_sub_key"`
	CardNameCn             string      `json:"card_name_cn"`
	CardVersionNumber      string      `json:"card_version_number"`
	CardVersionRarity      string      `json:"card_version_rarity"`
	CardVersionImage       string      `json:"card_version_image"`
	ProductCondition       int         `json:"product_condition"`
	ProductRemark          string      `json:"product_remark"`
	ProductPublishLocation interface{} `json:"product_publish_location"`
	ProductLanguage        string      `json:"product_language"`
}
