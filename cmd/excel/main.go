package main

import (
	"os"

	"github.com/DesistDaydream/dtcg/cmd/excel/fileparse"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/sirupsen/logrus"
)

func checkFile(rrFile string) {
	if _, err := os.Stat(rrFile); os.IsNotExist(err) {
		logrus.Fatalf("【%v】文件不存在，请使用 -f 指定域名的记录规则文件", rrFile)
	}
}

func main() {
	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/实卡统计.xlsx"
	// checkFile(file)
	cardGroup := "BTC-02"

	c := &models.FilterConditionReq{
		Page:             "",
		Limit:            "400",
		Name:             "",
		State:            "0",
		CardGroup:        cardGroup,
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

	fileparse.WriteExcelData(file, cardDescs, cardGroup)

}
