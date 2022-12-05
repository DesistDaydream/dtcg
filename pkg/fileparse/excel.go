package fileparse

import (
	"fmt"
	"reflect"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

type ExcelDataForPrice struct {
	Rows []models.JihuansheCardDescForPrice `json:"rows"`
}

func NewExcelDataForPrice(file string, test bool) (*ExcelDataForPrice, error) {
	var (
		d        ExcelDataForPrice
		cardSets []string
		err      error
	)

	if test {
		cardSets = []string{"STC-01"}
	} else {
		// 从数据库中获取卡牌集合信息
		cs, err := database.ListCardSets()
		if err != nil {
			return nil, fmt.Errorf("获取卡盒列表失败：%v", err)
		}

		for _, cardSet := range cs.Data {
			cardSets = append(cardSets, cardSet.SetPrefix)
		}
	}

	f, err := excelize.OpenFile(file)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			logrus.Errorln(err)
			return
		}
	}()

	for _, sheet := range cardSets {
		// 逐行读取Excel文件
		rows, err := f.GetRows(sheet)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"file":  file,
				"sheet": sheet,
			}).Errorf("读取中sheet页异常: %v", err)
			return nil, err
		}

		for r := 1; r < len(rows); r++ {
			var data models.JihuansheCardDescForPrice
			rType := reflect.TypeOf(&data).Elem()
			rVal := reflect.ValueOf(&data).Elem()

			// 通过遍历 struct 将 Excel 中的数据写入到 struct 的值中
			for i := 0; i < rType.NumField(); i++ {
				t := rType.Field(i)
				f := rVal.Field(i)
				v := rows[r][i]

				// 检查是否需要类型转换
				dataType := reflect.TypeOf(rows[r][i])
				structType := f.Type()
				if structType == dataType {
					f.Set(reflect.ValueOf(v))
				} else {
					if dataType.ConvertibleTo(structType) {
						// 转换类型
						f.Set(reflect.ValueOf(v).Convert(structType))
					} else {
						logrus.Errorf("%v 类型不匹配", t.Name)
					}
				}
			}

			if data.ParallCard == "1" {
				data.ParallCard = "原画"
			} else {
				data.ParallCard = "异画"
			}

			d.Rows = append(d.Rows, data)
		}
	}

	return &d, nil
}
