package main

import (
	"os"

	"github.com/DesistDaydream/dtcg/cmd/excel/fileparse"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func checkFile(rrFile string) {
	if _, err := os.Stat(rrFile); os.IsNotExist(err) {
		logrus.Fatalf("【%v】文件不存在，请使用 -f 指定域名的记录规则文件", rrFile)
	}
}

func statistics(cardGroup string, cardDescs *models.CardDesc) {
	var (
		原画  int
		sec int
		sr  int
	)

	for _, cardDesc := range cardDescs.Page.List {
		if cardDesc.ParallCard == "1" {
			原画++
			if cardDesc.RareDegree == "隐藏稀有（SEC）" {
				sec++
			}
			if cardDesc.RareDegree == "超稀有（SR）" {
				sr++
			}
		}
	}

	logrus.WithFields(logrus.Fields{
		"数量":  len(cardDescs.Page.List),
		"原画":  原画,
		"异画":  len(cardDescs.Page.List) - 原画,
		"SEC": sec,
		"SR":  sr,
	}).Infof("【%v】卡包统计", cardGroup)
}

func main() {
	file := pflag.StringP("file", "f", "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/实卡统计.xlsx", "指定文件")
	cardGroup := pflag.StringP("cardGroup", "c", "BTC-02", "卡包")
	pflag.Parse()

	checkFile(*file)

	c := &models.FilterConditionReq{
		Page:             "",
		Limit:            "400",
		Name:             "",
		State:            "0",
		CardGroup:        *cardGroup,
		RareDegree:       "",
		BelongsType:      "",
		CardLevel:        "",
		Form:             "",
		Attribute:        "",
		Type:             "",
		Color:            "",
		EnvolutionEffect: "",
		SafeEffect:       "",
		ParallCard:       "",
		KeyEffect:        "",
	}

	// 根据过滤条件获取卡片详情
	cardDescs, err := services.GetCardDescs(c)
	if err != nil {
		panic(err)
	}

	// 统计卡包信息
	statistics(*cardGroup, cardDescs)

	fileparse.WriteExcelData(*file, cardDescs, *cardGroup)
}
