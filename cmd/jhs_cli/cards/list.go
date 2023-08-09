package cards

import (
	"os"
	"strconv"

	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ListFlags struct {
	PackID     int
	CategoryID int
	Page       int
}

var listFlags ListFlags

func ListCmd() *cobra.Command {
	long := `
列出卡牌`

	listCardsCmd := &cobra.Command{
		Use:   "list",
		Short: "列出卡牌",
		Long:  long,
		Run:   listCards,
	}

	listCardsCmd.PersistentFlags().IntVar(&listFlags.PackID, "pid", 115, "卡包 ID")
	listCardsCmd.PersistentFlags().IntVar(&listFlags.CategoryID, "cid", 0, "集合类型")

	return listCardsCmd
}

// 列出商品
func listCards(cmd *cobra.Command, args []string) {
	table := tablewriter.NewWriter(os.Stdout)
	// table.SetHeader([]string{"名称", "数量", "编号", "稀有度", "异画", "集换单价", "集换价", "最低单价", "最低价"})

	page := 1
	for {
		resp, err := handler.H.JhsServices.Market.ListCardVersions(listFlags.PackID, listFlags.CategoryID, page)
		if err != nil {
			logrus.Errorf("%v", err)
		}
		for _, card := range resp.Data {
			table.Append([]string{strconv.Itoa(card.CardVersionID), strconv.Itoa(card.CardID), card.NameCN, card.Number})
		}

		if resp.NextPageURL == "" {
			break
		}

		page++
	}

	table.Render()

}
