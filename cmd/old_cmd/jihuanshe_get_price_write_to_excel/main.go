package main

import (
	"strconv"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

// 获取每张卡的价格写入到 Excel 中
func main() {
	// 初始化配置文件
	c := config.NewConfig("", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	client := products.NewProductsClient(core.NewClient(""))

	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/价格统计表-测试.xlsx"
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

	cardSets, err := database.ListCardSets()
	if err != nil {
		logrus.Fatalln(err)
	}

	// cardGroups = []string{"STC-01"}

	for _, dataSet := range cardSets.Data {
		rows, err := f.GetRows(dataSet.SetName)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"file":  file,
				"sheet": dataSet,
			}).Fatalf("读取中sheet页异常: %v", err)
		}

		f.SetCellValue(dataSet.SetName, "AA1", "最低价")
		f.SetCellValue(dataSet.SetName, "AB1", "集换价")

		for row := 1; row < len(rows); row++ {
			// cardVersionID := rows[row][25]
			cellCardVersionID, err := excelize.CoordinatesToCellName(26, row+1)
			if err != nil {
				logrus.Errorf("获取CardVersionID单元格名称错误：%v", err)
			}
			cardVersionID, err := f.GetCellValue(dataSet.SetName, cellCardVersionID)
			if err != nil {
				logrus.Errorf("获取CardVersionID单元格的值错误：%v", err)
			}

			// 获取卡牌加个信息
			productsGetResponse, err := client.Get(cardVersionID)
			if err != nil {
				logrus.Fatal(err)
			}
			minPrice, _ := strconv.ParseFloat(productsGetResponse.MinPrice, 64)
			avgPrice, _ := strconv.ParseFloat(productsGetResponse.AvgPrice, 64)

			// 向单元格写入数据
			cellMin, _ := excelize.CoordinatesToCellName(27, row+1)
			err = f.SetCellValue(dataSet.SetName, cellMin, minPrice)
			if err != nil {
				logrus.Errorf("设置最低价单元格的值错误：%v", err)
			}
			cellAvg, _ := excelize.CoordinatesToCellName(28, row+1)
			err = f.SetCellValue(dataSet.SetName, cellAvg, avgPrice)
			if err != nil {
				logrus.Errorf("设置集换价单元格的值错误：%v", err)
			}
		}
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}
