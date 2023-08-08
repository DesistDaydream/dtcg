package cards

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ListFlags struct {
	CategoryID string
	Page       int
}

var (
	listFlags ListFlags
)

func ListCmd() *cobra.Command {
	long := `
`
	listCardsCmd := &cobra.Command{
		Use:   "list",
		Short: "列出卡牌",
		Long:  long,
		Run:   listCards,
	}

	listCardsCmd.PersistentFlags().StringVar(&listFlags.CategoryID, "cid", "4793", "类型")
	listCardsCmd.PersistentFlags().IntVar(&listFlags.Page, "page", 1, "页")

	return listCardsCmd
}

// 列出商品
func listCards(cmd *cobra.Command, args []string) {
	resp, err := handler.H.JhsServices.Market.ListCardVersions(listFlags.CategoryID, listFlags.Page)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Println(resp)
}
