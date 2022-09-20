package main

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/DesistDaydream/dtcg/cards"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
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

func AddFlsgs(f *Flags) {
	pflag.StringVarP(&f.CardLevel, "level", "l", "", "Lv.2、Lv.3...Lv.7")
	pflag.StringVarP(&f.Color, "color", "c", "", "颜色，可用的值有：红、绿、蓝、黄、紫、黑、混色")
	pflag.StringVarP(&f.EffectKey, "effectKey", "k", "6000", "要查找带有该关键字效果的卡牌")
	pflag.StringSliceVarP(&f.CardGroups, "cardGroups", "g", []string{"STC-01"}, "卡盒列表")
	pflag.BoolVarP(&f.Test, "test", "t", false, "是否进行测试。")

}

// 从 进化源效果、安防效果、卡牌效果 中查找根据自己定义的关键字过滤卡片
func EffectKey(cardDesc models.CardDesc, flags Flags, cardGroup string) {
	if strings.Contains(cardDesc.Effect, flags.EffectKey) || strings.Contains(cardDesc.SafeEffect, flags.EffectKey) || strings.Contains(cardDesc.EnvolutionEffect, flags.EffectKey) {
		logrus.WithFields(logrus.Fields{
			"卡包":    cardGroup,
			"名称":    cardDesc.Name,
			"颜色":    cardDesc.Color,
			"效果":    cardDesc.Effect,
			"安防效果":  cardDesc.SafeEffect,
			"进化源效果": cardDesc.EnvolutionEffect,
		}).Infoln(flags.EffectKey)
	}
}

func main() {
	var flags Flags
	AddFlsgs(&flags)
	pflag.Parse()

	if !flags.Test {
		var err error
		if flags.CardGroups, err = cards.GetCardGroups(); err != nil {
			logrus.Fatalf("获取卡盒列表失败：%v", err)
		}
	}

	for _, cardGroup := range flags.CardGroups {
		file := path.Join("./cards", cardGroup+".json")
		//打开文件
		fileByte, _ := os.ReadFile(file)

		var cardsDesc []models.CardDesc
		err := json.Unmarshal(fileByte, &cardsDesc)
		if err != nil {
			logrus.Fatalln(err)
		}

		for _, cardDesc := range cardsDesc {
			if cardDesc.ParallCard == "1" {
				EffectKey(cardDesc, flags, cardGroup)
			}
		}
	}
}
