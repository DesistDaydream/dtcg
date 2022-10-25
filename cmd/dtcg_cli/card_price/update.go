package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	SetPrefix      []string
	StartAt        int
	UpdateImageURL bool
}

var updateFlags UpdateFlags

func UpdateCardPriceCommand() *cobra.Command {
	UpdateCardPriceCmd := &cobra.Command{
		Use:   "update",
		Short: "更新卡牌价格",
		Run:   updateCardPrice,
	}

	UpdateCardPriceCmd.Flags().StringSliceVar(&updateFlags.SetPrefix, "sets-name", []string{}, "更新哪些卡包的价格，使用 card-set list 子命令获取卡包名称。若不指定则更新所有")
	UpdateCardPriceCmd.Flags().IntVar(&updateFlags.StartAt, "start-at", 0, "从哪张卡牌开始更新")
	UpdateCardPriceCmd.Flags().BoolVar(&updateFlags.UpdateImageURL, "update-image", false, "是否更新卡牌的图片URL")

	return UpdateCardPriceCmd
}

func updateCardPrice(cmd *cobra.Command, args []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("获取全部卡牌描述失败: %v", err)
	}

	if len(updateFlags.SetPrefix) != 0 {
		for _, cardDesc := range cardsDesc.Data {
			updateCardPriceBaseonCardSet(cardDesc, updateFlags.SetPrefix)
		}
	} else {
		updateCardPriceBaseonStartAt(cardsDesc, updateFlags.StartAt)
	}
}

func updateCardPriceBaseonStartAt(cardsDesc *models.CardsDesc, startAtCardIDFromDB int) {
	var startAt int
	if startAtCardIDFromDB != 0 {
		for i, cardDesc := range cardsDesc.Data {
			if cardDesc.CardIDFromDB == startAtCardIDFromDB {
				startAt = i
			}
		}
	} else {
		startAt = 0
	}

	for i := startAt; i < len(cardsDesc.Data); i++ {
		updateRun(&cardsDesc.Data[i])
	}
}

func updateCardPriceBaseonCardSet(cardDesc models.CardDesc, setsPrefix []string) {
	for _, setPrefix := range setsPrefix {
		if cardDesc.SetPrefix == setPrefix {
			updateRun(&cardDesc)
		}
	}
}

func updateRun(cardDesc *models.CardDesc) {
	cardVersionID, minPrice, avgPrice := GetPrice(cardDesc)

	if updateFlags.UpdateImageURL {
		imageUrl := GetImageURL(cardVersionID)
		database.UpdateCardPrice(&models.CardPrice{
			CardIDFromDB:  cardDesc.CardIDFromDB,
			CardVersionID: cardVersionID,
			MinPrice:      minPrice,
			AvgPrice:      avgPrice,
			ImageUrl:      imageUrl,
		}, map[string]string{})
	} else {
		database.UpdateCardPrice(&models.CardPrice{
			CardIDFromDB: cardDesc.CardIDFromDB,
			MinPrice:     minPrice,
			AvgPrice:     avgPrice,
		}, map[string]string{})
	}
}
