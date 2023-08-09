package models

// 卡牌基础信息
type CardVersionsPriceHistoryResp []RequestElement

// Request
type RequestElement struct {
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}
