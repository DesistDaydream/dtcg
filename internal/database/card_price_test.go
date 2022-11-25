package database

import (
	"testing"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

func initDB() {
	// 初始化配置文件
	c := config.NewConfig("../../config", "")

	// 初始化数据库
	dbInfo := &DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	InitDB(dbInfo)
}

func TestGetCardPrice(t *testing.T) {
	initDB()

	got, err := GetCardPrice("2210")
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Info(got)
}

func TestUpdateCardPrice(t *testing.T) {
	initDB()

	// 实例化一个处理器，包括各种 SDK 的服务能力
	handler.H = handler.NewHandler()

	cardsPrice, err := ListCardsPrice()
	if err != nil {
		logrus.Fatalf("获取卡片价格信息失败: %v", err)
	}

	for _, cardPrice := range cardsPrice.Data {
		if cardPrice.ImageUrl == "" {
			UpdateCardPrice(&cardPrice, map[string]string{})
		}
	}
}

func TestGetCardPriceWhereSetPrefix(t *testing.T) {
	initDB()

	// 实例化一个处理器，包括各种 SDK 的服务能力
	// handler.H = handler.NewHandler()

	got, err := GetCardPriceWhereSetPrefix("STC-07")
	if err != nil {
		logrus.Errorf("%v", err)
	}
	logrus.Infof("%v", got)
}

func TestGetCardPriceByCondition(t *testing.T) {
	initDB()

	got, err := GetCardPriceByCondition(3, 1, &models.QueryCardPrice{
		CardPack:   0,
		ClassInput: false,
		Color:      []string{"红", "白"},
		EvoCond:    []models.EvoCond{},
		Keyword:    "奥米加",
		Language:   "",
		OrderType:  "",
		QField:     []string{},
		Rarity:     []string{},
		Tags:       []string{},
		TagsLogic:  "",
		Type:       "",
	})
	if err != nil {
		logrus.Errorln(err)
	}

	for _, v := range got.Data {
		logrus.WithFields(logrus.Fields{
			"CardIDFromDB":  v.CardIDFromDB,
			"CardVersionID": v.CardVersionID,
			"图片":            v.ImageUrl,
		}).Infof("查询结果")
	}
}

func TestGetCardPriceWithImageByCondition(t *testing.T) {
	initDB()

	got, err := GetCardPriceWithDtcgDBImgByCondition(3, 1, &models.QueryCardPrice{
		CardPack:   0,
		ClassInput: false,
		Color:      []string{"红", "白"},
		EvoCond:    []models.EvoCond{},
		Keyword:    "奥米加",
		Language:   "",
		OrderType:  "",
		QField: []string{
			"serial",
			"sc_name",
		},
		Rarity:    []string{},
		Tags:      []string{},
		TagsLogic: "",
		Type:      "",
	})
	if err != nil {
		logrus.Errorln(err)
	}

	for _, v := range got.Data {
		logrus.WithFields(logrus.Fields{
			"CardIDFromDB":  v.CardIDFromDB,
			"CardVersionID": v.CardVersionID,
			"图片":            v.Image,
		}).Infof("查询结果")
	}
}
