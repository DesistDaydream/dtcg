package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	SetPrefix []string
	StartAt   int
}

var updateFlags UpdateFlags

func UpdateCardPriceCommand() *cobra.Command {
	UpdateCardPriceCmd := &cobra.Command{
		Use:   "update",
		Short: "更新卡片集合",
		Run:   updateCardPrice,
	}

	UpdateCardPriceCmd.Flags().StringSliceVar(&updateFlags.SetPrefix, "sets-name", []string{}, "更新哪些卡包的价格，使用 card-set list 子命令获取卡包名称。若不指定则更新所有")
	UpdateCardPriceCmd.Flags().IntVar(&updateFlags.StartAt, "start-at", 0, "从哪个卡片开始更新")

	return UpdateCardPriceCmd
}

func updateCardPrice(cmd *cobra.Command, args []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("获取全部卡片描述失败: %v", err)
	}

	// client = services.NewSearchClient(core.NewClient(""))

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
	_, minPrice, avgPrice := GetPrice(cardDesc)

	database.UpdateCardPrice(&models.CardPrice{
		CardIDFromDB: cardDesc.CardIDFromDB,
		MinPrice:     minPrice,
		AvgPrice:     avgPrice,
	}, map[string]string{})
}
