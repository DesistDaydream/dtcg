package database

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

// 添加卡牌价格
func AddCardPirce(cardPrice *models.CardPrice) {
	result := DB.FirstOrCreate(cardPrice, models.CardPrice{CardIDFromDB: cardPrice.CardIDFromDB})
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 更新卡牌价格
func UpdateCardPrice(cardPrice *models.CardPrice, condition map[string]string) {
	// TODO: 如何在 condition 中添加多个条件，然后根据不同情况执行 WHERE
	result := DB.Model(cardPrice).Where("card_id_from_db = ?", cardPrice.CardIDFromDB).Updates(models.CardPrice{
		CardVersionID: cardPrice.CardVersionID,
		MinPrice:      cardPrice.MinPrice,
		AvgPrice:      cardPrice.AvgPrice,
		ImageUrl:      cardPrice.ImageUrl,
	})
	if result.Error != nil {
		logrus.Errorf("更新 %v %v 价格异常: %v", cardPrice.CardIDFromDB, cardPrice.ScName, result.Error)
	}

	logrus.Debugf("已更新 %v 条数据", result.RowsAffected)
}

// 列出所有卡牌价格详情
func ListCardsPrice() (*models.CardsPrice, error) {
	var cp []models.CardPrice
	result := DB.Find(&cp)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.CardsPrice{
		Count:       result.RowsAffected,
		PageSize:    -1,
		PageCurrent: 1,
		PageTotal:   1,
		Data:        cp,
	}, nil
}

// 根据条件获取卡牌价格详情
func GetCardsPrice(pageSize int, pageNum int) (*models.CardsPrice, error) {
	var (
		CardCount int64
		cp        []models.CardPrice
	)

	DB.Model(&models.CardPrice{}).Count(&CardCount)

	// result := DB.Find(&cp)
	result := DB.Limit(pageSize).Offset(pageSize * (pageNum - 1)).Find(&cp)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.CardsPrice{
		Count:       CardCount,
		PageSize:    pageSize,
		PageCurrent: pageNum,
		PageTotal:   (int(CardCount) / pageSize) + 1,
		Data:        cp,
	}, nil
}

// 根据 card_id_from_db 获取卡牌价格详情
func GetCardPrice(cardIDFromDB string) (*models.CardPrice, error) {
	var cardPrice models.CardPrice
	result := DB.Where("card_id_from_db = ?", cardIDFromDB).First(&cardPrice)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cardPrice, nil
}

// 根据 card_version_id 获取卡牌价格详情
func GetCardPriceWhereCardVersionID(cardVersionID string) (*models.CardPrice, error) {
	var cardPrice models.CardPrice
	result := DB.Where("card_version_id = ?", cardVersionID).First(&cardPrice)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cardPrice, nil
}

// 根据卡片集合前缀获取卡牌价格详情
func GetCardPriceWhereSetPrefix(setPrefix string) (*models.CardsPrice, error) {
	var cp []models.CardPrice
	result := DB.Where("set_prefix = ?", setPrefix).Find(&cp)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.CardsPrice{
		Count:       result.RowsAffected,
		PageSize:    -1,
		PageCurrent: 1,
		PageTotal:   1,
		Data:        cp,
	}, nil
}

// 根据条件获取卡牌价格详情
func GetCardPriceByCondition(pageSize int, pageNum int, queryCardPrice *models.QueryCardPrice) (*models.CardsPrice, error) {
	var (
		CardCount int64
		cp        []models.CardPrice
	)

	result := DB.Model(&models.CardPrice{})

	// 多列模糊查询
	if queryCardPrice.Keyword != "" {
		// QField 不为空时，只查询 QField 中的列
		if len(queryCardPrice.QField) > 0 {
			for _, qf := range queryCardPrice.QField {
				result = result.Or(qf+" LIKE ?", "%"+queryCardPrice.Keyword+"%")
			}
		} else {
			result = result.Where("sc_name LIKE ? OR serial LIKE ?",
				"%"+queryCardPrice.Keyword+"%",
				"%"+queryCardPrice.Keyword+"%",
			)
		}
	}

	// 查询最终结果
	result = result.Where(&models.CardPrice{
		AlternativeArt: queryCardPrice.Type,
	})

	// 分页、计数
	result = result.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&cp).Offset(-1).Limit(-1).Count(&CardCount)
	if condition := result.Error; condition != nil {
		return nil, condition
	}

	return &models.CardsPrice{
		Count:       CardCount,
		PageSize:    pageSize,
		PageCurrent: pageNum,
		PageTotal:   (int(CardCount) / pageSize) + 1,
		Data:        cp,
	}, nil
}
