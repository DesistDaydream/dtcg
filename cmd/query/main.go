package main

import (
	"log"
	"strings"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	CardLevel  string
	Color      string
	EffectKey  string
	CardGroups []string
	Test       bool
}

func AddFlags(f *Flags) {
	pflag.StringVarP(&f.CardLevel, "level", "l", "", "Lv.2、Lv.3...Lv.7")
	pflag.StringVarP(&f.Color, "color", "c", "", "颜色，可用的值有：红、绿、蓝、黄、紫、黑、混色")
	pflag.StringVarP(&f.EffectKey, "effectKey", "k", "6000", "要查找带有该关键字效果的卡牌")
	pflag.StringSliceVarP(&f.CardGroups, "cardGroups", "g", []string{"STC-01"}, "卡盒列表")
	pflag.BoolVarP(&f.Test, "test", "t", false, "是否进行测试。")

}

// 从 进化源效果、安防效果、卡牌效果 中查找根据自己定义的关键字过滤卡片
func EffectKey(cardDesc models.CardDesc, flags Flags, cardSet string) {
	if strings.Contains(cardDesc.Effect, flags.EffectKey) || strings.Contains(cardDesc.SecurityEffect, flags.EffectKey) || strings.Contains(cardDesc.EvoCoverEffect, flags.EffectKey) {
		logrus.WithFields(logrus.Fields{
			"卡包":    cardSet,
			"名称":    cardDesc.ScName,
			"颜色":    cardDesc.Color,
			"效果":    cardDesc.Effect,
			"安防效果":  cardDesc.SecurityEffect,
			"进化源效果": cardDesc.EvoCoverEffect,
		}).Infoln(flags.EffectKey)
	}
}

func main() {
	var (
		flags    Flags
		logFlags logging.LoggingFlags
	)
	AddFlags(&flags)
	logging.AddFlags(&logFlags)
	pflag.Parse()
	if err := logging.LogInit(&logFlags); err != nil {
		logrus.Fatalf("初始化日志失败: %v", err)
	}

	// 初始化配置文件
	c := config.NewConfig("", "")

	// 连接数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}
	database.InitDB(dbInfo)

	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		log.Fatalln(err)
	}

	logrus.Infof("共有 %v 张卡", cardsDesc.Count)

	for _, cardDesc := range cardsDesc.Data {
		if cardDesc.AlternativeArt == "否" {
			EffectKey(cardDesc, flags, cardDesc.SetPrefix)
		}
	}

}
