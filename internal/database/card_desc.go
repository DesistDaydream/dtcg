package database

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

func AddCardDesc(cardDesc *models.CardDesc) {
	result := db.FirstOrCreate(cardDesc, cardDesc)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 获取所有卡片描述
func ListCardDesc() (*models.CardsDesc, error) {
	var cd []models.CardDesc
	result := db.Find(&cd)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.CardsDesc{
		Count: result.RowsAffected,
		Data:  cd,
	}, nil
}

// 根据条件获取卡片描述
func GetCardDesc(pageSize int, pageNum int) (*models.CardsDesc, error) {
	var (
		CardCount int64
		cd        []models.CardDesc
	)

	db.Model(&models.CardDesc{}).Count(&CardCount)

	result := db.Limit(pageSize).Offset(pageSize * (pageNum - 1)).Find(&cd)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.CardsDesc{
		Count:       CardCount,
		PageSize:    pageSize,
		PageCurrent: pageNum,
		PageTotal:   (int(CardCount) / pageSize) + 1,
		Data:        cd,
	}, nil
}
