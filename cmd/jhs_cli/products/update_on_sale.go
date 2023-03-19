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

type UpdateSaleStateFlags struct {
	OldSaleState          string
	NewSaleState          string
	UpdateSaleStatePolicy UpdateSaleStatePolicy
}

type UpdateSaleStatePolicy struct {
	PriceChange float64
	isArt       string
}

var updateSaleStateFlags UpdateSaleStateFlags

func UpdateSaleStateCommand() *cobra.Command {
	long := `
	根据策略更新商品的售卖状态。1: 售卖。0: 下架
`

	updateProductsSaleStateCmd := &cobra.Command{
		Use:   "sale-state",
		Short: "更新商品的售卖状态，是上架还是下架",
		Long:  long,
		Run:   updateSaleState,
	}

	updateProductsSaleStateCmd.Flags().StringVarP(&updateSaleStateFlags.OldSaleState, "old-sale-state", "o", "0", "当前售卖状态。即获取什么状态的商品。1: 售卖。0: 下架")
	updateProductsSaleStateCmd.Flags().StringVarP(&updateSaleStateFlags.NewSaleState, "new-sale-state", "n", "1", "期望的售卖状态。")
	updateProductsSaleStateCmd.Flags().StringVar(&updateSaleStateFlags.UpdateSaleStatePolicy.isArt, "art", "", "是否更新异画，可用的值有两个：是、否。空值为更新所有卡牌")

	return updateProductsSaleStateCmd
}

// 更新商品售卖状态
func updateSaleState(cmd *cobra.Command, args []string) {
	// if updateSaleStateFlags.UpdateSaleStatePolicy.PriceRange == nil {
	// 	logrus.Error("请指定要更新的卡牌价格区间。比如 -r 0,2.99 表示将所有价格在 0-2.99 之间卡牌。")
	// 	return
	// }

	// 根据更新策略更新卡牌价格
	genNeedUpdateSaleStateProducts(updateSaleStateFlags.UpdateSaleStatePolicy.isArt, updateSaleStateFlags.UpdateSaleStatePolicy.PriceChange)
}

// 生成需要更新的卡牌信息
func genNeedUpdateSaleStateProducts(alternativeArt string, priceChange float64) {
	var (
		cards *dbmodels.CardsPrice
		err   error
	)

	// 生成需要更新的卡牌信息
	cards, err = database.GetCardPriceByCondition(300, 1, &dbmodels.CardPriceQuery{
		SetsPrefix:     updateFlags.SetPrefix,
		AlternativeArt: alternativeArt,
		MinPriceRange:  "",
	})
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}

	logrus.Infof("共有 %v 张卡牌需要更新", len(cards.Data))

	// TODO: 下面这俩接口有各自的优缺点，还有什么其他的好用的接口么，可以通过卡牌的唯一ID获取到商品信息~~o(╯□╰)o

	// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到商品信息
	for _, card := range cards.Data {
		products, err := handler.H.JhsServices.Products.List("1", card.Serial, updateSaleStateFlags.OldSaleState)
		if err != nil {
			logrus.Fatal(err)
		}
		// 通过卡牌编号获取到的商品信息不是唯一的，有异画的可能，所以需要先获取商品中的 card_version_id，同时获取到 product_id(商品ID)
		for _, product := range products.Data {
			logrus.WithFields(logrus.Fields{
				"当前状态": updateSaleStateFlags.OldSaleState,
				"预期状态": updateSaleStateFlags.NewSaleState,
			}).Infof("更新前检查【%v】【%v %v】商品", card.AlternativeArt, card.Serial, product.CardNameCn)
			// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
			if updateFlags.isRealRun {
				updateSaleStateRun(&product, priceChange, updateSaleStateFlags.NewSaleState)
			}
		}
	}

	// 使用 /api/market/products/bySellerCardVersionId 接口时提交卖家 ID 和 card_version_id，即可获得唯一指定卡牌的商品信息，而不用其他逻辑来判断该卡牌是原画还是异画。
	// 然后，只需要遍历修改这些商品即可。
	// 但是，该接口只能获取到在售的商品，已经下架的商品无法获取到。所以想要修改下架后的商品价格或者让商品的状态变为在售或下架，是不可能的。
	// for _, card := range cards.Data {
	// 	products, err := handler.H.JhsServices.Products.Get(fmt.Sprint(card.CardVersionID), updateSaleStateFlags.SellerUserID)
	// 	if err != nil {
	// 		logrus.Fatal(err)
	// 	}
	// 	for _, product := range products.Products {
	// 		logrus.WithFields(logrus.Fields{}).Infof("商品【%v】【%v %v】将要变为 %v 状态", card.AlternativeArt, card.Serial, product.CardNameCn, updateSaleStateFlags.NewSaleState)
	// 		// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
	// 		if updateSaleStateFlags.isUpdate {
	// 			updateSaleStateRun(&product, priceChange)
	// 		}
	// 	}
	// }
}

func updateSaleStateRun(product *models.ProductListData, priceChange float64, onSaleState string) {
	// func updateSaleStateRun(product *models.ProductData, priceChange float64) {
	resp, err := handler.H.JhsServices.Products.Update(&models.ProductsUpdateReqBody{
		AuthenticatorID:         "",
		Grading:                 "",
		Condition:               fmt.Sprint(product.Condition),
		OnSale:                  onSaleState,
		Price:                   fmt.Sprint(product.Price),
		ProductCardVersionImage: product.CardVersionImage,
		Quantity:                fmt.Sprint(product.Quantity),
		Remark:                  product.Remark,
	}, fmt.Sprint(product.ProductID))
	if err != nil {
		logrus.Errorf("商品 %v %v 更新失败：%v", product.ProductID, product.CardNameCn, err)
	} else {
		logrus.Infof("商品 %v %v 更新成功：%v", product.ProductID, product.CardNameCn, resp)
	}
}

// func updateSaleStateOneByOne() {
// 	page := 1 // 从获取到的数据的第一页开始
// 	for {
// 		products, err := handler.H.JhsServices.Products.List(strconv.Itoa(page), "", updateSaleStateFlags.OldSaleState)
// 		if err != nil || len(products.Data) <= 0 {
// 			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
// 		}
// 		for _, product := range products.Data {
// 			if product.Quantity != 0 {
// 				resp, err := handler.H.JhsServices.Products.Update(&models.ProductsUpdateReqBody{
// 					AuthenticatorID:         "",
// 					Grading:                 "",
// 					Condition:               fmt.Sprint(product.Condition),
// 					OnSale:                  updateSaleStateFlags.NewSaleState,
// 					Price:                   product.Price,
// 					ProductCardVersionImage: product.CardVersionImage,
// 					Quantity:                fmt.Sprint(product.Quantity),
// 					Remark:                  product.Remark,
// 				}, fmt.Sprint(product.ProductID))
// 				if err != nil {
// 					logrus.Errorf("商品 %v %v 修改失败：%v", product.ProductID, product.CardNameCn, err)
// 				} else {
// 					logrus.Infof("商品 %v %v 修改成功：%v", product.ProductID, product.CardNameCn, resp)
// 				}
// 			}
// 		}
// 		logrus.Infof("共 %v 页数据，已处理第 %v 页", products.LastPage, products.CurrentPage)
// 		// 如果当前处理的页等于最后页，则退出循环
// 		if products.CurrentPage == products.LastPage {
// 			logrus.Debugf("退出循环时共 %v 页,处理完 %v 页", products.LastPage, products.CurrentPage)
// 			break
// 		}
// 		// 每处理完一页，下一个循环需要处理的页+1
// 		page++
// 	}
// }
