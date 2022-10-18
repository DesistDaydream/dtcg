package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func AddCardPriceCommand() *cobra.Command {
	AddCardPriceCmd := &cobra.Command{
		Use:   "add",
		Short: "添加卡片集合",
		Run:   addCardPrice,
	}

	AddCardPriceCmd.Flags().Int("startAt", 0, "从哪个卡牌开始添加，使用从 dtcg db 中获取到的卡片 ID。")

	return AddCardPriceCmd
}

func addCardPrice(cmd *cobra.Command, args []string) {
	startAtCardIDFromDB, _ := cmd.Flags().GetInt("startAt")
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	// startAt 是我自己的编号。
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

	client = services.NewSearchClient(core.NewClient(""))

	for i := startAt; i < len(cardsDesc.Data); i++ {
		cardVersionID, minPrice, avgPrice := GetPrice(&cardsDesc.Data[i])

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
		})
	}
}
