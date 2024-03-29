package products

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdatePriceFlags struct {
	UpdateInterface string
	UpdatePolicy    UpdatePricePolicy
}

type UpdatePricePolicy struct {
	PriceChange float64
	Operator    string
}

var (
	updatePriceFlags   UpdatePriceFlags
	updateSuccessCount int = 0
	updateFailCount    int = 0
	updateSkip         int = 0
)

func UpdatePriceCommand() *cobra.Command {
	long := `
根据策略更新商品价格。
比如：
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 0,0.99 表示将所有价格在 0-0.99 之间卡牌的价格不增加，以集换价售卖。
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 1,9.99 -c 0.5 表示将所有价格在 1-9.99 之间卡牌的价格增加 0.5 元。
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 3,9.99 --art="否" -c 2  表示将所有价格在 3-9.99 之间的非异画卡牌的价格增加 0.5 元。
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 0,10000 --art="是" -c 50 表示将所有价格在 0-10000 之间的异画卡价格增加 50 元。
	go run cmd/jhs_cli/main.go products update price -s BTC-05 --rarity C,U,R --art 否 -o "*" -c 1.03 表示将所有 BTC-05 的 U、R、C 稀有度的原画卡的价格乘以 1.03
`
	UpdateProductsPriceCmd := &cobra.Command{
		Use:   "price",
		Short: "更新商品价格",
		Long:  long,
		Run:   updatePrice,
	}

	UpdateProductsPriceCmd.Flags().StringVar(&updatePriceFlags.UpdateInterface, "interface", "", "使用集换社的哪个接口获取商品信息。name: 通过卡牌名称从我在卖列出商品信息; id: 通过 card_version_id 直接获取唯一的商品信息")
	UpdateProductsPriceCmd.Flags().StringVarP(&updatePriceFlags.UpdatePolicy.Operator, "operator", "o", "+", "卡牌价格变化的计算方式，乘法还是加法。")
	UpdateProductsPriceCmd.Flags().Float64VarP(&updatePriceFlags.UpdatePolicy.PriceChange, "price-change", "c", 0, "卡牌需要变化的价格。")

	return UpdateProductsPriceCmd
}

func updatePrice(cmd *cobra.Command, args []string) {
	// 生成待处理的卡牌信息
	cards, err := GenNeedHandleCards()
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}

	// 生成待处理的商品信息
	var ps *NeedHandleProducts

	switch updatePriceFlags.UpdateInterface {
	case "name":
		// 生成需要更新的商品
		ps = genNeedUpdateProducts(cards)
		fmt.Printf("\n")
	case "id":
		ps = genNeedUpdateProductsWithBySellerCardVersionId(cards)
		fmt.Printf("\n")
	default:
		logrus.Fatalf("请通过 --interface 指定通过集换社的哪个接口获取商品信息")
	}

	logrus.Infof("共匹配到 %v 件商品", ps.count)

	var needHandleCount int

	// 逐一更新商品价格
	for _, p := range ps.products {
		// 生成商品将要更新的价格
		var newPrice string
		if updatePriceFlags.UpdatePolicy.Operator == "*" {
			newPrice = fmt.Sprintf("%.2f", p.card.AvgPrice*updatePriceFlags.UpdatePolicy.PriceChange)
		} else if updatePriceFlags.UpdatePolicy.Operator == "+" {
			newPrice = fmt.Sprintf("%.2f", p.card.AvgPrice+updatePriceFlags.UpdatePolicy.PriceChange)
		}

		// 只有期望价格与当前价格不一致时，才更新
		if newPrice != p.price {
			// 当期望价格低于当前价格 2 块以上时，不要更新。主要是有的 C、U、R 卡也很值钱，可以高价
			p1, _ := strconv.ParseFloat(p.price, 64)
			p2, _ := strconv.ParseFloat(newPrice, 64)
			if p2-p1 < -2 {
				logrus.WithFields(logrus.Fields{
					"原始价格": p.card.AvgPrice,
					"当前价格": p.price,
					"期望价格": newPrice,
					"调整方式": fmt.Sprintf("%v %v", updatePriceFlags.UpdatePolicy.Operator, updatePriceFlags.UpdatePolicy.PriceChange),
				}).Warnf("商品差价过大，等待后续手动确认: 【%v】【%v】【%v %v】【%v】", p.productID, p.cardVersionID, p.card.Serial, p.cardNameCN, p.card.Rarity)
				continue
			}

			logrus.WithFields(logrus.Fields{
				"原始价格": p.card.AvgPrice,
				"当前价格": p.price,
				"期望价格": newPrice,
				"调整方式": fmt.Sprintf("%v %v", updatePriceFlags.UpdatePolicy.Operator, updatePriceFlags.UpdatePolicy.PriceChange),
			}).Infof("检查将要更新的商品: 【%v】【%v】【%v %v】【%v】", p.productID, p.cardVersionID, p.card.Serial, p.cardNameCN, p.card.Rarity)

			p.price = newPrice

			if productsFlags.isRealRun {
				updateRun(&p)
			}

			needHandleCount++
		}
	}

	// 注意：总数不等于任何数量之和。
	logrus.WithFields(logrus.Fields{
		"需要处理数":      needHandleCount,
		"处理成功":       updateSuccessCount,
		"处理失败":       updateFailCount,
		"未匹配到商品的卡牌数": updateSkip,
	}).Infof("%v 张卡牌匹配到 %v 件商品的更新结果", len(cards.Data), ps.count)
}
