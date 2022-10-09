package fileparse

import (
	"strconv"
	"testing"

	"github.com/DesistDaydream/dtcg/cards"
	"github.com/sirupsen/logrus"
)

func TestNewExcelData(t *testing.T) {
	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/价格统计表.xlsx"

	cardGroups, err := cards.GetCardGroups("/mnt/d/Projects/DesistDaydream/dtcg/cards/card_package.json")
	if err != nil {
		logrus.Fatalln(err)
	}

	// cardGroups = []string{"STC-01"}

	got, err := NewExcelData(file, cardGroups)
	if err != nil {
		logrus.Errorln(err)
	}

	var count int64 = 0

	for _, data := range got.Rows {
		f, _ := strconv.ParseFloat(data.AvgPrice, 64)
		if f > 1 {
			count++
		}
	}
	logrus.Infof("在 %v 张卡中，集换价大于1块的有 %v 张", len(got.Rows), count)
}
