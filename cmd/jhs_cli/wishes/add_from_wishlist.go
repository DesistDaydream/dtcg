package wishes

import (
	"strconv"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func addFromWishListCmd() *cobra.Command {
	long := `
从 集换社 的 心愿单 获取卡牌列表，添加到我想收清单中
`
	addFromWishListCmd := &cobra.Command{
		Use:   "wishlist",
		Short: "从 集换社 的 心愿单 获取卡牌列表，添加到我想收清单中",
		Long:  long,
		Run:   addFromWishList,
	}

	return addFromWishListCmd
}

// 添加商品
func addFromWishList(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		logrus.Fatalln("请指定 WishListID")
	}

	wishList, err := handler.H.JhsServices.Wishes.CreateWashList(args[0])
	if err != nil {
		logrus.Fatalf("创建我想收清单失败: %v", err)
	}

	var page int = 1

	for {
		wishListGetResp, err := handler.H.JhsServices.Wishes.Get(args[0], page)
		if err != nil {
			logrus.Fatalf("%v", err)
		}

		for _, product := range wishListGetResp.Data {
			cardPrice, err := database.GetCardPriceWhereCardVersionID(strconv.Itoa(product.CardVersionID))
			if err != nil {
				logrus.Errorf("获取 %v 价格相关信息失败：%v", cardPrice.ScName, err)
			}
			handler.H.JhsServices.Wishes.Add(strconv.Itoa(cardPrice.CardVersionID), "0", strconv.Itoa(product.Quantity), "", strconv.Itoa(wishList.WishListID))
			if err != nil {
				logrus.Fatalln(err)
			}
		}

		if wishListGetResp.NextPageURL == "" {
			logrus.Infof("退出循环时共 %v 页,处理完 %v 页", wishListGetResp.LastPage, wishListGetResp.CurrentPage)
			break
		}

		// 每处理完一页，下一个循环需要处理的页+1
		page++
	}

}
