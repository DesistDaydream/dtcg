package fileparse

import (
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
	index := f.NewSheet(cardGroup)
	// }

	streamWriter, err := f.NewStreamWriter(cardGroup)
	if err != nil {
		panic(err)
	}

	err = streamWriter.SetRow("A1", []interface{}{"名称", "编号", "卡包", "稀有度", "颜色", "DP", "异画"})
	if err != nil {
		logrus.Error(err)
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
			panic(err)
		}
	}

	if err := streamWriter.Flush(); err != nil {
		logrus.Error(err)
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save xlsx file by the given path.
	// if err := f.SaveAs(file); err != nil {
	// 	logrus.Error(err)
	// }

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}

}
