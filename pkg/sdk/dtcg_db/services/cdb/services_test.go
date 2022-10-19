package cdb

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb/models"
	"github.com/sirupsen/logrus"
)

func TestSearchClient_CardDeckSearch(t *testing.T) {
	client := NewCdbClient(core.NewClient(""))
	got, err := client.PostCardSearch(50)
	if err != nil {
		logrus.Fatalln(err)
	}

	fmt.Println(got)
}

// 获取卡包信息，并以 JSON 格式写入到文件中
func TestSearchClient_GetSeries(t *testing.T) {
	client := NewCdbClient(core.NewClient(""))
	series, err := client.GetSeries()
	if err != nil {
		logrus.Fatalln(err)
	}

	var cardGroups []models.SeriesPack

	for _, serie := range series.Data {
		for _, pack := range serie.SeriesPack {
			if pack.Language == "chs" {
				logrus.WithFields(logrus.Fields{
					"前缀": pack.PackPrefix,
					"名称": pack.PackName,
					"ID": pack.PackID,
				}).Infof("%v 中的卡包信息", serie.SeriesName)

				cardGroups = append(cardGroups, pack)
			}
		}
	}

	jsonByte, _ := json.Marshal(cardGroups)

	// 将响应信息写入文件
	fileName := "/mnt/d/Projects/DesistDaydream/dtcg/cards/dtcg_db/card_package.json"
	os.WriteFile(fileName, jsonByte, 0666)
}

// 获取卡片信息，并以 JSON 格式写入到文件中
func TestSearchClient_PostCardSearch(t *testing.T) {
	client := NewCdbClient(core.NewClient(""))

	cardPacks := make(map[int]string)

	series, err := client.GetSeries()
	if err != nil {
		logrus.Fatalln(err)
	}
	for _, serie := range series.Data {
		for _, pack := range serie.SeriesPack {
			if pack.Language == "chs" {
				cardPacks[pack.PackID] = pack.PackPrefix
			}
		}
	}

	// for id, _ := range cardPacks {
	// 	cards, err := client.PostCardSearch(id)
	// 	if err != nil {
	// 		logrus.Errorln(err)
	// 	}

	// 	for _, card := range cards.Data.List {
	// 		if card.Serial == "ST1-05" {
	// 			fmt.Println(card.CardID)
	// 		}
	// 	}
	// }

	for id, name := range cardPacks {
		cards, err := client.PostCardSearch(id)
		if err != nil {
			logrus.Errorln(err)
		}

		jsonByte, _ := json.Marshal(cards.Data)

		// 将响应信息写入文件
		fileName := fmt.Sprintf("/mnt/d/Projects/DesistDaydream/dtcg/cards/dtcg_db/%v.json", name)
		os.WriteFile(fileName, jsonByte, 0666)

		logrus.Infof("%v 中有 %v 种卡", name, cards.Data.Count)
	}
}

func TestSearchClient_GetCardPrice(t *testing.T) {
	client := NewCdbClient(core.NewClient(""))

	got, err := client.GetCardPrice("1896")
	if err != nil {
		logrus.Fatalln(err)
	}

	fmt.Println(got)
}
