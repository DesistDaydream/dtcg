package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/internal/database"
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdatePriceFlags struct {
	UpdatePolicy UpdatePricePolicy
	Remark       string
}

type UpdatePricePolicy struct {
	PriceRange  []float64
	PriceChange float64
	Operator    string
	isArt       string
}

var updatePriceFlags UpdatePriceFlags

func UpdatePriceCommand() *cobra.Command {
	long := `
根据策略更新商品价格。
比如：
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 0,0.99 表示将所有价格在 0-0.99 之间卡牌的价格不增加，以集换价售卖。
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 1,9.99 -c 0.5 表示将所有价格在 1-9.99 之间卡牌的价格增加 0.5 元。
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 3,9.99 --art="否" -c 2  表示将所有价格在 3-9.99 之间的非异画卡牌的价格增加 0.5 元。
	go run cmd/jhs_cli/main.go products update price -s BTC-04 -r 0,10000 --art="是" -c 50 表示将所有价格在 0-10000 之间的异画卡价格增加 50 元。
`
	UpdateProductsPriceCmd := &cobra.Command{
		Use:   "price",
		Short: "更新商品价格",
		Long:  long,
		Run:   updatePrice,
	}

	UpdateProductsPriceCmd.Flags().Float64SliceVarP(&updatePriceFlags.UpdatePolicy.PriceRange, "price-range", "r", nil, "更新策略，卡牌价格区间。")
	UpdateProductsPriceCmd.Flags().Float64VarP(&updatePriceFlags.UpdatePolicy.PriceChange, "price-change", "c", 0, "卡牌需要变化的价格。")
	UpdateProductsPriceCmd.Flags().StringVarP(&updatePriceFlags.UpdatePolicy.Operator, "operator", "o", "+", "卡牌价格变化的计算方式，乘法还是加法。")
	UpdateProductsPriceCmd.Flags().StringVar(&updatePriceFlags.UpdatePolicy.isArt, "art", "", "是否更新异画，可用的值有两个：是、否。空值为更新所有卡牌")
	UpdateProductsPriceCmd.Flags().StringVar(&updatePriceFlags.Remark, "remark", "", "商品备注信息")

	return UpdateProductsPriceCmd
}

func updatePrice(cmd *cobra.Command, args []string) {
	if updatePriceFlags.UpdatePolicy.PriceRange == nil {
		logrus.Error("请指定要更新的卡牌价格区间。比如 -r 0,2.99 表示将所有价格在 0-2.99 之间卡牌。")
		return
	}

	// 生成待处理的卡牌信息
	cards, err := GenNeedHandleCards(updatePriceFlags.UpdatePolicy.PriceRange, updatePriceFlags.UpdatePolicy.isArt)
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}
	logrus.Infof("%v 价格区间中共有 %v 张卡牌需要更新", updatePriceFlags.UpdatePolicy.PriceRange, len(cards.Data))

	// 根据更新策略更新卡牌价格
	genNeedHandleProducts(cards, updatePriceFlags.UpdatePolicy.PriceChange)
}

// 生成待处理的商品信息
func genNeedHandleProducts(cards *dbmodels.CardsPrice, priceChange float64) {
	// TODO: 下面这俩接口有各自的优缺点，还有什么其他的好用的接口么，可以通过卡牌的唯一ID获取到商品信息~~o(╯□╰)o

	// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到商品信息
	// for _, card := range cards.Data {
	// 	products, err := handler.H.JhsServices.Products.List("1", card.Serial, "0")
	// 	if err != nil {
	// 		logrus.Fatal(err)
	// 	}
	// 	// 通过卡牌编号获取到的商品信息不是唯一的，有异画的可能，所以需要先获取商品中的 card_version_id，同时获取到 product_id(商品ID)
	// 	// 此时需要根据 card_version_id 获取到卡牌的价格信息，然后根据价格信息判断要更新的是哪个商品
	// 	for _, product := range products.Data {
	// 		cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(product.CardVersionID))
	// 		if err != nil {
	// 			logrus.Errorf("获取 %v 价格失败：%v", product.CardNameCn, err)
	// 		}
	// 		if cardPrice.AvgPrice >= avgPriceRange[0] && cardPrice.AvgPrice <= avgPriceRange[1] {
	// 			logrus.WithFields(logrus.Fields{
	// 				"原始价格": cardPrice.AvgPrice,
	// 				"更新价格": cardPrice.AvgPrice + priceChange,
	// 			}).Infof("商品【%v】【%v %v】将要调整 %v 元", card.AlternativeArt, card.Serial, product.CardNameCn, priceChange)
	// 			// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
	// 			if updateFlags.isUpdate {
	// 				updateRun(&product, cardPrice, priceChange)
	// 			}
	// 		}
	// 	}
	// }

	// 使用 /api/market/products/bySellerCardVersionId 接口时提交卖家 ID 和 card_version_id，即可获得唯一指定卡牌的商品信息，而不用其他逻辑来判断该卡牌是原画还是异画。
	// 然后，只需要遍历修改这些商品即可。
	// 但是，该接口只能获取到在售的商品，已经下架的商品无法获取到。所以想要修改下架后的商品价格或者让商品的状态变为在售或下架，是不可能的。
	for _, card := range cards.Data {
		products, err := handler.H.JhsServices.Products.Get(fmt.Sprint(card.CardVersionID), updateFlags.SellerUserID)
		if err != nil {
			logrus.Fatal(err)
		}
		cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(card.CardVersionID))
		if err != nil {
			logrus.Errorf("获取 %v 价格失败：%v", card.ScName, err)
		}

		var newPrice string

		if updatePriceFlags.UpdatePolicy.Operator == "*" {
			newPrice = fmt.Sprintf("%.2f", cardPrice.AvgPrice*priceChange)
		} else if updatePriceFlags.UpdatePolicy.Operator == "+" {
			newPrice = fmt.Sprintf("%.2f", cardPrice.AvgPrice+priceChange)
		}

		for _, product := range products.Products {
			logrus.WithFields(logrus.Fields{
				"原始价格": cardPrice.AvgPrice,
				"更新价格": newPrice,
				"调整价格": priceChange,
			}).Infof("更新前检查【%v】【%v %v】商品，使用【%v】运算符", card.AlternativeArt, card.Serial, product.CardNameCn, updatePriceFlags.UpdatePolicy.Operator)
			// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
			if updateFlags.isRealRun {
				updateRun(&product, cardPrice, newPrice)
			}
		}
	}
}

// func updateRun(product *models.ProductListData, cardPrice *databasemodels.CardPrice, priceChange float64) {
func updateRun(product *models.ProductData, cardPrice *dbmodels.CardPrice, newPrice string) {
	// 生成备注信息
	var remark string
	if updatePriceFlags.Remark != "" {
		remark = updatePriceFlags.Remark
	} else {
		remark = product.Remark
	}

	resp, err := handler.H.JhsServices.Products.Update(&models.ProductsUpdateReqBody{
		AuthenticatorID:         "",
		Grading:                 "",
		Condition:               fmt.Sprint(product.Condition),
		OnSale:                  "1",
		Price:                   newPrice,
		ProductCardVersionImage: cardPrice.ImageUrl,
		Quantity:                fmt.Sprint(product.Quantity),
		Remark:                  remark,
	}, fmt.Sprint(product.ProductID))
	if err != nil {
		logrus.Errorf("商品 %v %v 修改失败：%v", product.ProductID, product.CardNameCn, err)
	} else {
		logrus.Infof("商品 %v %v 修改成功：%v", product.ProductID, product.CardNameCn, resp)
	}
}
