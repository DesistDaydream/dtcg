package fileparse

import (
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewExcelData(t *testing.T) {
	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/价格统计表.xlsx"
	test := true
	got, err := NewExcelDataForPrice(file, test)
	if err != nil {
		logrus.Errorln(err)
	}

	var count int64 = 0

	for _, data := range got.Rows {
		f, _ := strconv.ParseFloat(data.AvgPrice, 64)
		if f > 1 {
			count++
			logrus.Infof("%v %v", data.Name, data.ParallCard)
		}
	}
	logrus.Infof("在 %v 张卡中，集换价大于 1 块的有 %v 张", len(got.Rows), count)
}
