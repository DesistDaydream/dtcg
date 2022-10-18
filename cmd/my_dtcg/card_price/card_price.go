package cardprice

import (
	"fmt"
	"strconv"

	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cardSetCmd := &cobra.Command{
		Use:   "card-price",
		Short: "控制卡片价格信息",
		// PersistentPreRun: cardSetPersistentPreRun,
	}

	cardSetCmd.AddCommand(
		AddCardPriceCommand(),
		UpdateCardPriceCommand(),
	)

	return cardSetCmd
}

var client *services.SearchClient

// func cardSetPersistentPreRun(cmd *cobra.Command, args []string) {
// 	client = services.NewSearchClient(core.NewClient(""))
// }

func GetPrice(cardDesc *models.CardDesc) (int, float64, float64) {
	cardPrice, err := client.GetCardPrice(fmt.Sprint(cardDesc.CardIDFromDB))
	if err != nil {
		logrus.Fatalf("获取卡片价格失败: %v", err)
	}

	var minPrice float64
	if len(cardPrice.Data.Products) == 0 {
		minPrice = 0
	} else {
		minPrice, _ = strconv.ParseFloat(cardPrice.Data.Products[0].MinPrice, 64)
	}

	avgPrice, _ := strconv.ParseFloat(cardPrice.Data.AvgPrice, 64)

	return int(cardPrice.Data.Products[0].CardVersionID), minPrice, avgPrice
}
