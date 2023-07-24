package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// 更新没有卡图的卡牌价格
func UpdateByNoImgCommand() *cobra.Command {
	updateNoImgCardPriceCmd := &cobra.Command{
		Use:   "no-img",
		Short: "更新没有卡图的卡牌价格",
		Run:   updateNoImageCard,
	}

	return updateNoImgCardPriceCmd
}

func updateNoImageCard(cmd *cobra.Command, args []string) {
	cardsPrice, err := database.ListCardsPrice()
	if err != nil {
		logrus.Fatalf("获取卡片价格信息失败: %v", err)
	}

	for _, cardPrice := range cardsPrice.Data {
		if cardPrice.ImageUrl == "" {
			logrus.Infof("开始处理 %v", cardPrice.ScName)
			imageUrl := GetImageURL(cardPrice.CardVersionID)
			database.UpdateCardPrice(&models.CardPrice{CardIDFromDB: cardPrice.CardIDFromDB}, map[string]interface{}{
				"set_id":          cardPrice.SetID,
				"set_prefix":      cardPrice.SetPrefix,
				"serial":          cardPrice.Serial,
				"sc_name":         cardPrice.ScName,
				"alternative_art": cardPrice.AlternativeArt,
				"rarity":          cardPrice.Rarity,
				"card_version_id": cardPrice.CardVersionID,
				"min_price":       cardPrice.MinPrice,
				"avg_price":       cardPrice.AvgPrice,
				"image_url":       imageUrl,
			})
		}
	}
}
