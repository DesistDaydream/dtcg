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

// 获取所有卡牌集合
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

// 分页获取卡牌集合
func GetCardSets(pageSize int, pageNum int) (*models.CardSets, error) {
	var (
		SetCount int64
		cs       []models.CardSet
	)

	DB.Model(&models.CardSet{}).Count(&SetCount)

	result := DB.Limit(pageSize).Offset(pageSize * (pageNum - 1)).Find(&cs)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.CardSets{
		Count:       SetCount,
		PageSize:    pageSize,
		PageCurrent: pageNum,
		PageTotal:   (int(SetCount) / pageSize) + 1,
		Data:        cs,
	}, nil
}
