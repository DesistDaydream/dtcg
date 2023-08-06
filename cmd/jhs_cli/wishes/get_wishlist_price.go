package wishes

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func GetWishListPriceCmd() *cobra.Command {
	long := `
根据集换社的我想收清单，获取卡组的集换价、最低价
`
	getWishListPriceCmd := &cobra.Command{
		Use:   "get-price",
		Short: "根据集换社的我想收清单，获取卡组的集换价、最低价",
		Long:  long,
		Run:   runGetWishListPrice,
	}

	return getWishListPriceCmd

}

func runGetWishListPrice(cmd *cobra.Command, args []string) {
	var allAvgPrice, allMinPrice float64
	wishListGetResp, err := handler.H.JhsServices.Wishes.Get(args[0])
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"名称", "数量", "编号", "稀有度", "异画", "集换单价", "集换价", "最低单价", "最低价"})

	for _, wish := range wishListGetResp.Data {
		cardPrice, err := database.GetCardPriceWhereCardVersionID(strconv.Itoa(wish.CardVersionID))
		if err != nil {
			logrus.Fatalf("%v", err)
		}

		minPrice := cardPrice.MinPrice * float64(wish.Quantity)
		avgPrice := cardPrice.AvgPrice * float64(wish.Quantity)

		table.Append([]string{
			cardPrice.ScName,
			strconv.Itoa(wish.Quantity),
			wish.Number,
			wish.Rarity,
			cardPrice.AlternativeArt,
			fmt.Sprintf("%.2f", cardPrice.AvgPrice),
			fmt.Sprintf("%.2f", avgPrice),
			fmt.Sprintf("%.2f", cardPrice.MinPrice),
			fmt.Sprintf("%.2f", minPrice),
		})

		allAvgPrice = allAvgPrice + avgPrice
		allMinPrice = allMinPrice + minPrice
	}

	table.Render()

	fmt.Println(fmt.Sprintf("%.2f", allAvgPrice), fmt.Sprintf("%.2f", allMinPrice))
}
