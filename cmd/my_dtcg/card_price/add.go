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
		Run:   AddCardPrice,
	}

	return AddCardPriceCmd
}

func AddCardPrice(cmd *cobra.Command, args []string) {
	addCardPrice("")
}

func addCardPrice(startAtForCardID string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	// 从 startAt 号卡之后开始添加价格信息
	var startAt int
	if startAtForCardID != "" {
		for i, cardDesc := range cardsDesc.Data {
			if fmt.Sprint(cardDesc.CardIDFromDB) == startAtForCardID {
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

		d := &models.CardPrice{
			CardIDFromDB:   cardsDesc.Data[i].CardIDFromDB,
			CardVersionID:  int(cardPrice.Data.Products[0].CardVersionID),
			AlternativeArt: cardsDesc.Data[i].AlternativeArt,
			Rarity:         cardsDesc.Data[i].Rarity,
			SetID:          cardsDesc.Data[i].SetID,
			SetPrefix:      cardsDesc.Data[i].SetPrefix,
			MinPrice:       f1,
			AvgPrice:       f2,
		}

		database.AddCardPirce(d)
	}
}
