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

// 列出所有卡牌描述
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
func GetCardsDesc(pageSize int, pageNum int) (*models.CardsDesc, error) {
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

// 根据条件获取卡牌描述
func GetCardDescByCondition(pageSize int, pageNum int, queryCardDesc *models.QueryCardDesc) (*models.CardsDesc, error) {
	var (
		CardCount int64
		cd        []models.CardDesc
	)

	result := DB.Model(&models.CardDesc{})

	// 通过关键字在多列模糊查询
	if queryCardDesc.Keyword != "" {
		// QField 不为空时，只查询 QField 中的列
		if len(queryCardDesc.QField) > 0 {
			// TODO: 每次循环都会生成一个 SQL 片段，导致查询结果异常。并且 Or 会导致后面所有的查询条件都会使用 Or()，导致查询结果异常。
			// 这么做无法通过 Group 功能，将这些 Or 放到一个括号中，然后再与其他查询条件进行 And
			for _, field := range queryCardDesc.QField {
				result = result.Or(field+" LIKE ?", "%"+queryCardDesc.Keyword+"%")
			}
		} else {
			result = result.Where("sc_name LIKE ? OR serial LIKE ? OR effect LIKE ? OR evo_cover_effect LIKE ? OR security_effect LIKE ?",
				"%"+queryCardDesc.Keyword+"%",
				"%"+queryCardDesc.Keyword+"%",
				"%"+queryCardDesc.Keyword+"%",
				"%"+queryCardDesc.Keyword+"%",
				"%"+queryCardDesc.Keyword+"%",
			)
			// 检查查询语句
			logrus.Debugf("检查关键字多列模糊查询SQL: %v", result.Statement.SQL.String())
		}
	}

	// 颜色多匹配查询
	if len(queryCardDesc.Color) > 0 {
		for _, color := range queryCardDesc.Color {
			// TODO: 这里不能用 Where，否则所有的颜色匹配都会是 AND 逻辑
			result = result.Where("color LIKE ?", "%"+color+"%")
			// TODO: 这里如果使用 Or 会与上面的关键字查询产生相同的问题，无法通过 Group 功能，将这些 Or 放到一个括号中，然后再与其他查询条件进行 And
			// 每次循环都会生成一个 SQL 片段，导致查询结果异常。并且 Or 会导致后面所有的查询条件都会使用 Or()，导致查询结果异常
			// result = result.Or("color LIKE ?", "%"+color+"%")
		}
	}

	// 稀有度多匹配查询
	if len(queryCardDesc.Rarity) > 0 {
		result = result.Where("rarity IN ?", queryCardDesc.Rarity)
	}

	// 查询最终结果
	result = result.Where(&models.CardDesc{
		Type: queryCardDesc.Type,
	})

	// 分页、计数
	result = result.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Debug().Find(&cd).Offset(-1).Limit(-1).Count(&CardCount)
	if condition := result.Error; condition != nil {
		return nil, condition
	}

	return &models.CardsDesc{
		Count:       CardCount,
		PageSize:    pageSize,
		PageCurrent: pageNum,
		PageTotal:   (int(CardCount) / pageSize) + 1,
		Data:        cd,
	}, nil
}
