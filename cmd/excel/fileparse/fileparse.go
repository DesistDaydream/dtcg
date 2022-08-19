package fileparse

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func WriteExcelData(file string, cardDescs *models.CardDesc, cardGroup string) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet(cardGroup)

	// Set value of a cell.
	f.SetCellValue(cardGroup, "A1", "名称")
	f.SetCellValue(cardGroup, "B1", "编号")
	f.SetCellValue(cardGroup, "C1", "卡包")
	f.SetCellValue(cardGroup, "D1", "稀有度")
	f.SetCellValue(cardGroup, "E1", "颜色")
	f.SetCellValue(cardGroup, "F1", "DP")
	f.SetCellValue(cardGroup, "G1", "异画")
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	for index, cardDesc := range cardDescs.Page.List {
		f.SetCellValue(cardGroup, "A"+fmt.Sprint(index+2), cardDesc.Name)
		f.SetCellValue(cardGroup, "B"+fmt.Sprint(index+2), cardDesc.Model)
		f.SetCellValue(cardGroup, "C"+fmt.Sprint(index+2), cardDesc.CardGroup)
		f.SetCellValue(cardGroup, "D"+fmt.Sprint(index+2), cardDesc.RareDegree)
		f.SetCellValue(cardGroup, "E"+fmt.Sprint(index+2), cardDesc.Color)
		f.SetCellValue(cardGroup, "F"+fmt.Sprint(index+2), cardDesc.Dp)
		if cardDesc.ParallCard == "1" {
			f.SetCellValue(cardGroup, "G"+fmt.Sprint(index+2), "否")
		} else {
			f.SetCellValue(cardGroup, "G"+fmt.Sprint(index+2), "是")
		}
	}

	// Save xlsx file by the given path.
	if err := f.SaveAs(file); err != nil {
		logrus.Error(err)
	}
}
