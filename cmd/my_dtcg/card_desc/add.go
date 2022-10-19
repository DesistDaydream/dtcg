package carddesc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func AddCardDescCommand() *cobra.Command {
	AddCardSetCmd := &cobra.Command{
		Use:   "add",
		Short: "添加卡片描述",
		Run:   addCardDesc,
	}

	return AddCardSetCmd
}

func addCardDesc(cmd *cobra.Command, args []string) {
	wantCardSets := getWantAddCardDesc()

	for _, cardSet := range wantCardSets {
		logrus.WithField("卡包名称", cardSet.PackPrefix).Infof("待下载卡包名称")
	}
	fmt.Printf("需要下载上述卡包，是否继续？(y/n) ")
	var confirm string
	fmt.Scanln(&confirm)
	if confirm != "y" {
		logrus.Infof("取消下载")
		return
	}

	for _, cardSet := range wantCardSets {
		client := cdb.NewSearchClient(core.NewClient(""))
		resp, err := client.PostCardSearch(cardSet.PackID)
		if err != nil {
			logrus.Fatalln(err)
		}

		for _, l := range resp.Data.List {
			var alternativeArt string
			if strings.Contains(l.Rarity, "-") {
				alternativeArt = "是"
			} else {
				alternativeArt = "否"
			}
			color, _ := json.Marshal(l.Color)
			class, _ := json.Marshal(l.Class)

			database.AddCardDesc(&models.CardDesc{
				CardIDFromDB:   l.CardID,
				SetID:          l.CardPack,
				SetName:        cardSet.PackName,
				SetPrefix:      cardSet.PackPrefix,
				Serial:         l.Serial,
				SubSerial:      l.SubSerial,
				JapName:        l.JapName,
				ScName:         l.ScName,
				AlternativeArt: alternativeArt,
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
			})
		}
	}
}

func getWantAddCardDesc() []models.CardSet {
	var wantCardSets []models.CardSet

	allCardSets, err := database.ListCardSets()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, cardSet := range allCardSets.Data {
		logrus.WithFields(logrus.Fields{
			"名称": cardSet.PackPrefix,
		}).Infof("卡包信息")
	}

	fmt.Printf("请选择需要下载图片的卡包，多个卡包用逗号分隔(使用 all 下载所有): ")

	var userInput string

	fmt.Scanln(&userInput)

	userInputCardSets := strings.Split(userInput, ",")
	for _, name := range userInputCardSets {
		for _, cardInfo := range allCardSets.Data {
			if cardInfo.PackPrefix == name {
				wantCardSets = append(wantCardSets, cardInfo)
			}
		}
	}

	switch userInput {
	case "all":
		return allCardSets.Data
	default:
		return wantCardSets
	}
}
