package models

type CardsPrice struct {
	Count int64       `json:"count"`
	Data  []CardPrice `json:"data"`
}

type CardPrice struct {
	CardID        int     `json:"card_id"`
	CardVersionID int     `json:"card_version_id"`
	MinPrice      float64 `json:"min_price"`
	AvgPrice      float64 `json:"avg_price"`
}
