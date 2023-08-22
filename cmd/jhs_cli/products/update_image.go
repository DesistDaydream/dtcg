package products

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/handler"
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
		// genNeedHandleImgProducts(cards)
		ps := genNeedUpdateProducts(cards)
		logrus.Infof("共匹配到 %v 件商品", ps.count)

		// 逐一更新商品图片
		for _, p := range ps.products {
			u1, _ := url.Parse(p.img)
			u2, _ := url.Parse(p.card.ImageUrl)
			// 只有当前商品图与数据库中的图不一样时，才更新
			if u1.Path != u2.Path {
				logrus.WithFields(logrus.Fields{}).Infof("检查将要更新的商品: 【%v】【%v】【%v】【%v %v】", p.productID, p.product.CardVersionID, p.card.AlternativeArt, p.card.Serial, p.product.CardNameCN)

				p.img = p.card.ImageUrl

				if productsFlags.isRealRun {
					updateRun(&p)
				}
			}
		}
	}
}

func updateNoImage() {
	page := 1 // 从获取到的数据的第一页开始
	for {
		products, err := handler.H.JhsServices.Market.SellersProductsList(page, "", updateFlags.CurSaleState, "published_at_desc")
		if err != nil || len(products.Data) <= 0 {
			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
		}

		for _, p := range products.Data {
			if !strings.Contains(p.CardVersionImage, "cdn-client") {
				cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(p.CardVersionID))
				if err != nil {
					logrus.Errorf("获取卡牌 %v 价格详情失败: %v", p.CardNameCN, err)
					continue
				}

				updateRun(&Product{
					card:           cardPrice,
					product:        &p,
					defaultProduct: nil,
					productID:      p.ProductID,
					condition:      p.Condition,
					cardVersionID:  p.CardVersionID,
					cardNameCN:     p.CardNameCN,
					img:            cardPrice.ImageUrl,
					onSale:         p.OnSale,
					price:          p.Price,
					quantity:       p.Quantity,
				})
			}
		}

		fmt.Printf("\r\033[2K 共 %v 页数据，已处理第 %v 页", products.LastPage, products.CurrentPage)

		// 如果当前处理的页等于最后页，则退出循环
		if products.CurrentPage == products.LastPage {
			logrus.Debugf("退出循环时共 %v 页,处理完 %v 页", products.LastPage, products.CurrentPage)
			break
		}

		// 每处理完一页，下一个循环需要处理的页+1
		page++
	}
}
