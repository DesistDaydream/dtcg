package cdb

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/sirupsen/logrus"
)

var client *CdbClient

func initTest() {
	client = NewCdbClient(core.NewClient("", 10))
}

// 列出卡牌集合
func TestCdbClient_GetSeries(t *testing.T) {
	initTest()
	series, err := client.GetSeries()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, serie := range series.Data {
		for _, pack := range serie.SeriesPack {
			if pack.Language == "ja" {
				logrus.WithFields(logrus.Fields{
					"前缀": pack.PackPrefix,
					"名称": pack.PackName,
					"ID": pack.PackID,
				}).Infof("%v 中的卡包信息", serie.SeriesName)
			}
		}
	}
}

// 获取卡牌集合详情(包括卡集中包含的所有卡牌)
func TestCdbClient_GetPackage(t *testing.T) {
	initTest()
	cardSet, err := client.GetPackage("STC-10")
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, card := range cardSet.Data.Cards {
		logrus.WithFields(logrus.Fields{
			"ID": card.CardID,
			"名称": card.ScName,
		}).Infof("卡牌信息")
	}
}

func TestSearchClient_CardDeckSearch(t *testing.T) {
	initTest()
	cardsDesc, err := client.PostCardSearch(79, "10", "chs", "")
	// cardsDesc, err := client.PostCardSearch(27, "10", "ja", "true")
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Infof("共查询到 %v 张卡", cardsDesc.Data.Count)

	for _, cardDesc := range cardsDesc.Data.List {
		logrus.WithFields(logrus.Fields{
			"名称": cardDesc.ScName,
		}).Infof("卡牌描述")
	}
}

// 获取卡片信息，并以 JSON 格式写入到文件中
func TestSearchClient_PostCardSearch(t *testing.T) {
	initTest()

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
		cards, err := client.PostCardSearch(id, "300", "chs", "")
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
	initTest()

	got, err := client.GetCardPrice("1896")
	if err != nil {
		logrus.Fatalln(err)
	}

	fmt.Println(got)
}
