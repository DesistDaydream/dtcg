package database

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type CardPriceA struct {
	PackPrefix    string `gorm:"column:pack_prefix"`
	CardID        int
	CardVersionID string  `gorm:"column:card_version_id"`
	Serial        string  `gorm:"column:serial"`
	ScName        string  `gorm:"column:sc_name"`
	Rarity        string  `gorm:"column:rarity"`
	MinPrice      float64 `gorm:"column:min_price"`
	AvgPrice      float64 `gorm:"column:avg_price"`
}

func TestDatabase(t *testing.T) {
	i := &DBInfo{
		FilePath: "my_dtcg.db",
	}
	InitDB(i)

	var cardsPrice []CardPrice
	sql := `
SELECT
	c_set.pack_prefix,
	card.card_id,
	card_version_id,
	serial,sc_name,rarity,min_price,avg_price
FROM
	card_desc_from_dtcg_dbs card
	LEFT JOIN card_prices price ON price.card_id=card.card_id
	LEFT JOIN card_group_from_dtcg_dbs c_set ON c_set.pack_id=card.card_pack`
	result := db.Raw(sql).Scan(&cardsPrice)
	if result.Error != nil {
		logrus.Fatalf("从数据库获取卡片信息失败: %v", result.Error)
	}

	for _, cardPrice := range cardsPrice {
		fmt.Println(cardPrice)
	}
}
