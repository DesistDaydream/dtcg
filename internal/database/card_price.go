package database

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

func AddCardPirce(cardPrice *models.CardPrice) {
	result := db.FirstOrCreate(cardPrice, models.CardPrice{CardID: cardPrice.CardID})
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

func UpdateCardPrice() {

}
