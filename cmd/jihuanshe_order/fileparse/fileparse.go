package fileparse

import (
	"fmt"
	"reflect"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func BuyerFileParse(file string, orders []int64, token string, sheet string) {
	opts := excelize.Options{}
	f, err := excelize.OpenFile(file, opts)
	if err != nil {
		logrus.Errorln(err)
	}

	f.NewSheet(sheet)

	var colNames []string
	desc := &models.BuyerOrderProduct{}
	s := reflect.TypeOf(desc).Elem()
	for i := 0; i < s.NumField(); i++ {
		colNames = append(colNames, s.Field(i).Name)
	}
	// 写入第一行数据，设置列名
	for i, colName := range colNames {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, colName)

	}

	var ops []models.BuyerOrderProduct

	for _, buyerOrder := range orders {
		orderProducts, err := services.GetBuyerOrderProducts(int(buyerOrder), token)
		if err != nil {
			logrus.Errorln(err)
		}
		ops = append(ops, orderProducts.OrderProducts...)
	}

	for i, orderProduct := range ops {
		// 通过反射获取结构体中的每一个值
		var values []string
		v := reflect.ValueOf(orderProduct)
		for k := 0; k < v.NumField(); k++ {
			v := fmt.Sprintf("%v", v.Field(k).Interface())
			values = append(values, v)
		}

		// 写入数据
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
			f.SetCellValue(sheet, cell, v)
		}
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}

func SellerFileParse(file string, orders []int64, token string, sheet string) {
	opts := excelize.Options{}
	f, err := excelize.OpenFile(file, opts)
	if err != nil {
		logrus.Errorln(err)
	}

	f.NewSheet(sheet)

	var colNames []string
	desc := &models.SellerOrderProduct{}
	s := reflect.TypeOf(desc).Elem()
	for i := 0; i < s.NumField(); i++ {
		colNames = append(colNames, s.Field(i).Name)
	}
	// 写入第一行数据，设置列名
	for i, colName := range colNames {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, colName)

	}

	var ops []models.SellerOrderProduct

	for _, buyerOrder := range orders {
		orderProducts, err := services.GetSellerOrderProducts(int(buyerOrder), token)
		if err != nil {
			logrus.Errorln(err)
		}
		ops = append(ops, orderProducts.OrderProducts...)
	}

	for i, orderProduct := range ops {
		// 通过反射获取结构体中的每一个值
		var values []string
		v := reflect.ValueOf(orderProduct)
		for k := 0; k < v.NumField(); k++ {
			v := fmt.Sprintf("%v", v.Field(k).Interface())
			values = append(values, v)
		}

		// 写入数据
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
			f.SetCellValue(sheet, cell, v)
		}
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}
