package database

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

func AddCardPirce(cardPrice *models.CardPrice) {
	result := DB.FirstOrCreate(cardPrice, models.CardPrice{CardIDFromDB: cardPrice.CardIDFromDB})
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

func UpdateCardPrice(cardPrice *models.CardPrice, condition map[string]string) {
	// TODO: 如何在 condition 中添加多个条件，然后根据不同情况执行 WHERE
	result := DB.Model(cardPrice).Where("card_id_from_db = ?", cardPrice.CardIDFromDB).Updates(models.CardPrice{
		MinPrice: cardPrice.MinPrice,
		AvgPrice: cardPrice.AvgPrice,
	})
	if result.Error != nil {
		logrus.Errorf("更新 %v %v 价格异常: %v", cardPrice.CardIDFromDB, cardPrice.ScName, result.Error)
	}

	logrus.Debugf("已更新 %v 条数据", result.RowsAffected)
}
