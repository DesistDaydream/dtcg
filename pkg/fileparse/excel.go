package fileparse

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

type ExcelData struct {
	Rows []models.JihuansheExporterCardDesc `json:"rows"`
}

func NewExcelData(file string, sheets []string) (*ExcelData, error) {
	var d ExcelData
	// var ed *models.JihuansheExporterCardDesc
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

	for _, sheet := range sheets {
		// 逐行读取Excel文件
		rows, err := f.GetRows(sheet)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"file":  file,
				"sheet": sheet,
			}).Errorf("读取中sheet页异常: %v", err)
			return nil, err
		}

		for i := 1; i < len(rows); i++ {
			var data models.JihuansheExporterCardDesc
			data.Model = rows[i][2]
			data.Name = rows[i][9]
			data.CardVersionID = rows[i][25]

			d.Rows = append(d.Rows, data)
		}
	}

	return &d, nil
}
