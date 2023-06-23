// 匹配官方推荐的清单的最低价，与自己的库存进行比较
package wishes

import (
	"fmt"
	"os"
	"text/tabwriter"

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
	w := tabwriter.NewWriter(os.Stdout, 4, 8, 4, ' ', 0)
	fmt.Fprintf(w, "Serial\tMy Price\tPrice\tMy Quantity\tQuantity\tCard Name\n")

	resp, err := handler.H.JhsServices.Wishes.WishListMatch(compareFlags.WishListID)
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, card := range resp[0].MatchCards {
		products, err := handler.H.JhsServices.Products.List("1", card.Number, "1", "price_asc")
		if err != nil {
			logrus.Fatal(err)
		}

		if len(products.Data) > 0 {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t\n", card.Number, products.Data[0].Price, card.Price, products.Data[0].Quantity, card.Quantity, card.CardName)
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t\n", card.Number, "0", card.Price, 0, card.Quantity, card.CardName)
		}
	}
	w.Flush()
}
