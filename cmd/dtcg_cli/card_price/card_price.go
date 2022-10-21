package cardprice

import (
	"fmt"
	"strconv"

	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cardSetCmd := &cobra.Command{
		Use:              "card-price",
		Short:            "控制卡片价格信息",
		PersistentPreRun: cardSetPersistentPreRun,
	}

	cardSetCmd.AddCommand(
		AddCardPriceCommand(),
		UpdateCardPriceCommand(),
	)

	return cardSetCmd
}

var client *cdb.CdbClient

func cardSetPersistentPreRun(cmd *cobra.Command, args []string) {
	// 执行根命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}
	client = cdb.NewCdbClient(core.NewClient(""))
}

func GetPrice(cardDesc *models.CardDesc) (int, float64, float64) {
	var (
		minPrice      float64
		avgPrice      float64
		cardVersionID int
	)

	cardPrice, err := client.GetCardPrice(fmt.Sprint(cardDesc.CardIDFromDB))
	if err != nil {
		logrus.Fatalf("获取卡片价格失败: %v", err)
	}

	avgPrice, _ = strconv.ParseFloat(cardPrice.Data.AvgPrice, 64)

	if cardPrice.Data.Total == 0 {
		minPrice = 0
		cardVersionID = 0
	} else {
		minPrice, _ = strconv.ParseFloat(cardPrice.Data.Products[0].MinPrice, 64)
		cardVersionID = int(cardPrice.Data.Products[0].CardVersionID)
	}

	return cardVersionID, minPrice, avgPrice
}
