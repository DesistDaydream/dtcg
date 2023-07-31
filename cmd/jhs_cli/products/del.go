package products

import (
	"fmt"
	"strconv"

	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type DelFlags struct {
	SaleState  string
	productIDs []string
}

var delFlags DelFlags

func DelCommand() *cobra.Command {
	delProdcutCmd := &cobra.Command{
		Use:   "del",
		Short: "删除商品",
		Run:   delProducts,
	}

	delProdcutCmd.Flags().StringVar(&delFlags.SaleState, "sale-state", "", "商品的售卖状态。删除指定状态的商品，一般使用 0 ，即下架的商品")
	delProdcutCmd.Flags().StringSliceVar(&delFlags.productIDs, "ids", nil, "要删除的商品 ID 列表")

	return delProdcutCmd
}

// 删除商品
func delProducts(cmd *cobra.Command, args []string) {
	if delFlags.productIDs != nil {
		delProductsForIDs()
	} else if delFlags.SaleState != "" {
		delProdcutsForSaleState()
	} else {
		logrus.Fatalln("请指至少一种命令行标志以获取商品 ID 的信息")
	}
}

func delProductsForIDs() {
	for _, productID := range delFlags.productIDs {
		resp, err := handler.H.JhsServices.Market.SellersProductsDel(productID)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("%v 删除成功：%v", productID, resp)
	}
}

func delProdcutsForSaleState() {
	page := 1 // 从获取到的数据的第一页开始
	for {
		products, err := handler.H.JhsServices.Market.SellersProductsList(strconv.Itoa(page), "", delFlags.SaleState, "published_at_desc")
		if err != nil || len(products.Data) <= 0 {
			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
		}

		for _, product := range products.Data {
			resp, err := handler.H.JhsServices.Market.SellersProductsDel(fmt.Sprint(product.ProductID))
			if err != nil {
				logrus.Errorf("商品 %v %v 删除失败：%v", product.ProductID, product.CardNameCN, err)
			} else {
				logrus.Infof("商品 %v %v 删除成功：%v", product.ProductID, product.CardNameCN, resp)
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
