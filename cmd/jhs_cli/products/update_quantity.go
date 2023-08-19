package products

import (
	"fmt"

	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateQuantityFlags struct {
	UpdatePolicy UpdateQuantityPolicy
}

type UpdateQuantityPolicy struct {
	ProductQuantity string
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

	UpdateProductsPriceCmd.Flags().StringVarP(&updateQuantityFlags.UpdatePolicy.ProductQuantity, "price-change", "q", "4", "卡牌需要变化的价格。")

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
	genNeedHandleQuantityProducts(cards)

	// 注意：总数不等于任何数量之和。
	logrus.WithFields(logrus.Fields{
		"总数": cards.Count,
		"成功": updateSuccessCount,
		"失败": updateFailCount,
		"跳过": updateSkip,
	}).Infof("更新结果")
}

// TODO: 下面这个接口与另一个 genNeedHandleProducts 接口各有优缺点，还有什么其他的好用的接口么，可以通过卡牌的唯一ID获取到商品信息~~o(╯□╰)o
// 生成待处理的商品信息
func genNeedHandleQuantityProducts(cards *dbmodels.CardsPrice) {
	for _, card := range cards.Data {
		products, err := handler.H.JhsServices.Products.Get(fmt.Sprint(card.CardVersionID), updateFlags.SellerUserID)
		if err != nil {
			logrus.Fatal(err)
		}

		if products.DefaultProduct.Quantity == 0 {
			logrus.Warnf("当前商品 %v %v 数量为 0，跳过", products.DefaultProduct.CardNameCN, card.CardVersionID)
			updateSkip++
			continue
		}

		var newQuantity string
		if products.DefaultProduct.Quantity < 4 {
			newQuantity = updateQuantityFlags.UpdatePolicy.ProductQuantity
		} else {
			newQuantity = fmt.Sprintf("%v", products.DefaultProduct.Quantity)
		}

		logrus.WithFields(logrus.Fields{
			"当前售卖数量": products.DefaultProduct.Quantity,
			"期望售卖数量": newQuantity,
		}).Infof("更新前检查【%v】【%v %v】商品", card.AlternativeArt, card.Serial, products.DefaultProduct.CardNameCN)

		if productsFlags.isRealRun {
			resp, err := handler.H.JhsServices.Market.SellersProductsUpdate(&models.ProductsUpdateReqBody{
				AuthenticatorID:         "",
				Grading:                 "",
				Condition:               fmt.Sprint(products.DefaultProduct.Condition),
				Default:                 "1",
				OnSale:                  updateFlags.ExpSaleState,
				Price:                   fmt.Sprintf("%.2f", products.DefaultProduct.Price),
				ProductCardVersionImage: card.ImageUrl,
				Quantity:                newQuantity,
				Remark:                  products.DefaultProduct.Remark,
			}, fmt.Sprint(products.DefaultProduct.ProductID))
			if err != nil {
				logrus.Errorf("商品 %v %v 修改失败：%v", products.DefaultProduct.ProductID, products.DefaultProduct.CardNameCN, err)
				updateFailCount++
			} else {
				logrus.Infof("商品 %v %v 修改成功：%v", products.DefaultProduct.ProductID, products.DefaultProduct.CardNameCN, resp)
				updateSuccessCount++
			}
		}
	}
}
