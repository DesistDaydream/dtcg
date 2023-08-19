package products

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateQuantityFlags struct {
	ProductQuantity string
}

type UpdateQuantityPolicy struct {
}

var (
	updateQuantityFlags UpdateQuantityFlags
)

func UpdateQuantityCommand() *cobra.Command {
	long := `
根据策略更新商品数量。
比如：
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 0,0.99 表示将所有价格在 0-0.99 之间卡牌的价格不增加，以集换价售卖。
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 1,9.99 -c 0.5 表示将所有价格在 1-9.99 之间卡牌的价格增加 0.5 元。
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 3,9.99 --art="否" -c 2  表示将所有价格在 3-9.99 之间的非异画卡牌的价格增加 0.5 元。
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 0,10000 --art="是" -c 50 表示将所有价格在 0-10000 之间的异画卡价格增加 50 元。
	go run cmd/jhs_cli/main.go products update price -s BTC-05 --rarity C,U,R --art 否 -o "*" -c 1.03 表示将所有 BTC-05 的 U、R、C 稀有度的原画卡的价格乘以 1.03
`
	UpdateProductsPriceCmd := &cobra.Command{
		Use:   "quantity",
		Short: "更新商品数量",
		Long:  long,
		Run:   updateQuantity,
	}

	UpdateProductsPriceCmd.Flags().StringVarP(&updateQuantityFlags.ProductQuantity, "quantity", "q", "", "商品数量")

	return UpdateProductsPriceCmd
}

func updateQuantity(cmd *cobra.Command, args []string) {
	// 生成待处理的卡牌信息
	cards, err := GenNeedHandleCards()
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}

	// 根据更新策略更新卡牌价格
	ps := genNeedUpdateProducts(cards, 0)
	logrus.Infof("共匹配到 %v 件商品", ps.count)
	for _, p := range ps.products {
		logrus.WithFields(logrus.Fields{
			"原始价格": p.card.AvgPrice,
			"更新价格": p.product.Price,
			"调整价格": fmt.Sprintf("%v %v", updatePriceFlags.UpdatePolicy.Operator, 0),
		}).Debugf("检查生成的商品: 【%v】【%v】【%v %v】", p.product.CardVersionID, p.card.AlternativeArt, p.card.Serial, p.product.CardNameCN)

		if productsFlags.isRealRun {
			updateRun(
				&p.product,
				fmt.Sprint(p.product.OnSale),
				p.product.Price,
				p.product.CardVersionImage,
				updateQuantityFlags.ProductQuantity,
			)
		}
	}

	// 注意：总数不等于任何数量之和。
	logrus.WithFields(logrus.Fields{
		"总数": cards.Count,
		"成功": updateSuccessCount,
		"失败": updateFailCount,
		"跳过": updateSkip,
	}).Infof("更新结果")
}
