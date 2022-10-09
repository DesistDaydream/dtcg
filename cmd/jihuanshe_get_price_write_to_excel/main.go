package main

import (
	"github.com/DesistDaydream/dtcg/cards"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

// 获取每张卡的价格写入到 Excel 中
func main() {
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

	cardGroups, err := cards.GetCardGroups("")
	if err != nil {
		logrus.Fatalln(err)
	}

	// cardGroups = []string{"BTC-02", "STC-06", "STC-05", "STC-04", "BTC-01", "PR卡", "STC-03", "STC-02", "STC-01"}

	for _, sheet := range cardGroups {
		rows, err := f.GetRows(sheet)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"file":  file,
				"sheet": sheet,
			}).Fatalf("读取中sheet页异常: %v", err)
		}

		f.SetCellValue(sheet, "AB1", "最低价")
		f.SetCellValue(sheet, "AC1", "集换价")

		for i := 1; i < len(rows); i++ {
			cardVersionID := rows[i][25]
			productsGetResponse, err := client.Get(cardVersionID)
			if err != nil {
				logrus.Fatal(err)
			}

			cellMin, _ := excelize.CoordinatesToCellName(28, i+1)
			err = f.SetCellValue(sheet, cellMin, productsGetResponse.MinPrice)
			if err != nil {
				logrus.Errorln(err)
			}
			cellAvg, _ := excelize.CoordinatesToCellName(29, i+1)
			err = f.SetCellValue(sheet, cellAvg, productsGetResponse.AvgPrice)
			if err != nil {
				logrus.Errorln(err)
			}
		}
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}
