package database

import (
	"testing"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

var jhsToken string

func initDB() {
	// 初始化配置文件
	c, _ := config.NewConfig("../../config", "")

	// 初始化数据库
	dbInfo := &DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	InitDB(dbInfo)

	jhsToken = c.JHS.Token
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
	handler.H = handler.NewHandler(false, jhsToken, "", "", "")

	cardsPrice, err := ListCardsPrice()
	if err != nil {
		logrus.Fatalf("获取卡片价格信息失败: %v", err)
	}

	for _, cardPrice := range cardsPrice.Data {
		if cardPrice.ImageUrl == "" {
			UpdateCardPrice(&models.CardPrice{CardIDFromDB: cardPrice.CardIDFromDB}, map[string]interface{}{
				"image_url": "找不到图片地址",
			})
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

	got, err := GetCardPriceByCondition(10, 1, &models.CardPriceQuery{
		CardVersionID: 0,
		// SetsPrefix:     []string{"BTC-02", "BTC-03"},
		SetsPrefix:     []string{"BTC-05"},
		Keyword:        "",
		QField:         []string{},
		Rarity:         []string{"u", "r"},
		AlternativeArt: "",
		MinPriceRange:  "",
		AvgPriceRange:  "",
	})
	if err != nil {
		logrus.Errorln(err)
	}

	for _, v := range got.Data {
		logrus.WithFields(logrus.Fields{
			"CardIDFromDB":  v.CardIDFromDB,
			"CardVersionID": v.CardVersionID,
			// "图片":            v.ImageUrl,
			"卡片名称": v.ScName,
			"集换价":  v.AvgPrice,
			"最低价":  v.MinPrice,
		}).Infof("查询结果")
	}
}

func TestGetCardPriceWithImageByCondition(t *testing.T) {
	initDB()

	got, err := GetCardPriceWithDtcgDBImgByCondition(3, 1, &models.CardPriceQuery{
		SetsPrefix:     []string{"BTC-02", "BTC-03"},
		Keyword:        "奥米加",
		Language:       "",
		QField:         []string{"serial", "sc_name"},
		Rarity:         []string{},
		AlternativeArt: "",
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
