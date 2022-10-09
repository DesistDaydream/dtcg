package fileparse

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewExcelData(t *testing.T) {
	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/价格统计表.xlsx"
	sheet := "STC-01"
	got, err := NewExcelData(file, sheet)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(got)
}
