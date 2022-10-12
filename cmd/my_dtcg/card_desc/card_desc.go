package carddesc

import (
	"encoding/json"
	"strings"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	"github.com/sirupsen/logrus"
)

func AddCardDesc() {
	d, err := database.ListCardSets()
	if err != nil {
		logrus.Fatalln(err)
	}
	for _, set := range d.Data {
		client := services.NewSearchClient(core.NewClient(""))
		resp, err := client.PostCardSearch(set.PackID)
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
			d := &models.CardDesc{
				CardIDFromDB:   l.CardID,
				SetID:          l.CardPack,
				SetName:        set.PackName,
				SetPrefix:      set.PackPrefix,
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
			}

			database.AddCardDesc(d)
		}
	}
}
