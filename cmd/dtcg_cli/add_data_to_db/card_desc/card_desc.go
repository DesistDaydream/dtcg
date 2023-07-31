package carddesc

import (
	"encoding/json"
	"log"

	"github.com/DesistDaydream/dtcg/pkg/database"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"

	"github.com/sirupsen/logrus"
)

func AddCardDescFromOfficial() {
	cardGroups, err := database.ListCardGroupsFromOfficial()
	if err != nil {
		log.Fatalln(err)
	}

	c := &models.FilterConditionReq{
		Limit: "3",
		State: "0",
	}

	for _, cardGroup := range cardGroups.Data {
		// 若要获取卡盒所有卡，需要将限制扩大
		c.Limit = "300"
		c.CardGroup = cardGroup.Name

		// 根据过滤条件获取卡片详情
		cardDescs, err := services.GetCardsDesc(c)
		if err != nil {
			panic(err)
		}

		for _, cardDesc := range cardDescs.Page.CardsDesc {
			d := &database.CardDescFromOfficial{
				OfficialID:           cardDesc.ID,
				CardGroup:            cardDesc.CardGroup,
				Model:                cardDesc.Model,
				RareDegree:           cardDesc.RareDegree,
				BelongsType:          cardDesc.BelongsType,
				CardLevel:            cardDesc.CardLevel,
				Color:                cardDesc.Color,
				Form:                 cardDesc.Form,
				Attribute:            cardDesc.Attribute,
				Name:                 cardDesc.Name,
				Dp:                   cardDesc.Dp,
				Type:                 cardDesc.Type,
				EntryConsumeValue:    cardDesc.EntryConsumeValue,
				EnvolutionConsumeOne: cardDesc.EnvolutionConsumeOne,
				EnvolutionConsumeTwo: cardDesc.EnvolutionConsumeTwo,
				GetWay:               cardDesc.GetWay,
				Effect:               cardDesc.Effect,
				SafeEffect:           cardDesc.SafeEffect,
				EnvolutionEffect:     cardDesc.EnvolutionEffect,
				ImageCover:           cardDesc.ImageCover,
				State:                cardDesc.State,
				ParallCard:           cardDesc.ParallCard,
				KeyEffect:            cardDesc.KeyEffect,
			}
			database.AddCardDescOfficial(d)
		}
	}
}

func AddCardDescFromDtcgDB() {
	d, err := database.ListCardGroupsFromDtcgDB()
	if err != nil {
		logrus.Fatalln(err)
	}
	for _, set := range d.Data {
		resp, err := handler.H.MoecardServices.Cdb.PostCardSearch(set.PackID, "300", "chs", "")
		if err != nil {
			logrus.Fatalln(err)
		}

		for _, l := range resp.Data.List {
			color, _ := json.Marshal(l.Color)
			class, _ := json.Marshal(l.Class)
			d := &database.CardDescFromDtcgDB{
				CardID:         l.CardID,
				CardPack:       l.CardPack,
				PackName:       set.PackName,
				PackPrefix:     set.PackPrefix,
				Serial:         l.Serial,
				SubSerial:      l.SubSerial,
				JapName:        l.JapName,
				ScName:         l.ScName,
				Rarity:         l.Rarity,
				Type:           l.Type,
				Color:          string(color),
				Level:          l.Level,
				Cost:           l.Cost,
				Cost1:          l.Cost1,
				EvoCond:        l.EvoCond,
				DP:             l.DP,
				Grade:          l.Grade,
				Attribute:      l.Attribute,
				Class:          string(class),
				Illustrator:    l.Illustrator,
				Effect:         l.Effect,
				EvoCoverEffect: l.EvoCoverEffect,
				SecurityEffect: l.SecurityEffect,
				IncludeInfo:    l.IncludeInfo,
				RaritySC:       l.RaritySC,
			}

			database.AddCardDescFromDtcgDB(d)
		}
	}
}
