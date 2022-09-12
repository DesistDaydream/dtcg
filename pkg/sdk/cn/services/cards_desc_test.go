package services

import (
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
)

func TestGetCardsDesc(t *testing.T) {
	// cardGroups := []string{"STC-01", "STC-02", "STC-03", "STC-04", "STC-05", "STC-06", "BTC-01", "BTC-02"}
	cardGroups := []string{"BTC-02"}

	for _, cardGroup := range cardGroups {
		filterConditionReq := &models.FilterConditionReq{
			Page:             "1",
			Limit:            "3",
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

		resp, err := GetCardDescs(filterConditionReq)
		if err != nil {
			logrus.Fatal(err)
		}

		for _, c := range resp.Page.CardsDesc {

			logrus.WithFields(logrus.Fields{
				"ID":    c.ID,
				"卡名":    c.Name,
				"是否显示":  c.State,
				"所属卡组":  c.CardGroup,
				"卡片编号":  c.Model,
				"稀有度":   c.RareDegree,
				"关键字效果": c.KeyEffect,
			}).Infoln("卡片详情")
		}

		// jsonByte, _ := json.Marshal(resp.Page.CardsDesc)
		// fileName := fmt.Sprintf("../../../../cards/%v.json", cardGroup)

		// os.WriteFile(fileName, jsonByte, 0666)
	}
}
