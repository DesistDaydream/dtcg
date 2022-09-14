package main

import (
	"strings"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
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

func main() {
	var flags Flags
	AddFlsgs(&flags)
	pflag.Parse()

	cardGroups := []string{"STC-01", "STC-02", "STC-03", "STC-04", "STC-05", "STC-06", "BTC-01", "BTC-02"}

	for _, cardGroup := range cardGroups {
		c := &models.FilterConditionReq{
			Page:             "",
			Limit:            "300",
			Name:             "",
			State:            "0",
			CardGroup:        cardGroup,
			RareDegree:       "",
			BelongsType:      "",
			CardLevel:        flags.CardLevel,
			Form:             "",
			Attribute:        "",
			Type:             "",
			Color:            flags.Color,
			EnvolutionEffect: "",
			SafeEffect:       "",
			ParallCard:       "1",
			KeyEffect:        "",
		}
		// 根据过滤条件获取卡片详情
		cardDescs, err := services.GetCardsDesc(c)
		if err != nil {
			panic(err)
		}

		for _, cardDesc := range cardDescs.Page.CardsDesc {
			logrus.WithFields(logrus.Fields{
				"名称":   cardDesc.Name,
				"效果":   cardDesc.Effect,
				"安防效果": cardDesc.SafeEffect,
				"进化效果": cardDesc.EnvolutionEffect,
			}).Debugln("检查")
			// 判断 cardDesc.Effect 中包含字符串 6000
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
	}
}
