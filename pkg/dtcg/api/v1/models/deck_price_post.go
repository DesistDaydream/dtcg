package models

type PostDeckPriceRequest struct {
	Deck  string `json:"deck"`
	Envir string `json:"envir"`
}

type PostDeckPriceResponse struct {
	MinPrice string         `json:"min_price"`
	AvgPrice string         `json:"avg_price"`
	Data     []MutCardPrice `json:"data"`
}

type MutCardPrice struct {
	Count          int    `json:"count"`
	Serial         string `json:"serial"`
	ScName         string `json:"sc_name"`
	AlternativeArt string `json:"alternative_art"`
	MinPrice       string `json:"min_price"`
	AvgPrice       string `json:"avg_price"`
}

type PostDeckPriceWithIDReq struct {
	IDs string `json:"ids"`
}

type PostDeckPriceWithIDReqTransform struct {
	CardsInfo []CardInfo `json:"cards_info"`
}

type CardInfo struct {
	Count        int    `json:"count"`
	CardIDFromDB string `json:"card_id_from_db"`
	ScName       string `json:"sc_name"`
	Serial       string `json:"serial"`
}

type PostDeckPriceWithIDResp struct{}
