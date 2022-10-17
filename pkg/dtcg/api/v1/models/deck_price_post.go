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
	Count          int
	Serial         string
	ScName         string
	AlternativeArt string
	MinPrice       float64
	AvgPrice       float64
}
