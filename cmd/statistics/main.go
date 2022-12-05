package main

import (
	"log"
	"sort"
	"strconv"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// 统计每个卡盒中各种稀有度的数量以及异画数量
func RareDegreeStatistics(cardSets *models.CardSets) {
	for _, cardSet := range cardSets.Data {
		var (
			原画  int
			sec int
			sr  int
		)
		cardsDesc, err := database.GetCardDescByCondition(300, 1, &models.QueryCardDesc{
			CardSet: int64(cardSet.SetID),
		})
		if err != nil {
			log.Fatalln(err)
		}
		if err != nil {
			logrus.Fatalln(err)
		}

		for _, cardDesc := range cardsDesc.Data {
			if cardDesc.AlternativeArt == "否" {
				原画++
				if cardDesc.RaritySC == "sec" {
					sec++
				}
				if cardDesc.RaritySC == "sr" {
					sr++
				}
			}
		}

		logrus.WithFields(logrus.Fields{
			"数量":  len(cardsDesc.Data),
			"原画":  原画,
			"异画":  len(cardsDesc.Data) - 原画,
			"SEC": sec,
			"SR":  sr,
		}).Infof("【%v】卡包统计", cardSet.SetName)
	}
}

// 统计每个级别的各种 DP 的数量
func DPStatistics(cardGroups *models.CardSets) {
	// 获取卡牌等级列表
	cardLevels, err := database.GetCardDescLevel()
	if err != nil {
		logrus.Fatalln(err)
	}
	// 获取卡牌详情列表
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Debugf("当前共有 %v 张卡", cardsDesc.Count)

	// 测试用，只给一个等级
	// cardLevels = []string{"3"}

	for _, cardLevel := range cardLevels {
		// 遍历所有卡片描述，获取指定等级的所有 DP 保存在该数组中
		var alldps []int
		for _, cardDesc := range cardsDesc.Data {
			if cardDesc.AlternativeArt == "否" {
				if cardDesc.Level == cardLevel {
					dp, _ := strconv.Atoi(cardDesc.DP)
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

		for i := 0; i < len(alldps); i++ {
			if newalldps[element] != alldps[i] {
				logrus.Infof("【%v】 级别中有 %v 个 %v DP 的数码宝贝", cardLevel, count, alldps[i-1])
				count = 1
				// logrus.Infof("发现新一个 DP：%v", alldps[i])
				newalldps = append(newalldps, alldps[i])
				element++
			} else {
				count++
			}
		}
		// 输出最后一波发现到的数据
		logrus.Infof("【%v】 级别中有 %v 个 %v DP 的数码宝贝", cardLevel, count, alldps[len(alldps)-1])

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
	pflag.StringVarP(&f.StatisticsInfo, "statistics", "s", "rare-degree", "统计信息的类型，可选值：rare-degree, dp")
	pflag.BoolVarP(&f.DownloadImg, "downloadImg", "d", false, "是否下载图片")
}

// 从已经下载的文件中统计一些卡片信息
func main() {
	var flags Flags
	AddFlsgs(&flags)
	pflag.Parse()

	// 初始化配置文件
	c := config.NewConfig("", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	cardSets, err := database.ListCardSets()
	if err != nil {
		logrus.Fatalln(err)
	}

	switch flags.StatisticsInfo {
	case "rare-degree":
		RareDegreeStatistics(cardSets)
	case "dp":
		DPStatistics(cardSets)
	default:
		logrus.Fatalln("指定要统计的信息")
	}
}
