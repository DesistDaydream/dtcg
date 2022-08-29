package main

import (
	"strings"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	var effectKey string
	pflag.StringVarP(&effectKey, "effectKey", "k", "6000", "要查找带有该关键字效果的卡牌")
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
			CardLevel:        "",
			Form:             "",
			Attribute:        "",
			Type:             "",
			Color:            "",
			EnvolutionEffect: "",
			SafeEffect:       "",
			ParallCard:       "1",
			KeyEffect:        "",
		}
		// 根据过滤条件获取卡片详情
		cardDescs, err := services.GetCardDescs(c)
		if err != nil {
			panic(err)
		}

		for _, cardDesc := range cardDescs.Page.List {
			logrus.WithFields(logrus.Fields{
				"名称":   cardDesc.Name,
				"效果":   cardDesc.Effect,
				"安防效果": cardDesc.SafeEffect,
				"进化效果": cardDesc.EnvolutionEffect,
			}).Debugln("检查")
			// 判断 cardDesc.Effect 中包含字符串 6000
			if strings.Contains(cardDesc.Effect, effectKey) || strings.Contains(cardDesc.SafeEffect, effectKey) || strings.Contains(cardDesc.EnvolutionEffect, effectKey) {
				logrus.WithFields(logrus.Fields{
					"卡包": cardGroup,
					"名称": cardDesc.Name,
					"效果": cardDesc.Effect,
				}).Infoln(effectKey)
			}
		}
	}
}
