package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	CardLevel string
	Color     string
	EffectKey string
}

func AddFlsgs(f *Flags) {
	pflag.StringVarP(&f.CardLevel, "level", "l", "", "Lv.2、Lv.3...Lv.7")
	pflag.StringVarP(&f.Color, "color", "c", "", "颜色，可用的值有：红、绿、蓝、黄、紫、黑、混色")
	pflag.StringVarP(&f.EffectKey, "effectKey", "k", "6000", "要查找带有该关键字效果的卡牌")
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

	// cardGroups := []string{"STC-01", "STC-02", "STC-03", "STC-04", "STC-05", "STC-06", "BTC-01", "BTC-02"}
	cardGroups := []string{"STC-01"}

	for _, cardGroup := range cardGroups {
		file := path.Join("./cards", cardGroup+".json")
		//打开文件
		fileByte, _ := os.ReadFile(file)

		for k, v := range string(fileByte) {
			fmt.Println(k, v)
		}

		var cardsDesc []models.CardDesc
		err := json.Unmarshal(fileByte, &cardsDesc)
		if err != nil {
			logrus.Fatalln(err)
		}

		for _, cardDesc := range cardsDesc {
			EffectKey(cardDesc, flags, cardGroup)
			if cardDesc.CardLevel == "Lv.3" && cardDesc.ParallCard == "1" {
				fmt.Println(cardDesc.Dp)
			}
		}
	}
}
