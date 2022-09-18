package main

import (
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

// 统计每个级别的各种 DP 的数量
func DPStatistics(cardGroups []string) {
	// 获取卡牌等级列表
	cardLevels, err := cards.GetCardLevel()
	if err != nil {
		logrus.Fatalln(err)
	}
	// 获取卡牌详情列表
	allCardsDesc, err := cards.MergeCardsDesc(cardGroups)
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Debugf("当前共有 %v 张卡", len(allCardsDesc))

	// 测试用，只给一个等级
	// cardLevels = []string{"Lv.3"}

	for _, cardLevel := range cardLevels {
		// 遍历所有卡片描述，获取指定等级的所有 DP 保存在该数组中
		var alldps []int
		for _, cardDesc := range allCardsDesc {
			if cardDesc.ParallCard == "1" {
				if cardDesc.CardLevel == cardLevel {
					dp, _ := strconv.Atoi(cardDesc.Dp)
					alldps = append(alldps, dp)
				}
			}
		}

		// 将 DP 排序
		sort.Ints(alldps[:])

		// logrus.Infof(alldps, len(alldps))

		// 排序之后，需要使用类似 Linux 中 uniq 命令的行为，来进行统计。
		// 存放元素不同的数组
		var newalldps []int
		// 不同元素的个数
		var element int
		// DP 计数
		var count int

		newalldps = append(newalldps, alldps[0])

		// 长度 +1 是为了当新 dp 只有 1 个的时候，可以再多来一次循环并输出信息
		for i := 0; i < len(alldps)+1; i++ {
			if i < len(alldps) {
				if newalldps[element] != alldps[i] {
					logrus.Infof("【%v】 级别中有 %v 个 %v DP 的数码宝贝", cardLevel, count, alldps[i-1])
					count = 1
					// logrus.Infof("发现新一个 DP：%v", alldps[i])
					newalldps = append(newalldps, alldps[i])
					element++
				} else {
					count++
				}
			} else { // 若最后新 DP 只有 1 个，通过这里输出一下信息
				logrus.Infof("【%v】 级别中有 %v 个 %v DP 的数码宝贝", cardLevel, count, alldps[i-1])
			}

		}

		// for i, v := range alldps {
		// 	if len(newalldps) == 0 {
		// 		newalldps = append(newalldps, v)
		// 		// logrus.Infof("发现第一个 DP：%v", v)
		// 	}

		// 	if newalldps[element] != v {
		// 		logrus.Infof("【%v】 级别中有 %v 个 %v DP 的数码宝贝", cardLevel, count, alldps[i-1])
		// 		count = 1
		// 		// logrus.Infof("发现新一个 DP：%v", v)
		// 		newalldps = append(newalldps, v)
		// 		element++
		// 	} else {
		// 		count++
		// 	}
		// }

		logrus.Infof("此级别的数码宝贝共有 【%v】 种 DP", len(newalldps))
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
