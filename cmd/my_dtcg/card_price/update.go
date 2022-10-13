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

type UpdateFlags struct {
	IDs []int
}

var updateFlags UpdateFlags

func UpdateCardPriceCommand() *cobra.Command {
	UpdateCardPriceCmd := &cobra.Command{
		Use:   "update",
		Short: "更新卡片集合",
		Run:   updateCardPrice,
	}

	// UpdateCardPriceCmd.Flags().IntSlice("id", []int{}, "从哪个卡牌开始添加，使用从 dtcg db 中获取到的卡片 ID。")
	UpdateCardPriceCmd.Flags().IntSliceVar(&updateFlags.IDs, "ids", []int{}, "从哪个卡牌开始添加，使用从 dtcg db 中获取到的卡片 ID。")

	return UpdateCardPriceCmd
}

func updateCardPrice(cmd *cobra.Command, args []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("获取全部卡片描述失败: %v", err)
	}

	for _, cardDesc := range cardsDesc.Data {
		if len(updateFlags.IDs) != 0 {
			updateCardPriceBaseonCardSet(cardDesc, updateFlags.IDs)
		}
	}
}

func updateCardPriceBaseonCardSet(cardDesc models.CardDesc, setsPrefix []int) {
	client := services.NewSearchClient(core.NewClient(""))

	for _, setPrefix := range setsPrefix {
		if cardDesc.CardIDFromDB == setPrefix {
			cardPrice, err := client.GetCardPrice(fmt.Sprint(setPrefix))
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
			database.UpdateCardPrice(&models.CardPrice{
				CardIDFromDB: setPrefix,
				MinPrice:     f1,
				AvgPrice:     f2,
			}, map[string]string{})
		}
	}
}
