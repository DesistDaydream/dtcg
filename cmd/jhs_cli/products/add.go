package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func AddCommand() *cobra.Command {
	AddCardSetCmd := &cobra.Command{
		Use:   "add",
		Short: "添加商品",
		Run:   addProducts,
	}

	return AddCardSetCmd
}

// 添加商品
func addProducts(cmd *cobra.Command, args []string) {
	cards := []string{"2954", "2955", "2956", "2957"}

	for _, cardVersionID := range cards {
		var price string

		cardPrice, err := database.GetCardPriceWhereCardVersionID(cardVersionID)
		if err != nil {
			logrus.Errorln("获取卡牌价格详情失败", err)
		}

		if cardPrice.AvgPrice == 0 {
			price = "10"
		} else {
			price = fmt.Sprint(cardPrice.AvgPrice + float64(2))
		}

		// 开始上架
		resp, err := handler.H.JhsServices.Products.Add(&models.ProductsAddReqBody{
			CardVersionID:        cardVersionID,
			Price:                price,
			Quantity:             "4",
			Condition:            "1",
			Remark:               "",
			GameKey:              "dgm",
			UserCardVersionImage: cardPrice.ImageUrl,
		})

		// resp, err := client.Add(cardModelToCardVersionID[rows[i][0]], rows[i][1], rows[i][2])
		if err != nil {
			logrus.Errorf("%v 上架失败：%v", cardPrice.ScName, err)
		} else {
			logrus.Infof("%v 上架成功：%v", cardPrice.ScName, resp)
		}
	}
}
