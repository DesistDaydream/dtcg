package database

import (
	"strings"

	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 添加卡牌价格
func AddCardPirce(cardPrice *models.CardPrice) {
	result := DB.FirstOrCreate(cardPrice, models.CardPrice{CardIDFromDB: cardPrice.CardIDFromDB})
	if result.Error != nil {
		logrus.Errorf("插入数据失败: %v", result.Error)
	}
}

// 更新卡牌价格
func UpdateCardPrice(cardPrice *models.CardPrice, condition map[string]interface{}) {
	// 注意：当使用 struct 进行更新时，GORM 只会更新非零值的字段。
	// result := DB.Model(cardPrice).Select("min_price,avg_price").Where("card_id_from_db = ?", cardPrice.CardIDFromDB).Updates(models.CardPrice{
	// 	SetID:          cardPrice.SetID,
	// 	SetPrefix:      cardPrice.SetPrefix,
	// 	Serial:         cardPrice.Serial,
	// 	ScName:         cardPrice.ScName,
	// 	AlternativeArt: cardPrice.AlternativeArt,
	// 	Rarity:         cardPrice.Rarity,
	// 	CardVersionID:  cardPrice.CardVersionID,
	// 	MinPrice:       cardPrice.MinPrice,
	// 	AvgPrice:       cardPrice.AvgPrice,
	// 	ImageUrl:       cardPrice.ImageUrl,
	// })
	// 所以我们使用 map 更新字段，这样就可以更新零值字段了，具体更新哪个字段，由 condition 决定，即由函数的调用者决定
	result := DB.Model(cardPrice).Where("card_id_from_db = ?", cardPrice.CardIDFromDB).Updates(condition)

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
		Count:     result.RowsAffected,
		PageSize:  -1,
		PageNum:   1,
		PageTotal: 1,
		Data:      cp,
	}, nil
}

// 列出所有卡牌价格详情，分页
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
		Count:     CardCount,
		PageSize:  pageSize,
		PageNum:   pageNum,
		PageTotal: (int(CardCount) / pageSize) + 1,
		Data:      cp,
	}, nil
}

// 列出所有卡牌价格详情，使用分页库的逻辑实现分页
func GetCardsPriceWithPaginationLib(c *gin.Context) (*models.CardsPrice, error) {
	// 实例化，并传递类型参数为 models.CardPrice 以约束 Pagination.Data 的类型
	pagination := NewPagination[models.CardPrice](models.CardPrice{}, c)
	err := pagination.Query()
	if err != nil {
		return nil, err
	}

	return &models.CardsPrice{
		Count:     pagination.Count,
		PageSize:  pagination.PageSize,
		PageNum:   pagination.PageNum,
		PageTotal: pagination.PageTotal,
		Data:      pagination.Data, // 由于 New 时约束了类型，所以不用断言可以直接赋值。
	}, nil
}

// 根据卡片集合前缀列出卡牌价格详情
func GetCardPriceWhereSetPrefix(setPrefix string) (*models.CardsPrice, error) {
	var cp []models.CardPrice
	result := DB.Where("set_prefix = ?", setPrefix).Find(&cp)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.CardsPrice{
		Count:     result.RowsAffected,
		PageSize:  -1,
		PageNum:   1,
		PageTotal: 1,
		Data:      cp,
	}, nil
}

// 根据条件列出卡牌价格详情，分页
func GetCardPriceByCondition(pageSize int, pageNum int, cardPriceQuery *models.CardPriceQuery) (*models.CardsPrice, error) {
	var (
		CardCount int64
		cp        []models.CardPrice
	)

	result := DB.Model(&models.CardPrice{})

	// 过滤指定 card_version_id 的卡牌
	if cardPriceQuery.CardVersionID != 0 {
		result = result.Where("card_version_id = ?", cardPriceQuery.CardVersionID)
	}

	// 过滤不包含某些 card_version_id 的卡牌
	if len(cardPriceQuery.NotInCardVersionID) > 0 {
		result = result.Where("card_version_id NOT IN ?", cardPriceQuery.NotInCardVersionID)
	}

	// 过滤指定卡牌集合的卡牌
	if len(cardPriceQuery.SetsPrefix) > 0 {
		result = result.Where("set_prefix IN ?", cardPriceQuery.SetsPrefix)
	}

	// 过滤指定稀有度的卡牌
	if len(cardPriceQuery.Rarity) > 0 {
		f := func(cardPriceQuery *models.CardPriceQuery, result *gorm.DB) *gorm.DB {
			// 通过 Session() 创建一个新的 DB 实例，避免影响原来的 DB 实例。用以实现为多个 Or 分组的功能
			newResult := result.Session(&gorm.Session{NewDB: true})
			for _, rarity := range cardPriceQuery.Rarity {
				newResult = newResult.Or("rarity LIKE ?", rarity+"%")
			}
			return newResult
		}(cardPriceQuery, result)

		// 通过 Group Conditions(组条件) 功能将查询分组
		result = result.Debug().Where(f)
	}

	// 根据关键字过滤卡牌
	if cardPriceQuery.Keyword != "" {
		result = result.Where("sc_name LIKE ? OR serial LIKE ?",
			"%"+cardPriceQuery.Keyword+"%",
			"%"+cardPriceQuery.Keyword+"%",
		)
	}

	// 是否是异画
	if cardPriceQuery.AlternativeArt != "" {
		result = result.Debug().Where("alternative_art = ?", cardPriceQuery.AlternativeArt)
	}

	// 根据集换价范围查询
	if cardPriceQuery.AvgPriceRange != "" {
		priceRange := strings.Split(cardPriceQuery.AvgPriceRange, "-")
		if len(priceRange) == 2 {
			result = result.Where("avg_price BETWEEN ? AND ?", priceRange[0], priceRange[1])
		}
	}

	// 根据最低价范围查询
	if cardPriceQuery.MinPriceRange != "" {
		priceRange := strings.Split(cardPriceQuery.MinPriceRange, "-")
		if len(priceRange) == 2 {
			result = result.Where("min_price BETWEEN ? AND ?", priceRange[0], priceRange[1])
		}
	}

	// 分页、计数
	result = result.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&cp).Offset(-1).Limit(-1).Count(&CardCount)
	if condition := result.Error; condition != nil {
		return nil, condition
	}

	return &models.CardsPrice{
		Count:     CardCount,
		PageSize:  pageSize,
		PageNum:   pageNum,
		PageTotal: (int(CardCount) / pageSize) + 1,
		Data:      cp,
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

// 根据条件从 card_price 表获取卡牌价格中带有 card_desc 表中的图片
func GetCardPriceWithDtcgDBImgByCondition(pageSize int, pageNum int, cardPriceQuery *models.CardPriceQuery) (*models.CardsPriceWithImageDB, error) {
	var (
		CardCount int64
		cp        []models.CardPriceWithImageDB
	)

	result := DB.Model(&models.CardPrice{})

	// 联表查询
	sqlSelect := `card_prices.card_id_from_db AS card_id_from_db,
card_prices.set_prefix AS set_prefix,
card_prices.serial AS serial,
card_prices.sc_name AS sc_name,
card_prices.alternative_art AS alternative_art,
card_prices.rarity AS rarity,
card_prices.card_version_id AS card_version_id,
card_prices.min_price AS min_price,
card_prices.avg_price AS avg_price,
card_descs.image AS image`
	result = result.Select(sqlSelect).Joins("LEFT JOIN card_descs ON card_prices.card_id_from_db = card_descs.card_id_from_db").Debug()

	// 根据关键字从多列模糊查询
	if cardPriceQuery.Keyword != "" {
		result = result.Where("card_prices.sc_name LIKE ? OR card_prices.serial LIKE ?",
			"%"+cardPriceQuery.Keyword+"%",
			"%"+cardPriceQuery.Keyword+"%",
		)
	}

	// 是否是异画
	if cardPriceQuery.AlternativeArt != "" {
		result = result.Where("card_prices.alternative_art = ?", cardPriceQuery.AlternativeArt)
	}

	// 分页、计数
	result = result.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&cp).Offset(-1).Limit(-1).Count(&CardCount)
	if condition := result.Error; condition != nil {
		return nil, condition
	}

	return &models.CardsPriceWithImageDB{
		Count:       CardCount,
		PageSize:    pageSize,
		PageCurrent: pageNum,
		PageTotal:   (int(CardCount) / pageSize) + 1,
		Data:        cp,
	}, nil
}
