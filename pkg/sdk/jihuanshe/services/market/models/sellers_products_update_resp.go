package models

type ProductsUpdateResp struct {
	Message        string `json:"message"`
	OnSaleQuantity int    `json:"on_sale_quantity"`
}

type ProductUpdateErrorResp struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

// 市场商品不存在
// {
//     "code": 229,
//     "error": "MARKET_PRODUCT_NOT_EXISTS",
//     "msg": "抱歉，该商品已删除"
// }
