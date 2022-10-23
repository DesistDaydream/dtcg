package cardprice

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cardPriceCmd := &cobra.Command{
		Use:              "card-price",
		Short:            "控制卡片价格信息",
		PersistentPreRun: cardPricePersistentPreRun,
	}

	cardPriceCmd.AddCommand(
		AddCardPriceCommand(),
		UpdateCardPriceCommand(),
	)

	return cardPriceCmd
}

var cdbClient *cdb.CdbClient

func cardPricePersistentPreRun(cmd *cobra.Command, args []string) {
	// 执行根命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}

	// 自身执行前操作
	cdbClient = cdb.NewCdbClient(core.NewClient(""))
}

func GetPrice(cardDesc *models.CardDesc) (int, float64, float64) {
	var (
		minPrice      float64
		avgPrice      float64
		cardVersionID int
	)

	cardPrice, err := cdbClient.GetCardPrice(fmt.Sprint(cardDesc.CardIDFromDB))
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

func GetImageURL(jhsClient *market.MarketClient, cardVersionID int) string {
	var imageUrl string
	productSellers, err := jhsClient.GetProductSellers(fmt.Sprint(cardVersionID))
	if err != nil {
		logrus.Errorf("%v", err)
	}

	for _, d := range productSellers.Data {
		if strings.Contains(d.CardVersionImage, "cdn-client") {
			imageUrl = d.CardVersionImage
			break
		}
	}

	// i, _ := url.QueryUnescape(imageUrl)

	return imageUrl
}
