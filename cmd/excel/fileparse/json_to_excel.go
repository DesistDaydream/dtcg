package fileparse

import (
	"fmt"
	"reflect"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

// 从官网获取到的 JSON 格式的卡片详情写入到 Excel 中
func JsonToExcel(file string, cardDescs *models.CardListResponse, cardGroupSheet string, colNames []string) {
	opts := excelize.Options{}
	f, err := excelize.OpenFile(file, opts)
	if err != nil {
		logrus.Errorln(err)
	}

	// 检查 sheet 是否存在
	// if !f.SheetExist(cardGroup) {
	// Create a new sheet.
	f.NewSheet(cardGroupSheet)
	// }

	// 设置列宽。四个参数分别为：Sheet 名，起始列号，结束列号，宽度
	err = f.SetColWidth(cardGroupSheet, "A", "H", 5.57)
	if err != nil {
		panic(err)
	}

	// 写入第一行数据，设置列名
	for i, colName := range colNames {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(cardGroupSheet, cell, colName)

	}

	// 从第二行开始写入数据
	for i, cardDesc := range cardDescs.Page.CardsDesc {
		// 通过反射获取结构体中的每一个值
		var values []string
		v := reflect.ValueOf(cardDesc)
		for k := 0; k < v.NumField(); k++ {
			v := fmt.Sprintf("%v", v.Field(k).Interface())
			values = append(values, v)
		}

		// 写入数据
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
			f.SetCellValue(cardGroupSheet, cell, v)
		}
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}
