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
	CardsIDFromDB []string `json:"cards_id_from_db"`
}
