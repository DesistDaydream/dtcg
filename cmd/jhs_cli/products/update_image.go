package products

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DesistDaydream/dtcg/internal/database"
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
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
		cards, err := GenNeedHandleCards()
		if err != nil {
			logrus.Errorf("%v", err)
			return
		}

		// 根据更新策略更新卡牌价格
		genNeedHandleImgProducts(cards)
	}
}

// 生成待处理的商品信息
func genNeedHandleImgProducts(cards *dbmodels.CardsPrice) {
	for _, card := range cards.Data {
		// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到该卡牌的商品列表
		products, err := handler.H.JhsServices.Market.SellersProductsList("1", card.Serial, updateFlags.CurSaleState, "published_at_desc")
		if err != nil {
			logrus.Fatal(err)
		}
		cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(card.CardVersionID))
		if err != nil {
			logrus.Errorf("获取 %v 价格失败：%v", card.ScName, err)
		}
		for _, product := range products.Data {
			logrus.Infof("更新前检查【%v】【%v %v】商品", card.AlternativeArt, card.Serial, product.CardNameCN)
			// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
			if productsFlags.isRealRun {
				resp, err := handler.H.JhsServices.Market.SellersProductsUpdate(&models.ProductsUpdateReqBody{
					Condition:               fmt.Sprint(product.Condition),
					OnSale:                  fmt.Sprint(product.OnSale),
					Price:                   product.Price,
					Quantity:                fmt.Sprint(product.Quantity),
					Remark:                  product.Remark,
					ProductCardVersionImage: cardPrice.ImageUrl,
				}, fmt.Sprint(product.ProductID))
				if err != nil {
					logrus.Errorf("商品 %v %v 修改失败：%v", product.ProductID, product.CardNameCN, err)
				} else {
					logrus.Infof("商品 %v %v 修改成功：%v", product.ProductID, product.CardNameCN, resp)
				}
			}
		}
	}
}

func updateNoImage() {
	page := 1 // 从获取到的数据的第一页开始
	for {
		products, err := handler.H.JhsServices.Market.SellersProductsList(strconv.Itoa(page), "", updateFlags.CurSaleState, "published_at_desc")
		if err != nil || len(products.Data) <= 0 {
			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
		}

		for _, product := range products.Data {
			if !strings.Contains(product.CardVersionImage, "cdn-client") {
				cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(product.CardVersionID))
				if err != nil {
					logrus.Errorf("获取卡牌 %v 价格详情失败: %v", product.CardNameCN, err)
					continue
				}

				resp, err := handler.H.JhsServices.Market.SellersProductsUpdate(&models.ProductsUpdateReqBody{
					Condition:               fmt.Sprint(product.Condition),
					OnSale:                  fmt.Sprint(product.OnSale),
					Price:                   product.Price,
					Quantity:                fmt.Sprint(product.Quantity),
					Remark:                  product.Remark,
					ProductCardVersionImage: cardPrice.ImageUrl,
				}, fmt.Sprint(product.ProductID))
				if err != nil {
					logrus.Errorf("商品 %v %v 修改失败：%v", product.ProductID, product.CardNameCN, err)
				} else {
					logrus.Infof("商品 %v %v 修改成功：%v", product.ProductID, product.CardNameCN, resp)
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
