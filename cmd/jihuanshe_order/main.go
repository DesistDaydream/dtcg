package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/DesistDaydream/dtcg/cmd/jihuanshe_order/fileparse"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/orders"
)

func checkFile(rrFile string) {
	if _, err := os.Stat(rrFile); os.IsNotExist(err) {
		logrus.Fatalf("【%v】文件不存在，请使用 -f 指定文件", rrFile)
	}
}

type Flags struct {
	File  string
	Token string
}

func AddFlsgs(f *Flags) {
	pflag.StringVarP(&f.File, "file", "f", "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/订单.xlsx", "指定文件")
	pflag.StringVarP(&f.Token, "token", "t", "", "用户认证信息")
	// pflag.StringVarP(&f.File, "file", "f", "test.xlsx", "指定文件")
}

type Orders struct {
}

func GetBuyerOrderList(client *orders.OrdersClient) ([]int64, error) {
	var orders []int64

	page := 1
	for {
		buyerOrders, err := client.GetBuyerOrders(strconv.Itoa(page))
		if err != nil {
			return nil, err
		}

		for _, order := range buyerOrders.Data {
			orders = append(orders, int64(order.OrderID))
		}

		logrus.Infof("买入订单共 %v 页，已处理完第 %v 页", buyerOrders.LastPage, buyerOrders.CurrentPage)
		if buyerOrders.CurrentPage == buyerOrders.LastPage {
			logrus.Debugln("%v/%v 已处理完成，退出循环", buyerOrders.CurrentPage, buyerOrders.LastPage)
			break
		}

		page = buyerOrders.CurrentPage + 1
	}

	return orders, nil
}

func GetSellerOrderList(client *orders.OrdersClient) ([]int64, error) {
	var orders []int64

	page := 1
	// 分页
	for {
		sellerOrders, err := client.GetSellerOrders(strconv.Itoa(page))
		if err != nil {
			return nil, err
		}

		for _, order := range sellerOrders.Data {
			orders = append(orders, int64(order.OrderID))
		}

		logrus.Infof("卖出订单共 %v 页，已处理完第 %v 页", sellerOrders.LastPage, sellerOrders.CurrentPage)
		if sellerOrders.CurrentPage == sellerOrders.LastPage {
			logrus.Debugln("%v/%v 已处理完成，退出循环", sellerOrders.CurrentPage, sellerOrders.LastPage)
			break
		}

		page = sellerOrders.CurrentPage + 1
	}

	return orders, nil
}

// 获取买入和卖出订单中所有产品信息，写入到 Excel 中
func main() {
	var flags Flags
	AddFlsgs(&flags)
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()
	pflag.Parse()

	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	checkFile(flags.File)

	client := orders.NewOrdersClient(core.NewClient(flags.Token))

	buyerOrderList, err := GetBuyerOrderList(client)
	if err != nil {
		logrus.Error(err)
	}

	sellerOrderList, err := GetSellerOrderList(client)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Debugln("买入订单号", buyerOrderList, len(buyerOrderList))
	logrus.Debugln("卖出订单号", sellerOrderList, len(sellerOrderList))

	fileparse.FileParse(client, flags.File, buyerOrderList, "买入")
	fileparse.FileParse(client, flags.File, sellerOrderList, "卖出")
}
