package community

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/sirupsen/logrus"
)

var token string = ""

// var cardVersionID string = "2544"

func getToken() {
	file, err := os.ReadFile("../token.txt")
	if err != nil {
		logrus.Fatal(err)
	}
	token = string(file)
}

func initDB() {
	// 初始化配置文件
	c := config.NewConfig("../../../../../config", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)
}

func TestCommunityClient_PostConvertDeck(t *testing.T) {
	getToken()
	decksjson := `["Exported from http://digimon.card.moe","ST1-01","ST1-03","ST1-03","ST1-03","ST1-06","ST1-06","ST1-07","ST1-07","ST1-07","ST1-07","ST1-16","ST1-16","BT1-010","BT1-010","BT1-020","BT1-020","BT1-020","BT1-020","BT1-025","BT1-025","BT1-084","BT1-085","P-009","P-009","P-009","P-009","BT4-019","BT4-019","BT4-092","BT4-099","BT4-099","BT4-100","BT5-001","BT5-001","BT5-001","BT5-001","BT5-007","BT5-007","BT5-007","BT5-007","BT5-010","BT5-010","BT5-010","BT5-010","BT5-015","BT5-015","BT5-015","BT5-015","BT5-016","BT5-016","BT5-086","BT5-086","BT5-092","BT5-092","BT5-092"]`
	client := NewCommunityClient(core.NewClient(token))
	decks, err := client.PostDeckConvert(decksjson)
	if err != nil {
		logrus.Fatalln(err)
	}

	var (
		minPrice float64
		avgPrice float64
	)

	var cardsID []string

	for _, card := range decks.Data.DeckInfo.Eggs {
		cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
	}

	for _, card := range decks.Data.DeckInfo.Main {
		cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
	}

	clientSearch := cdb.NewCdbClient(core.NewClient(""))
	for _, cardID := range cardsID {
		cardPrice, err := clientSearch.GetCardPrice(cardID)
		if err != nil {
			logrus.Errorf("获取卡片价格失败: %v", err)
		}

		var fMin float64
		if len(cardPrice.Data.Products) == 0 {
			fMin = 0
		} else {
			fMin, _ = strconv.ParseFloat(cardPrice.Data.Products[0].MinPrice, 64)
		}

		fAvg, _ := strconv.ParseFloat(cardPrice.Data.AvgPrice, 64)

		minPrice = minPrice + fMin
		avgPrice = avgPrice + fAvg
	}

	logrus.WithFields(logrus.Fields{
		"最低价": minPrice,
		"集换价": avgPrice,
	}).Infof("卡组价格")

}

// 从自己的卡组中获取卡组详情
func TestCommunityClient_GetDeckCloud(t *testing.T) {
	getToken()
	client := NewCommunityClient(core.NewClient(token))
	decks, err := client.GetDeckCloud("106981")
	if err != nil {
		logrus.Errorln(err)
	}

	var cardsID []string

	for _, card := range decks.Data.DeckInfo.Eggs {
		for i := 0; i < card.Number; i++ {
			cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
		}
	}

	for _, card := range decks.Data.DeckInfo.Main {
		for i := 0; i < card.Number; i++ {
			cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
		}
	}

	logrus.Infoln(len(cardsID), cardsID)

	var (
		minPrice float64
		avgPrice float64
	)
	initDB()

	for _, cardID := range cardsID {
		cardPrice, err := database.GetCardPrice(cardID)
		if err != nil {
			logrus.Fatalln(err)
		}
		minPrice = minPrice + cardPrice.MinPrice
		avgPrice = avgPrice + cardPrice.AvgPrice
	}

	logrus.WithFields(logrus.Fields{
		"最低价": minPrice,
		"集换价": avgPrice,
	}).Infof("卡组价格")

}

// 从卡组广场的卡组中获取卡组详情
func TestCommunityClient_GetDeck(t *testing.T) {
	client := NewCommunityClient(core.NewClient(token))
	decks, err := client.GetDeck("6cea907f6a001007281eaa8f52feb517a811a5bd")
	if err != nil {
		logrus.Errorln(err)
	}

	var cardsID []string

	for _, card := range decks.Data.DeckInfo.Eggs {
		for i := 0; i < card.Number; i++ {
			cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
		}
	}

	for _, card := range decks.Data.DeckInfo.Main {
		for i := 0; i < card.Number; i++ {
			cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
		}
	}

	logrus.Infoln(len(cardsID), cardsID)

	cardsIDString, _ := json.Marshal(&cardsID)

	logrus.Infoln(string(cardsIDString))
}
