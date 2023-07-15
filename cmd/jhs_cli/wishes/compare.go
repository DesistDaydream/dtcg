// 匹配官方推荐的清单的最低价，与自己的库存进行比较
package wishes

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CompareFlags struct {
	WishListID string
}

var compareFlags CompareFlags

func CompareCommand() *cobra.Command {
	long := `
匹配官方推荐的清单的最低价，与自己的库存进行比较
`
	compareCmd := &cobra.Command{
		Use:   "compare",
		Short: "匹配官方推荐的清单的最低价，与自己的库存进行比较",
		Long:  long,
		Run:   compare,
	}

	compareCmd.Flags().StringVar(&compareFlags.WishListID, "wlid", "2610301", "清单 ID")

	return compareCmd
}

func compare(cmd *cobra.Command, args []string) {
	// 注意，这里使用 go 标准库中的 text/tbwriter
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"卡名", "编号", "我的价格", "ta的价格", "我的数量", "ta的数量"})

	resp, err := handler.H.JhsServices.Wishes.WishListMatch(compareFlags.WishListID)
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, card := range resp[0].MatchCards {
		products, err := handler.H.JhsServices.Sellers.ProductList("1", card.Number, "1", "price_asc")
		if err != nil {
			logrus.Fatal(err)
		}

		if len(products.Data) > 0 {
			table.Append([]string{card.CardName, card.Number, products.Data[0].Price, card.Price, strconv.Itoa(products.Data[0].Quantity), strconv.FormatInt(card.Quantity, 10)})
		} else {
			table.Append([]string{card.CardName, card.Number, "0", card.Price, "0", strconv.FormatInt(card.Quantity, 10)})
		}
	}
	table.Render()
}
