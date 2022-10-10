package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
	"github.com/sirupsen/logrus"
)

func addCard() {
	filePath := "cards/card_package.json"
	file, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Fatalln(err)
	}

	var cardGroups *models.CacheListResp

	err = json.Unmarshal(file, &cardGroups)
	if err != nil {
		logrus.Fatalln(err)
	}

	c := &models.FilterConditionReq{
		Limit: "3",
		State: "0",
	}

	for _, cardGroup := range cardGroups.List {
		// 若要获取卡盒所有卡，需要将限制扩大
		c.Limit = "300"
		c.CardGroup = cardGroup.Name

		// 根据过滤条件获取卡片详情
		cardDescs, err := services.GetCardsDesc(c)
		if err != nil {
			panic(err)
		}

		for _, cardDesc := range cardDescs.Page.CardsDesc {
			d := &database.CardDesc{
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
			database.AddCard(d)
		}
	}
}

func addCardGroup(wirteToJSON bool) {
	cardPackageResp, err := services.GetCardGroups()
	if err != nil {
		logrus.Fatalln(err)
	}

	sort.Slice(cardPackageResp.List, func(i, j int) bool {
		return cardPackageResp.List[i].CreateTime < cardPackageResp.List[j].CreateTime
	})

	if wirteToJSON {
		jsonByte, _ := json.Marshal(cardPackageResp)

		fileName := filepath.Join("cards", "card_package.json")
		os.WriteFile(fileName, jsonByte, 0666)
	}

	for _, cardGroup := range cardPackageResp.List {
		g := &database.CardGroup{
			Name:       cardGroup.Name,
			Image:      cardGroup.Image,
			State:      cardGroup.State,
			Position:   cardGroup.Position,
			CreateTime: cardGroup.CreateTime,
			UpdateTime: cardGroup.UpdateTime,
		}
		database.AddCardGroup(g)
	}
}

func main() {
	database.InitDB()
	// addCard()
	addCardGroup(false)
}
