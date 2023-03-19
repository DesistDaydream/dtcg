package products

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/internal/database"
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateImageFlags struct {
	noImage bool
}

var updateImageFlags UpdateImageFlags

func UpdateImageCommand() *cobra.Command {
	updateProductsImageCmd := &cobra.Command{
		Use:   "image",
		Short: "更新商品图片",
		Run:   updateImage,
	}

	updateProductsImageCmd.Flags().BoolVar(&updateImageFlags.noImage, "no-img", false, "是否全量更新没有卡图的商品")

	return updateProductsImageCmd
}

// 更新商品卡图
func updateImage(cmd *cobra.Command, args []string) {

	if updateImageFlags.noImage {
		updateNoImage()
	} else {
		// 生成待处理的卡牌信息
		cards, err := GenNeedHandleCards(updatePriceFlags.UpdatePolicy.PriceRange, updatePriceFlags.UpdatePolicy.isArt)
		if err != nil {
			logrus.Errorf("%v", err)
			return
		}
		logrus.Infof("%v 价格区间中共有 %v 张卡牌需要更新", updatePriceFlags.UpdatePolicy.PriceRange, len(cards.Data))

		// 根据更新策略更新卡牌价格
		genNeedHandleImgProducts(cards, updatePriceFlags.UpdatePolicy.PriceChange)
	}
}

// 生成待处理的商品信息
func genNeedHandleImgProducts(cards *dbmodels.CardsPrice, priceChange float64) {
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
		for _, product := range products.Products {
			logrus.WithFields(logrus.Fields{
				"原始价格": cardPrice.AvgPrice,
				"更新价格": cardPrice.AvgPrice + priceChange,
				"调整价格": priceChange,
			}).Infof("更新前检查【%v】【%v %v】商品", card.AlternativeArt, card.Serial, product.CardNameCn)
			// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
			if productsFlags.isRealRun {
				resp, err := handler.H.JhsServices.Products.Update(&models.ProductsUpdateReqBody{
					ProductCardVersionImage: cardPrice.ImageUrl,
				}, fmt.Sprint(product.ProductID))
				if err != nil {
					logrus.Errorf("商品 %v %v 修改失败：%v", product.ProductID, product.CardNameCn, err)
				} else {
					logrus.Infof("商品 %v %v 修改成功：%v", product.ProductID, product.CardNameCn, resp)
				}
			}
		}
	}
}

func updateNoImage() {
	page := 1 // 从获取到的数据的第一页开始
	for {
		products, err := handler.H.JhsServices.Products.List(strconv.Itoa(page), "", updateFlags.CurSaleState)
		if err != nil || len(products.Data) <= 0 {
			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
		}

		for _, product := range products.Data {
			if !strings.Contains(product.CardVersionImage, "cdn-client") {
				cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(product.CardVersionID))
				if err != nil {
					logrus.Errorf("获取卡牌 %v 价格详情失败: %v", product.CardNameCn, err)
					continue
				}

				resp, err := handler.H.JhsServices.Products.Update(&models.ProductsUpdateReqBody{
					Condition:               fmt.Sprint(product.Condition),
					OnSale:                  fmt.Sprint(product.OnSale),
					Price:                   product.Price,
					Quantity:                fmt.Sprint(product.Quantity),
					Remark:                  product.Remark,
					ProductCardVersionImage: cardPrice.ImageUrl,
				}, fmt.Sprint(product.ProductID))
				if err != nil {
					logrus.Errorf("商品 %v %v 修改失败：%v", product.ProductID, product.CardNameCn, err)
				} else {
					logrus.Infof("商品 %v %v 修改成功：%v", product.ProductID, product.CardNameCn, resp)
				}
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
