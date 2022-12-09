package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/internal/database"
	databasemodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	SetPrefix    []string
	isUpdate     bool
	UpdatePolicy UpdatePolicy
}

type UpdatePolicy struct {
	PriceRange  []float64
	PriceChange float64
	isArt       string
}

var updateFlags UpdateFlags

func UpdateCommand() *cobra.Command {
	long := `
根据策略更新我在卖的卡片。
比如：
  jhs_cli products update -s BTC-03 -r 0,2.99 表示将所有价格在 0-2.99 之间卡牌的价格增加 0.1 元。
  jhs_cli products update -s BTC-03 -r 3,9.99 -c 0.5 --art="否" 表示将所有价格在 3-9.99 之间的非异画卡牌的价格增加 0.5 元。
  jhs_cli products update -s BTC-03 -r 10,50 -c 5 --art="是" 表示将所有价格在 0-1000 之间的异画卡牌的价格增加 5 元。
  jhs_cli products update -s BTC-03 -r 50.01,1000 -c 10 表示将所有价格在 50.01-1000 之间的异画卡牌的价格增加 10 元。
`
	updateProductsCmd := &cobra.Command{
		Use:              "update",
		Short:            "更新我在卖的卡片",
		Long:             long,
		Run:              updateProducts,
		PersistentPreRun: updatePersistentPreRun,
	}

	updateProductsCmd.Flags().StringSliceVarP(&updateFlags.SetPrefix, "sets-name", "s", nil, "要上架哪些卡包的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	updateProductsCmd.Flags().BoolVarP(&updateFlags.isUpdate, "update", "u", false, "是否真实更新卡牌信息。")
	updateProductsCmd.Flags().Float64SliceVarP(&updateFlags.UpdatePolicy.PriceRange, "price-range", "r", nil, "更新策略，卡牌价格区间。")
	updateProductsCmd.Flags().Float64VarP(&updateFlags.UpdatePolicy.PriceChange, "price-change", "c", 0, "卡牌需要变化的价格。")
	updateProductsCmd.Flags().StringVar(&updateFlags.UpdatePolicy.isArt, "art", "", "是否更新异画，可用的值有两个：是、否。空值为更新所有卡牌")

	updateProductsCmd.AddCommand(
		UpdateImageCommand(),
	)

	return updateProductsCmd
}

func updatePersistentPreRun(cmd *cobra.Command, args []string) {
	// 执行父命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}
}

func updateProducts(cmd *cobra.Command, args []string) {
	if updateFlags.SetPrefix == nil {
		logrus.Error("请指定要更新的卡牌集合，使用 dtcg_cli card-set list 子命令获取卡包名称。")
		return
	}

	if updateFlags.UpdatePolicy.PriceRange == nil {
		logrus.Error("请指定要更新的卡牌价格区间。比如 -r 0,2.99 表示将所有价格在 0-2.99 之间卡牌。")
		return
	}

	// 根据更新策略更新卡牌价格
	genNeedUpdateProducts(updateFlags.UpdatePolicy.PriceRange, updateFlags.UpdatePolicy.isArt, updateFlags.UpdatePolicy.PriceChange)
}

// 生成需要更新的卡牌信息
func genNeedUpdateProducts(avgPriceRange []float64, alternativeArt string, priceChange float64) {
	var (
		cards *databasemodels.CardsPrice
		err   error
	)

	// 生成需要更新的卡牌信息
	cards, err = database.GetCardPriceByCondition(200, 1, &databasemodels.CardPriceQuery{
		SetsPrefix:     updateFlags.SetPrefix,
		AlternativeArt: alternativeArt,
		MinPriceRange:  "",
		AvgPriceRange:  fmt.Sprintf("%v-%v", avgPriceRange[0], avgPriceRange[1]),
	})
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}

	logrus.Infof("%v 价格区间中共有 %v 张卡牌需要更新", avgPriceRange, len(cards.Data))

	for _, card := range cards.Data {
		// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到商品信息
		products, err := handler.H.JhsServices.Products.List("1", card.Serial)
		if err != nil {
			logrus.Fatal(err)
		}

		// 通过卡牌编号获取到的商品信息不是唯一的，有异画的可能，所以需要先获取商品中的 card_version_id，同时获取到 product_id(商品ID)
		// 此时需要根据 card_version_id 获取到卡牌的价格信息，然后根据价格信息判断要更新的是哪个商品
		for _, product := range products.Data {
			cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(product.CardVersionID))
			if err != nil {
				logrus.Errorf("获取 %v 价格失败：%v", product.CardNameCn, err)
			}

			if cardPrice.AvgPrice >= avgPriceRange[0] && cardPrice.AvgPrice <= avgPriceRange[1] {
				// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
				if updateFlags.isUpdate {
					updateRun(&product, cardPrice, priceChange)
				}
				logrus.WithFields(logrus.Fields{
					"原始价格": cardPrice.AvgPrice,
					"更新价格": cardPrice.AvgPrice + priceChange,
				}).Infof("商品【%v】【%v %v】将要调整 %v 元", card.AlternativeArt, card.Serial, product.CardNameCn, priceChange)
			}
		}
	}
}

func updateRun(product *models.ProductListData, cardPrice *databasemodels.CardPrice, priceChange float64) {
	resp, err := handler.H.JhsServices.Products.Update(&models.ProductsUpdateReqBody{
		Condition:            fmt.Sprint(product.Condition),
		OnSale:               fmt.Sprint(product.OnSale),
		Price:                fmt.Sprint(cardPrice.AvgPrice + priceChange),
		Quantity:             fmt.Sprint(product.Quantity),
		Remark:               "",
		UserCardVersionImage: cardPrice.ImageUrl,
	}, fmt.Sprint(product.ProductID))
	if err != nil {
		logrus.Errorf("商品 %v %v 修改失败：%v", product.ProductID, product.CardNameCn, err)
	} else {
		logrus.Infof("商品 %v %v 修改成功：%v", product.ProductID, product.CardNameCn, resp)
	}
}
