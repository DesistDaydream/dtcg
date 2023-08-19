package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
	pmodels "github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	SellerUserID string // 集换社卖家 ID
	CurSaleState string // 当前商品的售卖状态
	ExpSaleState string // 期望商品变成哪种售卖状态
	Remark       string // 商品备注
}

var updateFlags UpdateFlags

func UpdateCommand() *cobra.Command {

	updateProductsCmd := &cobra.Command{
		Use:   "update",
		Short: "更新商品",
		// Run:   updatePrice,
		PersistentPreRun: updatePersistentPreRun,
	}

	updateProductsCmd.PersistentFlags().StringVarP(&updateFlags.SellerUserID, "seller-user-id", "u", "934972", "卖家用户ID。")
	updateProductsCmd.PersistentFlags().StringVar(&updateFlags.CurSaleState, "cur-sale-state", "1", "当前售卖状态。即获取什么状态的商品。1: 售卖。0: 下架")
	updateProductsCmd.PersistentFlags().StringVar(&updateFlags.ExpSaleState, "exp-sale-state", "1", "期望的售卖状态。")
	updateProductsCmd.PersistentFlags().StringVar(&updateFlags.Remark, "remark", "", "商品备注信息")

	updateProductsCmd.AddCommand(
		UpdatePriceCommand(),
		UpdateImageCommand(),
		UpdateSaleStateCommand(),
		UpdateQuantityCommand(),
	)

	return updateProductsCmd
}

func updatePersistentPreRun(cmd *cobra.Command, args []string) {
	if productsFlags.SetPrefix == nil && productsFlags.CardVersionID == 0 {
		logrus.Fatalln("请指定要更新的卡牌，可以使用 dtcg_cli card-set list 子命令获取卡包名称；或者直接指定卡牌的 card_version_id。")
	}
}

type NeedHandleProducts struct {
	count    int
	products []Product
}

type Product struct {
	card    dbmodels.CardPrice
	product models.ProductListData
	// 更新商品用的新数据，并不是一定会用上。主要用于不同更新场景时使用
	newOnSale   string // 根据命令行标志设置商品在售状态
	newPrice    string // 根据条件生成商品价格
	newImg      string // 从数据库的 card_prices 表中获取卡图
	newQuantity string // 根据条件生成商品数量
}

// 生成需要更新的商品信息
func genNeedUpdateProducts(cards *dbmodels.CardsPrice, priceChange float64) *NeedHandleProducts {
	var needHandleProducts NeedHandleProducts

	// 逐一生成待处理卡牌的商品信息
	for _, card := range cards.Data {
		// 用于记录待处理的卡牌的商品是否已更新
		isUpdate := 0
		// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到该卡牌的商品列表
		products, err := handler.H.JhsServices.Market.SellersProductsList(1, card.Serial, updateFlags.CurSaleState, "published_at_desc")
		if err != nil {
			logrus.Errorf("获取 %v 卡牌的商品失败: %v", card.ScName, err)
			updateFailCount++
			continue
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

		// 通过卡牌编号获取到的商品信息不是唯一的，这个编号的卡有可能包含异画，所以需要先获取商品中的 card_version_id，
		// 然后将商品的 card_version_id 与当前待更新卡牌的 card_version_id 对比，以确定唯一的 product_id(商品ID)
		for _, p := range products.Data {
			if p.CardVersionID == card.CardVersionID {
				// 生成商品将要更新的数量
				var newQuantity string
				if p.Quantity < 4 {
					newQuantity = updateQuantityFlags.UpdatePolicy.ProductQuantity
				} else {
					newQuantity = fmt.Sprintf("%v", p.Quantity)
				}

				needHandleProducts.products = append(needHandleProducts.products, Product{
					card:        card,
					product:     p,
					newOnSale:   updateFlags.ExpSaleState,
					newPrice:    newPrice,
					newImg:      card.ImageUrl,
					newQuantity: newQuantity,
				})

				logrus.WithFields(logrus.Fields{
					"原始价格": card.AvgPrice,
					"更新价格": newPrice,
					"调整价格": fmt.Sprintf("%v %v", updatePriceFlags.UpdatePolicy.Operator, updatePriceFlags.UpdatePolicy.PriceChange),
				}).Infof("更新前检查商品: 【%v】【%v】【%v %v】", p.CardVersionID, card.AlternativeArt, card.Serial, p.CardNameCN)

				needHandleProducts.count++

				isUpdate = 1
				break
			} else {
				logrus.Debugf("当前商品 [%v %v %v %v] 与期望处理的商品 [%v %v %v %v] 不匹配，跳过",
					p.CardVersionID, p.CardVersionNumber, p.CardNameCN, p.CardVersionRarity,
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

	return &needHandleProducts
}

// TODO: 下面这个接口与 genNeedUpdateProducts 接口各有优缺点，还有什么其他的好用的接口么，可以通过卡牌的唯一ID获取到商品信息~~o(╯□╰)o
// 生成需要更新的商品信息
func genNeedUpdateProductsWithBySellerCardVersionId(cards *dbmodels.CardsPrice, priceChange float64) {
	for _, card := range cards.Data {
		// 使用 /api/market/products/bySellerCardVersionId 接口时提交卖家 ID 和 card_version_id，即可获得唯一指定卡牌的商品信息，而不用其他逻辑来判断该卡牌是原画还是异画。
		// 然后，只需要遍历修改这些商品即可。
		// 但是，该接口只能获取到在售的商品，已经下架的商品无法获取到。所以想要修改下架后的商品价格或者让商品的状态变为在售或下架，是不可能的。
		// 后来，官方添加了一个 default_product 的字段，这里也可以获得 product_id、价格、等 信息。都不用使用 for 轮询商品了，这点还是不错的。但是感觉用起来还是有点奇怪，待补充...
		products, err := handler.H.JhsServices.Products.Get(fmt.Sprint(card.CardVersionID), updateFlags.SellerUserID)
		if err != nil {
			logrus.Errorf("获取 %v 卡牌的商品失败: %v", card.ScName, err)
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

		logrus.WithFields(logrus.Fields{
			"原始价格": cardPrice.AvgPrice,
			"更新价格": newPrice,
			"调整价格": priceChange,
		}).Infof("更新前检查【%v】【%v %v】商品，使用【%v】运算符", card.AlternativeArt, card.Serial, products.DefaultProduct.CardNameCN, updatePriceFlags.UpdatePolicy.Operator)
		if productsFlags.isRealRun {
			updateRunWithDefaultProdcut(&products.DefaultProduct, card.ImageUrl, newPrice)
		}
	}
}

// 使用默认商品信息更新商品
func updateRunWithDefaultProdcut(product *pmodels.DefaultProduct, imageUrl string, price string) {
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
		Default:                 "1",
		OnSale:                  updateFlags.ExpSaleState,
		Price:                   price,
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

// 更新商品
func updateRun(product *models.ProductListData, onSale, price, imageUrl, quantity string) {
	// 生成备注信息
	var remark string
	if updateFlags.Remark != "" {
		remark = updateFlags.Remark
	} else {
		remark = product.Remark
	}

	// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
	resp, err := handler.H.JhsServices.Market.SellersProductsUpdate(&models.ProductsUpdateReqBody{
		AuthenticatorID:         "",
		Grading:                 "",
		Condition:               fmt.Sprint(product.Condition),
		Default:                 "1",
		OnSale:                  onSale,
		Price:                   price,
		ProductCardVersionImage: imageUrl,
		Quantity:                quantity,
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
