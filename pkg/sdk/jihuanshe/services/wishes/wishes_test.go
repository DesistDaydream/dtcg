package wishes

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/pkg/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/sirupsen/logrus"
)

// var sellerUserID string = "609077"
var token string = ""
var wishListID string = ""

// var cardVersionID string = "3982"

func initConfig() {
	// 初始化配置文件
	c := config.NewConfig("../../../../../config", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	token = c.JHS.Token
}

// 创建清单测试
func TestWishesClient_CreateList(t *testing.T) {
	initConfig()
	client := NewWishesClient(core.NewClient(token))

	resp, err := client.CreateList("测试清单")
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Infoln(resp)

	wishListID = fmt.Sprint(resp.WishListID)
}

// 向清单中添加卡牌测试
func TestWishesClient_Add(t *testing.T) {
	initConfig()
	client := NewWishesClient(core.NewClient(token))

	wishListID = "1794222"
	resp, err := client.Add("3850", "0", "4", "", wishListID)
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Infoln(resp)
}
