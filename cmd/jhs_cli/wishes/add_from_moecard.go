package wishes

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func addFromMoecardCmd() *cobra.Command {
	long := `
从 Moecard 的 云卡组 获取卡牌列表，添加到我想收清单中
`
	addFromWishListCmd := &cobra.Command{
		Use:   "moecard",
		Short: "从 Moecard 的 云卡组 获取卡牌列表，添加到我想收清单中",
		Long:  long,
		Run:   addFromMoecard,
	}

	return addFromWishListCmd
}

// 添加商品
func addFromMoecard(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		logrus.Fatalln("请指定 CDID")
	}

	deck, err := handler.H.MoecardServices.Community.GetDeckCloud(args[0])
	if err != nil {
		logrus.Fatalln(err)
	}

	wishList, err := handler.H.JhsServices.Wishes.CreateWashList(args[0])
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
