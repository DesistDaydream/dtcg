package products

import (
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateSaleStateFlags struct {
	OneByOne              bool
	UpdateSaleStatePolicy UpdateSaleStatePolicy
}

type UpdateSaleStatePolicy struct {
	isArt string
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

	updateProductsSaleStateCmd.Flags().StringVar(&updateSaleStateFlags.UpdateSaleStatePolicy.isArt, "art", "", "是否更新异画，可用的值有两个：是、否。空值为更新所有卡牌")
	updateProductsSaleStateCmd.Flags().BoolVar(&updateSaleStateFlags.OneByOne, "one-by-one", false, "是否一条一条得变更所有商品的价格")

	return updateProductsSaleStateCmd
}

// 更新商品售卖状态
func updateSaleState(cmd *cobra.Command, args []string) {
	if updateSaleStateFlags.OneByOne {
		if productsFlags.isRealRun {
			updateSaleStateOneByOne()
		}
	} else {
		// 生成待处理的卡牌信息
		cards, err := GenNeedHandleCards()
		if err != nil {
			logrus.Errorf("%v", err)
			return
		}

		// 根据更新策略更新卡牌价格
		genNeedUpdateSaleStateProducts(cards, updateSaleStateFlags.UpdateSaleStatePolicy.isArt)
	}
}

// 生成需要更新的卡牌信息
func genNeedUpdateSaleStateProducts(cards *dbmodels.CardsPrice, alternativeArt string) {
	// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到商品信息
	for _, card := range cards.Data {
		products, err := handler.H.JhsServices.Market.SellersProductsList(1, card.Serial, updateFlags.CurSaleState, "published_at_desc")
		if err != nil {
			logrus.Fatal(err)
		}
		for _, p := range products.Data {
			logrus.WithFields(logrus.Fields{
				"当前状态": updateFlags.CurSaleState,
				"预期状态": updateFlags.ExpSaleState,
			}).Infof("更新前检查【%v】【%v %v】商品", card.AlternativeArt, card.Serial, p.CardNameCN)
			// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
			if productsFlags.isRealRun {
				updateRun(&Product{
					card:      &dbmodels.CardPrice{},
					product:   &p,
					productID: p.ProductID,
					onSale:    updateFlags.ExpSaleState,
					price:     p.Price,
					img:       p.CardVersionImage,
					quantity:  p.Quantity,
				})
			}
		}
	}
}

// 逐一更新所有商品的售卖状态
func updateSaleStateOneByOne() {
	page := 1 // 从获取到的数据的第一页开始
	for {
		products, err := handler.H.JhsServices.Market.SellersProductsList(page, "", updateFlags.CurSaleState, "published_at_desc")
		if err != nil || len(products.Data) <= 0 {
			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
		}
		for _, p := range products.Data {
			if p.Quantity != 0 {
				updateRun(&Product{
					card:      &dbmodels.CardPrice{},
					product:   &p,
					productID: p.ProductID,
					onSale:    updateFlags.ExpSaleState,
					price:     p.Price,
					img:       p.CardVersionImage,
					quantity:  p.Quantity,
				})
			}
		}
		logrus.Infof("共 %v 页数据，已处理第 %v 页", products.LastPage, products.CurrentPage)
		// 如果当前处理的页等于最后页，则退出循环
		if products.CurrentPage == products.LastPage {
			logrus.Debugf("退出循环时共 %v 页,处理完 %v 页", products.LastPage, products.CurrentPage)
			break
		}
		// 每处理完一页，下一个循环需要处理的页+1
		page++
	}
}
