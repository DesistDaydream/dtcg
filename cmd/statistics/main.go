package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/DesistDaydream/dtcg/cards"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// 统计每个卡盒中各种稀有度的数量以及异画数量
func RareDegreeStatistics(cardGroups []string) {
	for _, cardGroup := range cardGroups {
		var (
			原画  int
			sec int
			sr  int
		)

		cardsDesc, err := cards.GetCardDesc(cardGroup)
		if err != nil {
			logrus.Fatalln(err)
		}

		for _, cardDesc := range cardsDesc {
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
			"数量":  len(cardsDesc),
			"原画":  原画,
			"异画":  len(cardsDesc) - 原画,
			"SEC": sec,
			"SR":  sr,
		}).Infof("【%v】卡包统计", cardGroup)
	}
}

func DPStatistics(cardGroups []string) {
	var alldps []int
	for _, cardGroup := range cardGroups {
		cardsDesc, err := cards.GetCardDesc(cardGroup)
		if err != nil {
			logrus.Fatalln(err)
		}

		var dps []int

		for _, cardDesc := range cardsDesc {
			if cardDesc.ParallCard == "1" {
				if cardDesc.CardLevel == "Lv.3" {
					dp, _ := strconv.Atoi(cardDesc.Dp)
					dps = append(dps, dp)
				}
			}
		}

		alldps = append(alldps, dps...)
	}

	sort.Ints(alldps[:])
	// 统计 alldps 中的重复元素的个数

	// fmt.Println(alldps)
	for _, v := range alldps {
		fmt.Println(v)
	}
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

	switch flags.StatisticsInfo {
	case "rare-degree":
		RareDegreeStatistics(cardGroups)
	case "dp":
		DPStatistics(cardGroups)
	default:
		logrus.Fatalln("指定要统计的信息")

	}
}
