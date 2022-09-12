package fileparse

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func downloadImg(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func WriteExcelData(file string, cardDescs *models.CardListResponse, cardGroupSheet string, isTruedownloadImg bool) {
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

	var colNames []string = []string{"名称", "编号", "卡包", "稀有度", "颜色", "DP", "异画", "图片"}

	// 写入第一行数据，设置列名
	for i, colName := range colNames {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(cardGroupSheet, cell, colName)

	}

	// 从第二行开始写入数据
	for i, cardDesc := range cardDescs.Page.CardsDesc {
		// 设置行高。三个参数分别为：Sheet 名，行号，高度
		err = f.SetRowHeight(cardGroupSheet, i+2, 45)
		if err != nil {
			panic(err)
		}

		var parallCard string

		if cardDesc.ParallCard == "1" {
			parallCard = "否"
		} else {
			parallCard = "是"
		}

		values := []string{cardDesc.Name,
			cardDesc.Model,
			cardDesc.CardGroup,
			cardDesc.RareDegree,
			cardDesc.Color,
			cardDesc.Dp,
			parallCard,
		}

		// 写入数据
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2)
			f.SetCellValue(cardGroupSheet, cell, v)
		}

		if isTruedownloadImg {
			// 获取图片
			imgBytes, err := downloadImg(cardDesc.ImageCover)
			if err != nil {
				logrus.Errorf("获取图片异常：%v", err)
			}
			// 插入图片
			format := `{
			"autofit": true,
			"lock_aspect_ratio": true
		}`
			f.AddPictureFromBytes(cardGroupSheet, fmt.Sprintf("%v%v", string('A'+len(values)), i+2), format, "Excel Logo", ".png", imgBytes)
		}
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}
