package database

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

// 添加卡牌描述
func AddCardDesc(cardDesc *models.CardDesc) {
	result := DB.FirstOrCreate(cardDesc, cardDesc)
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 更新卡牌描述
func UpdateCardDesc(cardDesc *models.CardDesc, condition map[string]string) {
	// TODO: 如何在 condition 中添加多个条件，然后根据不同情况执行 WHERE
	result := DB.Model(cardDesc).Where("card_id_from_db = ?", cardDesc.CardIDFromDB).Updates(cardDesc)
	if result.Error != nil {
		logrus.Errorf("更新 %v %v 价格异常: %v", cardDesc.CardIDFromDB, cardDesc.ScName, result.Error)
	}

	logrus.Debugf("已更新 %v 条数据", result.RowsAffected)
}

// 获取所有卡片描述
func ListCardDesc() (*models.CardsDesc, error) {
	var cd []models.CardDesc
	result := DB.Find(&cd)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.CardsDesc{
		Count:       result.RowsAffected,
		PageSize:    -1,
		PageCurrent: 1,
		PageTotal:   1,
		Data:        cd,
	}, nil
}

// 分页获取卡牌描述
func GetCardDesc(pageSize int, pageNum int) (*models.CardsDesc, error) {
	var (
		CardCount int64
		cd        []models.CardDesc
	)

	DB.Model(&models.CardDesc{}).Count(&CardCount)

	result := DB.Limit(pageSize).Offset(pageSize * (pageNum - 1)).Find(&cd)
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

// 根据 card_id_from_db 获取卡片描述
func GetCardDescByCardIDFromDB(cardIDFromDB string) (*models.CardDesc, error) {
	var cardDesc models.CardDesc
	result := DB.Where("card_id_from_db = ?", cardIDFromDB).First(&cardDesc)
	if condition := result.Error; condition != nil {
		return nil, condition
	}

	return &cardDesc, nil
}
