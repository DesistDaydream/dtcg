package orders

import (
	"strconv"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func GetAllOrderPriceCommand() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "获取所有买入以及卖出订单的总价格",
		Run:   getAllOrderPriceRun,
	}

	return getCmd
}

func getAllOrderPriceRun(cmd *cobra.Command, args []string) {
	var totalBuyerPrcie float64
	var totalSellerPrcie float64

	buyerPage := 1
	sellerPage := 1

	for {
		buyerOrders, err := handler.H.JhsServices.Market.OrderList(strconv.Itoa(buyerPage))
		if err != nil {
			logrus.Error(err)
		}

		for _, order := range buyerOrders.Data {
			totalBuyerPrcie = totalBuyerPrcie + order.TotalPrice
		}

		logrus.Infof("买入订单共 %v 页，已处理完第 %v 页", buyerOrders.LastPage, buyerOrders.CurrentPage)
		if buyerOrders.CurrentPage == buyerOrders.LastPage {
			logrus.Debugf("%v/%v 已处理完成，退出循环", buyerOrders.CurrentPage, buyerOrders.LastPage)
			break
		}

		buyerPage = buyerOrders.CurrentPage + 1
	}

	// 分页
	for {
		sellerOrders, err := handler.H.JhsServices.Market.SellerOrderList(strconv.Itoa(sellerPage))
		if err != nil {
			logrus.Error(err)
		}

		for _, order := range sellerOrders.Data {
			totalSellerPrcie = totalSellerPrcie + order.TotalPrice
		}

		logrus.Infof("卖出订单共 %v 页，已处理完第 %v 页", sellerOrders.LastPage, sellerOrders.CurrentPage)
		if sellerOrders.CurrentPage == sellerOrders.LastPage {
			logrus.Debugf("%v/%v 已处理完成，退出循环", sellerOrders.CurrentPage, sellerOrders.LastPage)
			break
		}

		sellerPage = sellerOrders.CurrentPage + 1
	}

	logrus.Infof("当前买入订单的总额为：%v", totalBuyerPrcie)
	logrus.Infof("当前卖出订单的总额为：%v", totalSellerPrcie)

}
