package models

type PostDeckPriceRequest struct {
	Deck  string `json:"deck"`
	Envir string `json:"envir"`
}

type PostDeckPriceResponse struct {
	MinPrice float64        `json:"min_price"`
	AvgPrice float64        `json:"avg_price"`
	Data     []MutCardPrice `json:"data"`
}

type MutCardPrice struct {
	Count          int     `json:"count"`
	Serial         string  `json:"serial"`
	ScName         string  `json:"sc_name"`
	AlternativeArt string  `json:"alternative_art"`
	MinPrice       float64 `json:"min_price"`
	AvgPrice       float64 `json:"avg_price"`
}
