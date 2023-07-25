package wishes

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type AddFlags struct {
	CDID string
}

var addFlags AddFlags

func AddCommand() *cobra.Command {
	long := `
从 dtcg db 的 云卡组 获取卡牌列表，添加到我想收清单中
`
	addProdcutCmd := &cobra.Command{
		Use:   "add",
		Short: "从 dtcg db 的 云卡组 获取卡牌列表，添加到我想收清单中",
		Long:  long,
		Run:   addProducts,
	}

	addProdcutCmd.Flags().StringVarP(&addFlags.CDID, "cdid", "c", "", "DTCG DB 中我的卡组的 ID")

	return addProdcutCmd
}

// 添加商品
func addProducts(cmd *cobra.Command, args []string) {
	if addFlags.CDID == "" {
		logrus.Fatalln("请使用 -c 指定 CDID")
	}
	deck, err := handler.H.MoecardServices.Community.GetDeckCloud(addFlags.CDID)
	if err != nil {
		logrus.Fatalln(err)
	}

	wishList, err := handler.H.JhsServices.Wishes.CreateWashList(addFlags.CDID)
	if err != nil {
		logrus.Fatalf("创建我想收清单失败: %v", err)
	}

	for _, main := range deck.Data.DeckInfo.Main {
		cardPrice, err := database.GetCardPrice(fmt.Sprint(main.Cards.CardID))
		if err != nil {
			logrus.Errorf("获取 %v 价格相关信息失败：%v", cardPrice.ScName, err)
		}
		handler.H.JhsServices.Wishes.Add(fmt.Sprint(cardPrice.CardVersionID), "0", fmt.Sprint(main.Number), "", fmt.Sprint(wishList.WishListID))
		if err != nil {
			logrus.Fatalln(err)
		}
	}

	for _, egg := range deck.Data.DeckInfo.Eggs {
		cardPrice, err := database.GetCardPrice(fmt.Sprint(egg.Cards.CardID))
		if err != nil {
			logrus.Errorf("获取 %v 价格相关信息失败：%v", cardPrice.ScName, err)
		}
		handler.H.JhsServices.Wishes.Add(fmt.Sprint(cardPrice.CardVersionID), "0", fmt.Sprint(egg.Number), "", fmt.Sprint(wishList.WishListID))
		if err != nil {
			logrus.Fatalln(err)
		}
	}
}
