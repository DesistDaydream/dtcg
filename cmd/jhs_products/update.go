package main

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

// 从 excel 中读取需要卖的卡数据，更新我在卖的商品
func UpdateProducts() {
	client := products.NewProductsClient(core.NewClient(""))
	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/我在卖.xlsx"
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
	sheet := "Sheet1"
	rows, err := f.GetRows(sheet)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"sheet": sheet,
		}).Fatalf("读取中sheet页异常: %v", err)
	}

	// for i := 1; i < len(rows); i++ {
	for i := 1; i < 3; i++ {
		resp, err := client.Update(&models.ProductsUpdateReqBody{
			Condition:            "1",
			OnSale:               "1",
			Price:                rows[i][1],
			Quantity:             rows[i][2],
			Remark:               "",
			UserCardVersionImage: rows[i][13],
		}, rows[i][0])
		if err != nil {
			logrus.Errorf("商品 %v %v 修改失败：%v", rows[i][10], rows[i][8], err)
		} else {
			logrus.Infof("商品 %v %v 修改成功：%v", rows[i][10], rows[i][8], resp)
		}
	}
}
