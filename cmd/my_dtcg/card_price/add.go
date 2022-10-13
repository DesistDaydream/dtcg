package cardprice

import (
	"fmt"
	"strconv"

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

	for i := startAt; i < len(cardsDesc.Data); i++ {
		client := services.NewSearchClient(core.NewClient(""))

		cardPrice, err := client.GetCardPrice(fmt.Sprint(cardsDesc.Data[i].CardIDFromDB))
		if err != nil {
			logrus.Errorf("获取卡片价格失败: %v", err)
		}

		var f1 float64
		if len(cardPrice.Data.Products) == 0 {
			f1 = 0
		} else {
			f1, _ = strconv.ParseFloat(cardPrice.Data.Products[0].MinPrice, 64)
		}

		f2, _ := strconv.ParseFloat(cardPrice.Data.AvgPrice, 64)

		database.AddCardPirce(&models.CardPrice{
			CardIDFromDB:   cardsDesc.Data[i].CardIDFromDB,
			SetID:          cardsDesc.Data[i].SetID,
			SetPrefix:      cardsDesc.Data[i].SetPrefix,
			Serial:         cardsDesc.Data[i].Serial,
			ScName:         cardsDesc.Data[i].ScName,
			AlternativeArt: cardsDesc.Data[i].AlternativeArt,
			Rarity:         cardsDesc.Data[i].Rarity,
			CardVersionID:  int(cardPrice.Data.Products[0].CardVersionID),
			MinPrice:       f1,
			AvgPrice:       f2,
		})
	}
}
