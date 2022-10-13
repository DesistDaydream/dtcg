package database

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

func AddCardSet(cardGroup *models.CardSet) {
	result := DB.FirstOrCreate(cardGroup, cardGroup)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 获取所有卡包
func ListCardSets() (*models.CardSets, error) {
	var cg []models.CardSet
	result := DB.Find(&cg)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.CardSets{
		Count: result.RowsAffected,
		Data:  cg,
	}, nil
}
