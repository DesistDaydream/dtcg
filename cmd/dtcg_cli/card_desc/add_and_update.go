package carddesc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func AddCardDescCommand() *cobra.Command {
	AddCardSetCmd := &cobra.Command{
		Use:   "add",
		Short: "添加卡片描述",
		Run:   addAndUpdateCardDesc,
	}

	return AddCardSetCmd
}

func UpdateCardDescCommand() *cobra.Command {
	UpdateCardDescCmd := &cobra.Command{
		Use:   "update",
		Short: "更新卡牌描述",
		Run:   addAndUpdateCardDesc,
	}

	return UpdateCardDescCmd
}

func addAndUpdateCardDesc(cmd *cobra.Command, args []string) {
	wantCardSets := getWantAddCardDesc()

	for _, cardSet := range wantCardSets {
		logrus.WithField("卡包名称", cardSet.SetPrefix).Infof("待下载卡包名称")
	}
	fmt.Printf("需要下载上述卡包，是否继续？(y/n) ")
	var confirm string
	fmt.Scanln(&confirm)
	if confirm != "y" {
		logrus.Infof("取消下载")
		return
	}

	switch cmd.Use {
	case "add":
		add(wantCardSets)
	case "update":
		update(wantCardSets)
	}
}

func add(wantCardSets []models.CardSet) {
	for _, cardSet := range wantCardSets {
		resp, err := handler.H.MoecardServices.Cdb.PostCardSearch(cardSet.SetID, "300", "chs", "")
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
			var image string
			if len(l.Images) > 0 {
				image = fmt.Sprintf("https://dtcg-pics.moecard.cn/img/%s~thumb.jpg", l.Images[0].ImgPath)
			} else {
				logrus.Errorf("无法获取 %v %v 卡图", l.CardID, l.ScName)
			}

			database.AddCardDesc(&models.CardDesc{
				CardIDFromDB:   l.CardID,
				SetID:          l.CardPack,
				SetName:        cardSet.SetName,
				SetPrefix:      cardSet.SetPrefix,
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
				Image:          image,
			})
		}
	}
}

func update(wantCardSets []models.CardSet) {
	for _, cardSet := range wantCardSets {
		resp, err := handler.H.MoecardServices.Cdb.PostCardSearch(cardSet.SetID, "300", "chs", "")
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
			var image string
			if len(l.Images) > 0 {
				image = fmt.Sprintf("https://dtcg-pics.moecard.cn/img/%s~thumb.jpg", l.Images[0].ImgPath)
			} else {
				logrus.Errorf("无法获取 %v %v 卡图", l.CardID, l.ScName)
			}

			database.UpdateCardDesc(&models.CardDesc{
				CardIDFromDB:   l.CardID,
				SetID:          l.CardPack,
				SetName:        cardSet.SetName,
				SetPrefix:      cardSet.SetPrefix,
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
				Image:          image,
			}, map[string]string{})
		}
	}
}
