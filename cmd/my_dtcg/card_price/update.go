package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func UpdateCardPriceCommand() *cobra.Command {
	UpdateCardPriceCmd := &cobra.Command{
		Use:   "update",
		Short: "更新卡片集合",
		Run:   updateCardPrice,
	}

	return UpdateCardPriceCmd
}

func updateCardPrice(cmd *cobra.Command, args []string) {
	UpdateCardPrice([]string{})
}

func UpdateCardPrice(setsPrefix []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("获取全部卡片描述失败: %v", err)
	}

	for _, cardDesc := range cardsDesc.Data {
		if setsPrefix != nil {
			updateCardPriceBaseonCardSet(cardDesc, setsPrefix)
		}
	}
}

func updateCardPriceBaseonCardSet(cardDesc models.CardDesc, setsPrefix []string) {
	for _, setPrefix := range setsPrefix {
		if cardDesc.SetPrefix == setPrefix {
			database.UpdateCardPrice()
		}
	}
}
