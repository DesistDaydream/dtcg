package main

import (
	"github.com/DesistDaydream/dtcg/cards"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func RareDegreeStatistics(cardGroup string, cardDescs []models.CardDesc) {
	var (
		原画  int
		sec int
		sr  int
	)

	for _, cardDesc := range cardDescs {
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
		"数量":  len(cardDescs),
		"原画":  原画,
		"异画":  len(cardDescs) - 原画,
		"SEC": sec,
		"SR":  sr,
	}).Infof("【%v】卡包统计", cardGroup)
}

type Flags struct {
	StatisticsInfo string
	DownloadImg    bool
}

func AddFlsgs(f *Flags) {
	pflag.StringVarP(&f.StatisticsInfo, "statistics", "s", "rare-degree", "指定文件")
	pflag.BoolVarP(&f.DownloadImg, "downloadImg", "d", false, "是否下载图片")
}

// 从已经下载的文件中统计一些卡片信息
func main() {
	var flags Flags
	AddFlsgs(&flags)
	pflag.Parse()

	cardGroups, err := cards.GetCardGroups()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, cardGroup := range cardGroups {
		cardDescs, err := cards.GetCardDesc(cardGroup)
		if err != nil {
			logrus.Fatalln(err)
		}

		switch flags.StatisticsInfo {
		case "rare-degree":
			// 统计卡盒稀有度信息
			RareDegreeStatistics(cardGroup, cardDescs)
		default:
			logrus.Fatalln("指定要统计的信息")
		}

	}
}
