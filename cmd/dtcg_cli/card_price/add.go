package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type AddFlags struct {
	StartAt int
}

var addFlags AddFlags

func AddCardPriceCommand() *cobra.Command {
	AddCardPriceCmd := &cobra.Command{
		Use:   "add",
		Short: "添加卡牌价格",
		Run:   addCardPrice,
	}

	AddCardPriceCmd.Flags().IntVar(&addFlags.StartAt, "start-at", 0, "从哪个卡牌开始添加，使用从 dtcg db 中获取到的卡片 ID。")

	return AddCardPriceCmd
}

func addCardPrice(cmd *cobra.Command, args []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	// startAt 是我自己的编号。
	var startAt int
	if addFlags.StartAt != 0 {
		for i, cardDesc := range cardsDesc.Data {
			if cardDesc.CardIDFromDB == addFlags.StartAt {
				startAt = i
			}
		}
	} else {
		startAt = 0
	}

	for i := startAt; i < len(cardsDesc.Data); i++ {
		cardVersionID, minPrice, avgPrice := GetPriceFromDtcgdb(&cardsDesc.Data[i])
		// cardVersionID, minPrice, avgPrice := 0, 0.0, 0.0

		// 如果没有卡牌获取到价格，就跳过获取图片 URL
		var imageUrl string
		if cardVersionID != 0 {
			imageUrl = GetImageURL(cardVersionID)
		}

		database.AddCardPirce(&models.CardPrice{
			CardIDFromDB:   cardsDesc.Data[i].CardIDFromDB,
			SetID:          cardsDesc.Data[i].SetID,
			SetPrefix:      cardsDesc.Data[i].SetPrefix,
			Serial:         cardsDesc.Data[i].Serial,
			ScName:         cardsDesc.Data[i].ScName,
			AlternativeArt: cardsDesc.Data[i].AlternativeArt,
			Rarity:         cardsDesc.Data[i].Rarity,
			CardVersionID:  cardVersionID,
			MinPrice:       minPrice,
			AvgPrice:       avgPrice,
			ImageUrl:       imageUrl,
		})
	}
}
