package database

import (
	"testing"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/config"
	"github.com/sirupsen/logrus"
)

func TestGetCardPrice(t *testing.T) {
	dbInfo := &DBInfo{
		FilePath: "my_dtcg.db",
	}
	InitDB(dbInfo)
	got, err := GetCardPrice("2210")
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Info(got)
}

func TestUpdateCardPrice(t *testing.T) {
	// 初始化配置文件
	c := config.NewConfig("", "")

	// 初始化数据库
	dbInfo := &DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	InitDB(dbInfo)

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

	// 初始化配置文件
	c := config.NewConfig("../../config", "")

	// 初始化数据库
	dbInfo := &DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	InitDB(dbInfo)

	// 实例化一个处理器，包括各种 SDK 的服务能力
	// handler.H = handler.NewHandler()

	got, err := GetCardPriceWhereSetPrefix("STC-07")
	if err != nil {
		logrus.Errorf("%v", err)
	}
	logrus.Infof("%v", got)
}
