package cards

import (
	"encoding/json"
	"fmt"

	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type GetFlags struct {
}

var getFlags GetFlags

func GetCmd() *cobra.Command {
	long := `获取卡牌详情
`
	getCardsCmd := &cobra.Command{
		Use:   "get",
		Short: "获取卡牌详情",
		Long:  long,
		Run:   getCards,
	}

	return getCardsCmd
}

// 列出商品
func getCards(cmd *cobra.Command, args []string) {
	resp, err := handler.H.JhsServices.Market.GetCardVersions(cardsFlags.CardVersionID)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	a, _ := json.Marshal(resp)
	fmt.Printf("%s", a)
}
