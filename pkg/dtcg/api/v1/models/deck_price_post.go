package models

type PostDeckPriceWithJSONReqBody struct {
	Deck  string `json:"deck"`
	Envir string `json:"envir"`
}

type PostDeckPriceResp struct {
	MinPrice string                  `json:"min_price"`
	AvgPrice string                  `json:"avg_price"`
	Data     []PostDeckPriceRespData `json:"data"`
}

type PostDeckPriceRespData struct {
	CardIDFromDB   int    `json:"card_id_from_db"`
	Count          int    `json:"count"`
	Serial         string `json:"serial"`
	ScName         string `json:"sc_name"`
	AlternativeArt string `json:"alternative_art"`
	MinPrice       string `json:"min_price"`
	AvgPrice       string `json:"avg_price"`
	MinUnitPrice   string `json:"min_unit_price"`
	AvgUnitPrice   string `json:"avg_unit_price"`
}

type PostDeckPriceWithIDReq struct {
	IDs string `json:"ids"`
}

// 通过 PostDeckPriceWithIDReq 中的 IDs 字段获取卡牌信息
// 就是将数组字符串转换一下，然后再获取卡牌信息
type PostDeckPriceWithIDReqTransform struct {
	CardsInfo []CardInfo `json:"cards_info"`
}

type CardInfo struct {
	Count        int    `json:"count"`
	CardIDFromDB string `json:"card_id_from_db"`
	ScName       string `json:"sc_name"`
	Serial       string `json:"serial"`
}
