package products

import (
	"fmt"
	"strconv"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type DelFlags struct {
	SaleState string
}

var delFlags DelFlags

func DelCommand() *cobra.Command {
	long := `
根据策略添加商品。
比如：
  jhs_cli products add -s BTC-03 -r 0,1000 -c 20 表示将所有价格在 0-1000 之间卡牌的价格增加 20 块售卖。
`
	delProdcutCmd := &cobra.Command{
		Use:   "del",
		Short: "删除商品",
		Long:  long,
		Run:   delProducts,
	}

	delProdcutCmd.Flags().StringVarP(&delFlags.SaleState, "sale-state", "s", "0", "商品的售卖状态。删除指定状态的商品")

	return delProdcutCmd
}

// 添加商品
func delProducts(cmd *cobra.Command, args []string) {
	page := 1 // 从获取到的数据的第一页开始
	for {
		products, err := handler.H.JhsServices.Products.List(strconv.Itoa(page), "", delFlags.SaleState)
		if err != nil || len(products.Data) <= 0 {
			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
		}

		for _, product := range products.Data {
			resp, err := handler.H.JhsServices.Products.Del(fmt.Sprint(product.ProductID))
			if err != nil {
				logrus.Errorf("商品 %v %v 删除失败：%v", product.ProductID, product.CardNameCn, err)
			} else {
				logrus.Infof("商品 %v %v 删除成功：%v", product.ProductID, product.CardNameCn, resp)
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
