package database

import "github.com/sirupsen/logrus"

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

func AddCardPirce(cardPrice *CardPrice) {
	result := db.FirstOrCreate(cardPrice, CardPrice{CardID: cardPrice.CardID})
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

func UpdateCardPrice() {

}
