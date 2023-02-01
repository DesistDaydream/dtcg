package services

import (
	"encoding/json"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/copier"
)

// 处理卡片描述中的 KeyEffect，因为响应体是字符串，我们可以将其改为数组
func NewCardsDesc(resp *models.CardListResponse) ([]byte, error) {
	var newCardsDesc []*models.NewCardDesc

	for _, c := range resp.Page.CardsDesc {
		// 首先，将原始结构体中的值拷贝到新结构体中，并根据注释忽略某些字段
		var newCardDesc models.NewCardDesc
		copier.Copy(&newCardDesc, &c)

		// 转换 KeyEffect 字段的值为数组
		slice := []string{}
		json.Unmarshal([]byte(c.KeyEffect), &slice)

		// 将新的 KeyEffect 的值添加到新的结构体中
		newCardDesc.KeyEffect = append(newCardDesc.KeyEffect, slice...)

		// logrus.WithFields(logrus.Fields{
		// 	"ID":    newCardDesc.ID,
		// 	"卡名":    newCardDesc.Name,
		// 	"是否显示":  newCardDesc.State,
		// 	"所属卡组":  newCardDesc.CardGroup,
		// 	"卡片编号":  newCardDesc.Model,
		// 	"稀有度":   newCardDesc.RareDegree,
		// 	"关键字效果": newCardDesc.KeyEffect,
		// }).Infoln("卡片详情")

		newCardsDesc = append(newCardsDesc, &newCardDesc)
	}

	jsonByte, _ := json.Marshal(newCardsDesc)

	return jsonByte, nil

}

func NewReq() *models.FilterConditionReq {
	return &models.FilterConditionReq{
		Page:             "1",
		Limit:            "300",
		Name:             "",
		State:            "0",
		CardGroup:        "",
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
}

// 获取卡牌描述
func TestDownloadCardsDesc(t *testing.T) {
	cardGroups := []string{"BTC-04"}

	for _, cardPackage := range cardGroups {
		filterConditionReq := NewReq()
		filterConditionReq.Limit = "300"
		filterConditionReq.CardGroup = cardPackage
		filterConditionReq.ParallCard = "0"

		// 获取卡片描述
		resp, err := GetCardsDesc(filterConditionReq)
		if err != nil {
			logrus.Fatal(err)
		}

		for _, cardDesc := range resp.Page.CardsDesc {
			logrus.WithFields(logrus.Fields{
				"编号": cardDesc.Model,
			}).Infof("%v 卡牌信息", cardDesc.Name)
		}

		logrus.Infof("共获取到 %v 张卡牌", len(resp.Page.CardsDesc))
	}
}
