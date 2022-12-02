package main

import (
	"fmt"

	cardprice "github.com/DesistDaydream/dtcg/cmd/dtcg_cli/card_price"
	"github.com/DesistDaydream/dtcg/cmd/dtcg_cli/handler"
	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化配置文件
	c := config.NewConfig("", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	// 实例化一个处理器，包括各种 SDK 的服务能力
	handler.H = handler.NewHandler()

	cardsPrice, err := database.ListCardsPrice()
	if err != nil {
		logrus.Fatalf("获取卡片价格信息失败: %v", err)
	}

	for _, cardPrice := range cardsPrice.Data {
		if cardPrice.ImageUrl == "" {
			fmt.Println("开始处理 ", cardPrice.ScName)
			img := cardprice.GetImageURL(cardPrice.CardVersionID)
			cardPrice.ImageUrl = img
			database.UpdateCardPrice(&cardPrice, map[string]interface{}{})
		}
	}
}
