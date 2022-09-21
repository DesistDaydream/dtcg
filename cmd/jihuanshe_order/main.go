package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/DesistDaydream/dtcg/cmd/jihuanshe_order/fileparse"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services"
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

func GetBuyerOrderList(page string, token string) ([]int64, error) {
	buyerOrders, err := services.GetBuyerOrders(page, token)
	if err != nil {
		return nil, err
	}

	var buyerOrderList []int64

	for _, buyerOrder := range buyerOrders.Data {
		buyerOrderList = append(buyerOrderList, int64(buyerOrder.OrderID))
	}

	// 如果查询到的记录条数大于 pageSize 的值，那么需要分页查询。并将查询到的记录合并
	if buyerOrders.LastPage > 1 {
		for i := 2; i <= buyerOrders.LastPage; i++ {
			buyerOrders, err := services.GetBuyerOrders(strconv.Itoa(i), token)
			if err != nil {
				return nil, err
			}

			for _, buyerOrder := range buyerOrders.Data {
				buyerOrderList = append(buyerOrderList, int64(buyerOrder.OrderID))
			}
		}
	}

	return buyerOrderList, nil
}

func GetSellerOrderList(page string, token string) ([]int64, error) {
	sellerOrders, err := services.GetSellerOrders(page, token)
	if err != nil {
		return nil, err
	}

	var sellerOrderList []int64

	for _, sellerOrder := range sellerOrders.Data {
		sellerOrderList = append(sellerOrderList, int64(sellerOrder.OrderID))
	}

	// 如果查询到的记录条数大于 pageSize 的值，那么需要分页查询。并将查询到的记录合并
	if sellerOrders.LastPage > 1 {
		for i := 2; i <= sellerOrders.LastPage; i++ {
			sellerOrders, err := services.GetSellerOrders(strconv.Itoa(i), token)
			if err != nil {
				return nil, err
			}

			for _, sellerOrder := range sellerOrders.Data {
				sellerOrderList = append(sellerOrderList, int64(sellerOrder.OrderID))
			}
		}
	}

	return sellerOrderList, nil
}

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

	buyerOrderList, err := GetBuyerOrderList("1", flags.Token)
	if err != nil {
		logrus.Error(err)
	}

	sellerOrderList, err := GetSellerOrderList("1", flags.Token)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Debugln("买入订单号", buyerOrderList, len(buyerOrderList))
	logrus.Debugln("卖出订单号", sellerOrderList, len(sellerOrderList))

	fileparse.FileParse(flags.File, buyerOrderList, flags.Token, "买入")
	fileparse.FileParse(flags.File, sellerOrderList, flags.Token, "卖出")
}
