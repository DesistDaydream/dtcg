package database

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 添加卡牌描述
func AddCardDesc(cardDesc *models.CardDesc) {
	// 若 card_id_from_db 已存在，则不添加。
	// 毕竟是从 https://digimon.card.moe/ 爬取的数据，不会有重复的 card_id_from_db，如果对方的 ID 已经在本地了，那就不需要再添加了。
	// TODO: 但是如果对方的数据有更新，那么本地的数据就会不同步了，需要有一个更新的机制。
	result := DB.FirstOrCreate(cardDesc, &models.CardDesc{CardIDFromDB: cardDesc.CardIDFromDB})
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
func GetCardDescByCondition(pageSize int, pageNum int, queryCardDesc *models.CardDescQuery) (*models.CardsDesc, error) {
	var (
		CardCount int64
		cd        []models.CardDesc
	)

	result := DB.Model(&models.CardDesc{})

	// 卡牌集合匹配查询
	if queryCardDesc.CardSet != 0 {
		result = result.Where("set_id = ?", queryCardDesc.CardSet)
	}

	// 通过关键字在多列模糊查询
	if queryCardDesc.Keyword != "" {
		// QField 不为空时，只查询 QField 中的列
		if len(queryCardDesc.QField) > 0 {
			// TODO: 每次循环都会生成一个 SQL 片段，导致查询结果异常。并且 Or 会导致后面所有的查询条件都会使用 Or()，导致查询结果异常。
			// 这么做无法通过 Group 功能，将这些 Or 放到一个括号中，然后再与其他查询条件进行 And
			// for _, field := range queryCardDesc.QField {
			// 	result = result.Or(field+" LIKE ?", "%"+queryCardDesc.Keyword+"%")
			// }

			// 通过这种方式解决了？参考：https://github.com/go-gorm/gorm/issues/5052
			f := func(queryCardDesc *models.CardDescQuery, result *gorm.DB) *gorm.DB {
				// 通过 Session() 创建一个新的 DB 实例，避免影响原来的 DB 实例。用以实现为多个 Or 分组的功能
				newResult := result.Session(&gorm.Session{NewDB: true})
				for _, field := range queryCardDesc.QField {
					newResult = newResult.Or(field+" LIKE ?", "%"+queryCardDesc.Keyword+"%")
				}
				return newResult
			}(queryCardDesc, result)

			// 通过 Group Conditions(组条件) 功能将查询分组
			result = result.Where(f)
		} else {
			result = result.Where("sc_name LIKE ? OR serial LIKE ? OR effect LIKE ? OR evo_cover_effect LIKE ? OR security_effect LIKE ?",
				"%"+queryCardDesc.Keyword+"%",
				"%"+queryCardDesc.Keyword+"%",
				"%"+queryCardDesc.Keyword+"%",
				"%"+queryCardDesc.Keyword+"%",
				"%"+queryCardDesc.Keyword+"%",
			)
		}
	}

	// 颜色多匹配查询
	if len(queryCardDesc.Color) > 0 {
		f := func(queryCardDesc *models.CardDescQuery, result *gorm.DB) *gorm.DB {
			// 通过 Session() 创建一个新的 DB 实例，避免影响原来的 DB 实例。用以实现为多个 Or 分组的功能
			newResult := result.Session(&gorm.Session{NewDB: true})
			for _, color := range queryCardDesc.Color {
				newResult = newResult.Or("color LIKE ?", "%"+color+"%")
			}
			return newResult
		}(queryCardDesc, result)

		// 通过 Group Conditions(组条件) 功能将查询分组
		result = result.Debug().Where(f)
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
	result = result.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&cd).Offset(-1).Limit(-1).Count(&CardCount)
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

// 获取所有卡牌可用的等级。即.获取 level 列的值并去重
func GetCardDescLevel() ([]string, error) {
	var (
		level []string
	)

	result := DB.Model(&models.CardDesc{}).Distinct().Where("level != ''").Pluck("level", &level)
	if result.Error != nil {
		return nil, result.Error
	}

	return level, nil
}
