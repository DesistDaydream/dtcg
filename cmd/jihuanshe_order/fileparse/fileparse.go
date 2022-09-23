package fileparse

import (
	"fmt"
	"reflect"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/orders"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/orders/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func FileParse(client *orders.OrdersClient, file string, orders []int64, buyOrSell string) {
	opts := excelize.Options{}
	f, err := excelize.OpenFile(file, opts)
	if err != nil {
		logrus.Errorln(err)
	}

	f.NewSheet(buyOrSell)

	var colNames []string
	desc := &models.OrderProduct{}
	s := reflect.TypeOf(desc).Elem()
	for i := 0; i < s.NumField(); i++ {
		colNames = append(colNames, s.Field(i).Name)
	}

	// 写入第一行数据，设置列名
	for i, colName := range colNames {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(buyOrSell, cell, colName)

	}

	switch buyOrSell {
	case "买入":
		BuyerFileParse(client, f, orders, buyOrSell)
	case "卖出":
		SellerFileParse(client, f, orders, buyOrSell)
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}

func BuyerFileParse(client *orders.OrdersClient, f *excelize.File, allOrders []int64, buyOrSell string) {
	// 从第二行开始写入产品信息，所以 row = 2
	for i, row := 0, 2; i < len(allOrders); i++ {
		ops, err := client.GetBuyerOrderProducts(int(allOrders[i]))
		if err != nil {
			logrus.Errorln(err)
		}

		// 每个订单里包含多个产品，我们需要将每个产品的信息写到 Excel 中的一行
		for i, orderProduct := range ops.OrderProducts {
			var values []string
			v := reflect.ValueOf(orderProduct)
			for k := 0; k < v.NumField(); k++ {
				v := fmt.Sprintf("%v", v.Field(k).Interface())
				values = append(values, v)
			}

			// 写入数据
			for j, v := range values {
				// 将要写入的数据的行受上一次写入影响
				// 每个订单的产品写入完成后，下一个订单开始写入的行不能从头开始，所以要加上本次订单产品的数量
				// 所以需要最后的 row = row + len(ops.OrderProducts) 这条逻辑
				cell, _ := excelize.CoordinatesToCellName(j+1, i+row)
				f.SetCellValue(buyOrSell, cell, v)
			}
		}

		row = row + len(ops.OrderProducts)
	}
}

func SellerFileParse(client *orders.OrdersClient, f *excelize.File, allOrders []int64, buyOrSell string) {
	for i, row := 0, 2; i < len(allOrders); i++ {
		ops, err := client.GetSellerOrderProducts(int(allOrders[i]))
		if err != nil {
			logrus.Errorln(err)
		}

		for i, orderProduct := range ops.OrderProducts {
			var values []string
			v := reflect.ValueOf(orderProduct)
			for k := 0; k < v.NumField(); k++ {
				v := fmt.Sprintf("%v", v.Field(k).Interface())
				values = append(values, v)
			}

			for j, v := range values {
				cell, _ := excelize.CoordinatesToCellName(j+1, i+row)
				f.SetCellValue(buyOrSell, cell, v)
			}
		}

		row = row + len(ops.OrderProducts)
	}
}
