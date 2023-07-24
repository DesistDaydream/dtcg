package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdatePriceFlags struct {
	UpdatePolicy UpdatePricePolicy
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

	// 根据更新策略更新卡牌价格
	genNeedHandleProducts(cards, updatePriceFlags.UpdatePolicy.PriceChange)

	// 注意：总数不等于任何数量之和。
	logrus.WithFields(logrus.Fields{
		"总数": cards.Count,
		"成功": updateSuccessCount,
		"失败": updateFailCount,
		"跳过": updateSkip,
	}).Infof("更新结果")
}

// 生成待处理的商品信息
func genNeedHandleProducts(cards *dbmodels.CardsPrice, priceChange float64) {
	// 逐一更新待处理卡牌的商品信息
	for _, card := range cards.Data {
		// 用于记录待处理的卡牌的商品是否已更新
		isUpdate := 0

		// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到该卡牌的商品列表
		products, err := handler.H.JhsServices.Market.SellersProductsList("1", card.Serial, updateFlags.CurSaleState, "published_at_desc")
		if err != nil {
			logrus.Fatal(err)
		}

		// 判断一下这个这个卡牌有几个商品，若商品为0，则继续处理下一个
		if len(products.Data) == 0 {
			logrus.Errorf("%v %v 没有可供处理的商品，跳过", card.CardVersionID, card.ScName)
			updateSkip++
			continue
		}

		// 生成商品将要更新的价格
		var newPrice string
		if updatePriceFlags.UpdatePolicy.Operator == "*" {
			newPrice = fmt.Sprintf("%.2f", card.AvgPrice*priceChange)
		} else if updatePriceFlags.UpdatePolicy.Operator == "+" {
			newPrice = fmt.Sprintf("%.2f", card.AvgPrice+priceChange)
		}

		for _, product := range products.Data {
			// 通过卡牌编号获取到的商品信息不是唯一的，这个编号的卡有可能包含异画，所以需要先获取商品中的 card_version_id，
			// 然后将商品的 card_version_id 与当前待更新卡牌的 card_version_id 对比，以确定唯一的 product_id(商品ID)
			if product.CardVersionID == card.CardVersionID {
				logrus.WithFields(logrus.Fields{
					"原始价格": card.AvgPrice,
					"更新价格": newPrice,
					"调整价格": fmt.Sprintf("%v %v", updatePriceFlags.UpdatePolicy.Operator, priceChange),
				}).Infof("更新前检查商品: 【%v】【%v】【%v %v】", product.CardVersionID, card.AlternativeArt, card.Serial, product.CardNameCN)
				// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
				if productsFlags.isRealRun {
					updateRun(&product, card.ImageUrl, newPrice)
				}
				isUpdate = 1
				break
			} else {
				logrus.Debugf("当前商品 [%v %v %v %v] 与期望处理的商品 [%v %v %v %v] 不匹配，跳过",
					product.CardVersionID, product.CardVersionNumber, product.CardNameCN, product.CardVersionRarity,
					card.CardVersionID, card.Serial, card.ScName, card.Rarity)
				updateSkip++
			}
		}

		// 挺尴尬的做法，获取到需要更新的卡牌列表后，只能通过名字获取商品，但是通过名称获取到的商品可能是其他卡牌的商品(各种异画)。。。o(╯□╰)o
		// 所以需要一个有状态的数据来记录待更新卡牌是否获取到商品并成功更新
		if isUpdate == 0 {
			logrus.Errorf("%v 卡牌没有商品可以更新", card.ScName)
			updateFailCount++
		}
	}
}

func updateRun(product *models.ProductListData, imageUrl, newPrice string) {
	// 这个是给 genNeedHandleProducts() 用的
	// func updateRun(product *models.ProductData, cardPrice *dbmodels.CardPrice, newPrice string) {

	// 生成备注信息
	var remark string
	if updateFlags.Remark != "" {
		remark = updateFlags.Remark
	} else {
		remark = product.Remark
	}

	resp, err := handler.H.JhsServices.Market.SellersProductsUpdate(&models.ProductsUpdateReqBody{
		AuthenticatorID:         "",
		Grading:                 "",
		Condition:               fmt.Sprint(product.Condition),
		OnSale:                  updateFlags.ExpSaleState,
		Price:                   newPrice,
		ProductCardVersionImage: imageUrl,
		Quantity:                fmt.Sprint(product.Quantity),
		Remark:                  remark,
	}, fmt.Sprint(product.ProductID))
	if err != nil {
		logrus.Errorf("商品 %v %v 修改失败：%v", product.ProductID, product.CardNameCN, err)
		updateFailCount++
	} else {
		logrus.Infof("商品 %v %v 修改成功：%v", product.ProductID, product.CardNameCN, resp)
		updateSuccessCount++
	}
}

// TODO: 下面这个接口与另一个 genNeedHandleProducts 接口各有优缺点，还有什么其他的好用的接口么，可以通过卡牌的唯一ID获取到商品信息~~o(╯□╰)o
// 生成待处理的商品信息
// func genNeedHandleProducts(cards *dbmodels.CardsPrice, priceChange float64) {
// 	for _, card := range cards.Data {
// 		// 使用 /api/market/products/bySellerCardVersionId 接口时提交卖家 ID 和 card_version_id，即可获得唯一指定卡牌的商品信息，而不用其他逻辑来判断该卡牌是原画还是异画。
// 		// 然后，只需要遍历修改这些商品即可。
// 		// 但是，该接口只能获取到在售的商品，已经下架的商品无法获取到。所以想要修改下架后的商品价格或者让商品的状态变为在售或下架，是不可能的。
// 		products, err := handler.H.JhsServices.Products.Get(fmt.Sprint(card.CardVersionID), updateFlags.SellerUserID)
// 		if err != nil {
// 			logrus.Fatal(err)
// 		}
// 		cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(card.CardVersionID))
// 		if err != nil {
// 			logrus.Errorf("获取 %v 价格失败：%v", card.ScName, err)
// 		}
// 		var newPrice string
// 		if updatePriceFlags.UpdatePolicy.Operator == "*" {
// 			newPrice = fmt.Sprintf("%.2f", cardPrice.AvgPrice*priceChange)
// 		} else if updatePriceFlags.UpdatePolicy.Operator == "+" {
// 			newPrice = fmt.Sprintf("%.2f", cardPrice.AvgPrice+priceChange)
// 		}
// 		for _, product := range products.Products {
// 			logrus.WithFields(logrus.Fields{
// 				"原始价格": cardPrice.AvgPrice,
// 				"更新价格": newPrice,
// 				"调整价格": priceChange,
// 			}).Infof("更新前检查【%v】【%v %v】商品，使用【%v】运算符", card.AlternativeArt, card.Serial, product.CardNameCN, updatePriceFlags.UpdatePolicy.Operator)
// 			// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
// 			if productsFlags.isRealRun {
// 				updateRun(&product, cardPrice, newPrice)
// 			}
// 		}
// 	}
// }
