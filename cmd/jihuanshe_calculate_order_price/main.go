package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/xuri/excelize/v2"
)

// 根据 Excel 计算所有卡片的集换价和最低价
func main() {
	var (
		// flags    Flags
		logFlags logging.LoggingFlags
	)
	// AddFlsgs(&flags)
	logging.AddFlags(&logFlags)
	pflag.Parse()

	if err := logging.LogInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	cardModelToCardVersionID := make(map[string]string)
	filea, err := os.ReadFile("/mnt/d/Projects/DesistDaydream/dtcg/cards/card_version_id_and_card_modle.json")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = json.Unmarshal(filea, &cardModelToCardVersionID)
	if err != nil {
		logrus.Fatalln(err)
	}

	client := products.NewProductsClient(core.NewClient(""))
	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/价格统计表.xlsx"
	f, err := excelize.OpenFile(file)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			logrus.Errorln(err)
			return
		}
	}()
	sheet := "待计算"
	rows, err := f.GetRows(sheet)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"sheet": sheet,
		}).Fatalf("读取中sheet页异常: %v", err)
	}

	var AllAvgPrice []float64
	var AllMinPrice []float64

	for i := 1; i < len(rows); i++ {
		if cardModelToCardVersionID[rows[i][0]] == "" {
			logrus.Errorf("获取 %v 的集换社 ID 【%v】异常，可能为空", rows[i][0], cardModelToCardVersionID[rows[i][0]])
		}

		// 开始上架
		resp, err := client.Get(cardModelToCardVersionID[rows[i][0]])
		if err != nil {
			logrus.Errorf("获取商品价格失败：%v", err)
		}

		quantity, _ := strconv.ParseFloat(rows[i][1], 64)
		avgPrice, _ := strconv.ParseFloat(resp.AvgPrice, 64)
		minPrice, _ := strconv.ParseFloat(resp.MinPrice, 64)
		avgPrice = avgPrice * quantity
		minPrice = minPrice * quantity

		logrus.WithFields(logrus.Fields{
			"集换价":  resp.AvgPrice,
			"总集换价": avgPrice,
			"最低价":  resp.MinPrice,
			"总最低价": minPrice,
		}).Infof("%v 张 %v 的价格", rows[i][1], rows[i][0])

		AllAvgPrice = append(AllAvgPrice, avgPrice)
		AllMinPrice = append(AllMinPrice, minPrice)
	}

	var TotalAvgPrice float64
	var TotalMinPrice float64

	for _, avgPrice := range AllAvgPrice {
		TotalAvgPrice = TotalAvgPrice + avgPrice
	}

	for _, minPrice := range AllMinPrice {
		TotalMinPrice = TotalMinPrice + minPrice
	}

	logrus.WithFields(logrus.Fields{
		"集换价": TotalAvgPrice,
		"最低价": TotalMinPrice,
	}).Infof("所有卡牌价格")
}
