package fileparse

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func WriteExcelData(file string, cardDescs *models.CardDesc, cardGroup string) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		logrus.Errorln(err)
	}

	// 检查 sheet 是否存在
	// if !f.SheetExist(cardGroup) {
	// Create a new sheet.
	f.NewSheet(cardGroup)
	// }

	streamWriter, err := f.NewStreamWriter(cardGroup)
	if err != nil {
		logrus.Errorln(err)
	}

	err = streamWriter.SetRow("A1", []interface{}{"名称", "编号", "卡包", "稀有度", "颜色", "DP", "异画", "图片"})
	if err != nil {
		logrus.Errorln(err)
	}

	for i, cardDesc := range cardDescs.Page.List {
		var parallCard string
		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		if cardDesc.ParallCard == "1" {
			parallCard = "否"
		} else {
			parallCard = "是"
		}
		// Write some data to the stream writer.
		err := streamWriter.SetRow(cell, []interface{}{
			cardDesc.Name,
			cardDesc.Model,
			cardDesc.CardGroup,
			cardDesc.RareDegree,
			cardDesc.Color,
			cardDesc.Dp,
			parallCard,
		})
		if err != nil {
			logrus.Errorln(err)
		}

		resp, err := http.Get(cardDesc.ImageCover)
		if err != nil {
			logrus.Errorln(err)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logrus.Errorln(err)
		}
		format := `{
			"x_scale": 0.1,
			"y_scale": 0.1,
			"autofit": true,
			"positioning": "oneCell"
		}`
		f.AddPictureFromBytes(cardGroup, cell, format, "Excel Logo", ".png", bodyBytes)
	}

	if err := streamWriter.Flush(); err != nil {
		logrus.Errorln(err)
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}
