package products

import (
	"fmt"
	"strconv"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateSaleStateFlags struct {
	OldSaleState string
	NewSaleState string
}

var updateSaleStateFlags UpdateSaleStateFlags

func UpdateSaleStateCommand() *cobra.Command {
	UpdateProductsSaleStateCmd := &cobra.Command{
		Use:   "sale-state",
		Short: "更新商品的售卖状态，是上架还是下架",
		Run:   updateSaleState,
	}

	UpdateProductsSaleStateCmd.Flags().StringVarP(&updateSaleStateFlags.OldSaleState, "old-sale-state", "o", "0", "当前售卖状态。即获取什么状态的商品")
	UpdateProductsSaleStateCmd.Flags().StringVarP(&updateSaleStateFlags.NewSaleState, "new-sale-state", "n", "1", "期望的售卖状态。")

	return UpdateProductsSaleStateCmd
}

// 更新我在卖卡牌的卡图
func updateSaleState(cmd *cobra.Command, args []string) {
	page := 1 // 从获取到的数据的第一页开始
	for {
		products, err := handler.H.JhsServices.Products.List(strconv.Itoa(page), "", updateSaleStateFlags.OldSaleState)
		if err != nil || len(products.Data) <= 0 {
			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
		}

		for _, product := range products.Data {
			if product.Quantity != 0 {
				resp, err := handler.H.JhsServices.Products.Update(&models.ProductsUpdateReqBody{
					AuthenticatorID:         "",
					Grading:                 "",
					Condition:               fmt.Sprint(product.Condition),
					OnSale:                  updateSaleStateFlags.NewSaleState,
					Price:                   product.Price,
					ProductCardVersionImage: product.CardVersionImage,
					Quantity:                fmt.Sprint(product.Quantity),
					Remark:                  product.Remark,
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
