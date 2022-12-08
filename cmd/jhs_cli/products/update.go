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
	UpdatePolicy UpdatePolicy
}

type UpdatePolicy struct {
	Less3             bool
	Greater3AndLess10 bool
	Greater5AndLess50 bool
	Greater50         bool
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
	updateProductsCmd.Flags().BoolVar(&updateFlags.UpdatePolicy.Less3, "1", false, "更新策略，小于 5 元的卡牌。")
	updateProductsCmd.Flags().BoolVar(&updateFlags.UpdatePolicy.Greater3AndLess10, "3", false, "更新策略，大于 3 元小于 10 元的卡牌。")
	updateProductsCmd.Flags().BoolVar(&updateFlags.UpdatePolicy.Greater5AndLess50, "5", false, "更新策略，大于 5 元小于 50 元的卡牌。")
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

func genNeedUpdateProducts(setsPrefix []string) ([]string, error) {
	// 生成需要更新的卡牌信息
	var serials []string

	for _, setPrefix := range setsPrefix {
		cardsPrice, err := database.GetCardPriceWhereSetPrefix(setPrefix)
		if err != nil {
			return nil, err
		}

		for _, cardPrice := range cardsPrice.Data {
			serials = append(serials, cardPrice.Serial)
		}
	}

	logrus.Debugf("当前需要更新 %v 张卡牌：%v", len(serials), serials)
	return serials, nil
}

func updateProducts(cmd *cobra.Command, args []string) {
	if updateFlags.SetPrefix == nil {
		logrus.Error("请指定要更新的卡牌集合，使用 dtcg_cli card-set list 子命令获取卡包名称。")
		return
	}

	// 生成需要更新的卡牌信息
	cards, err := genNeedUpdateProducts(updateFlags.SetPrefix)
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}

	// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到商品信息
	for _, card := range cards {
		products, err := handler.H.JhsServices.Products.List("1", card)
		if err != nil {
			logrus.Fatal(err)
		}

		for _, product := range products.Data {
			// 这里有异画的可能，所以需要先获取商品中的 card_version_id，同时获取到 product_id(商品ID)
			// 使用 card_version_id 从本地数据库中获取卡牌价格
			cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(product.CardVersionID))
			if err != nil {
				logrus.Errorf("获取 %v 价格失败：%v", product.CardNameCn, err)
			}

			// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
			// TODO: 这里可以写更新策略
			switch policy := updateFlags.UpdatePolicy; {
			case policy.Less3:
				// 如果集换价小于3元，则以集换价更新
				if cardPrice.AvgPrice < 3 {
					updateRun(&product, cardPrice, 0)
					logrus.Infof("商品【%v】%v 的价格更新为 %v", cardPrice.AlternativeArt, product.CardNameCn, cardPrice.AvgPrice)
				}
			case policy.Greater3AndLess10:
				// 如果集换价大于3元小于10元，则增加 0.5 元
				if cardPrice.AvgPrice >= 3 && cardPrice.AvgPrice <= 10 {
					updateRun(&product, cardPrice, 0.5)
					logrus.Infof("商品【%v】%v 的价格增加了0.5 块", cardPrice.AlternativeArt, product.CardNameCn)
				}
			case policy.Greater5AndLess50:
				// 如果集换价大于5块小于50块，且不是异画，则增加5元
				if cardPrice.AvgPrice > 5 && cardPrice.AvgPrice < 50 && cardPrice.AlternativeArt == "否" {
					updateRun(&product, cardPrice, 5)
					logrus.Infof("商品【%v】%v 的价格增加了5 块", cardPrice.AlternativeArt, product.CardNameCn)
				}

				// 如果集换价大于5块小于50块，且是异画，则增加10块
				if cardPrice.AvgPrice > 5 && cardPrice.AvgPrice < 50 && cardPrice.AlternativeArt == "是" {
					updateRun(&product, cardPrice, 10)
					logrus.Infof("商品【%v】%v 的价格增加了10 块", cardPrice.AlternativeArt, product.CardNameCn)
				}
			case policy.Greater50:
				if cardPrice.AvgPrice > 50 {
					updateRun(&product, cardPrice, 10)
					logrus.Infof("商品【%v】%v 的价格增加了10 块", cardPrice.AlternativeArt, product.CardNameCn)
				}
			default:
				// 更新商品，比集换价高 5 元
				// 防止误操作，默认不要更新
				// updateRun(&product, cardPrice, 5)
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
