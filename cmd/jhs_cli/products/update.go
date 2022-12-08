package products

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/internal/database"
	databasemodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	SetPrefix    []string
	UpdatePolicy UpdatePolicy
}

type UpdatePolicy struct {
	// PriceRange        string
	Less3              bool
	Greater3AndLess10  bool
	Greater10AndLess50 bool
	Greater50          bool
}

var updateFlags UpdateFlags

func UpdateCommand() *cobra.Command {
	updateProductsCmd := &cobra.Command{
		Use:              "update",
		Short:            "更新我在卖的卡片",
		Run:              updateProducts,
		PersistentPreRun: updatePersistentPreRun,
	}

	updateProductsCmd.Flags().StringSliceVarP(&updateFlags.SetPrefix, "sets-name", "s", nil, "要上架哪些卡包的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	// updateProductsCmd.Flags().StringVarP(&updateFlags.UpdatePolicy.PriceRange, "price-range", "p", "", "更新策略，卡牌价格区间。")
	updateProductsCmd.Flags().BoolVar(&updateFlags.UpdatePolicy.Less3, "0", false, "更新策略，小于 3 元的卡牌。")
	updateProductsCmd.Flags().BoolVar(&updateFlags.UpdatePolicy.Greater3AndLess10, "3", false, "更新策略，大于等于 3 元小于 10 元的卡牌。")
	updateProductsCmd.Flags().BoolVar(&updateFlags.UpdatePolicy.Greater10AndLess50, "10", false, "更新策略，大于等于 10 元小于 50 元的卡牌。")
	updateProductsCmd.Flags().BoolVar(&updateFlags.UpdatePolicy.Greater50, "50", false, "更新策略，大于 50 元的卡牌。")

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

	// 根据更新策略更新的卡牌价格
	switch policy := updateFlags.UpdatePolicy; {
	case policy.Less3:
		genNeedUpdateProducts("0-2.99", "", 0)
	case policy.Greater3AndLess10:
		genNeedUpdateProducts("3-3.99", "否", 0.5)
	case policy.Greater10AndLess50:
		genNeedUpdateProducts("10-50", "是", 5)
	case policy.Greater50:
		genNeedUpdateProducts("50.01-10000", "", 10)
	default:
		// updateRun(&product, cardPrice, 5)
	}
}

// 生成需要更新的卡牌信息
func genNeedUpdateProducts(avgPriceRange string, alternativeArt string, risingPrices float64) {
	var (
		cards      *databasemodels.CardsPrice
		err        error
		priceRange = strings.Split(avgPriceRange, "-")
	)

	// 将 priceRange 转为 float64 类型切片
	floatPriceRange := make([]float64, len(priceRange))

	for i, str := range priceRange {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			logrus.Errorf("%v", err)
		}
		floatPriceRange[i] = f
	}

	// 生成需要更新的卡牌信息
	cards, err = database.GetCardPriceByCondition(200, 1, &databasemodels.CardPriceQuery{
		SetsPrefix:     updateFlags.SetPrefix,
		AlternativeArt: alternativeArt,
		MinPriceRange:  "",
		AvgPriceRange:  avgPriceRange,
	})
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}

	logrus.Infof("【%v】价格区间中共有 %v 张卡牌需要更新", avgPriceRange, len(cards.Data))

	// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到商品信息
	for _, card := range cards.Data {
		products, err := handler.H.JhsServices.Products.List("1", card.Serial)
		if err != nil {
			logrus.Fatal(err)
		}

		// 通过卡牌编号获取到的商品有异画的可能，所以需要先获取商品中的 card_version_id，同时获取到 product_id(商品ID)
		// 然后还需要再判断一下价格区间，防止更新到价格不在区间内的商品
		for _, product := range products.Data {
			cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(product.CardVersionID))
			if err != nil {
				logrus.Errorf("获取 %v 价格失败：%v", product.CardNameCn, err)
			}

			if cardPrice.AvgPrice >= floatPriceRange[0] && cardPrice.AvgPrice <= floatPriceRange[1] {
				// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
				// updateRun(&product, cardPrice, risingPrices)
				logrus.WithFields(logrus.Fields{
					"原始价格": cardPrice.AvgPrice,
					"更新价格": cardPrice.AvgPrice + risingPrices,
				}).Infof("商品【%v】【%v %v】将要上调 %v 元", card.AlternativeArt, card.Serial, product.CardNameCn, risingPrices)
			}

		}
	}
}

func updateRun(product *models.ProductListData, cardPrice *databasemodels.CardPrice, price float64) {
	resp, err := handler.H.JhsServices.Products.Update(&models.ProductsUpdateReqBody{
		Condition:            fmt.Sprint(product.Condition),
		OnSale:               fmt.Sprint(product.OnSale),
		Price:                fmt.Sprint(cardPrice.AvgPrice + price),
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
