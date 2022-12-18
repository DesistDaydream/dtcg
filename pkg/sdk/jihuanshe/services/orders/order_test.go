package orders

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/pkg/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/sirupsen/logrus"
)

var token string = ""

// var client *CommunityClient

// var cardVersionID string = "2544"

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

// 获取用户订单列表（买入）
func TestOrdersClient_GetBuyerOrders(t *testing.T) {
	initConfig()
	client := NewOrdersClient(core.NewClient(token))

	resp, err := client.GetBuyerOrders("1")
	if err != nil {
		logrus.Errorln(err)
	}
	for _, data := range resp.Data {
		logrus.Infoln(data.OrderID)
	}
}

func TestOrdersClient_GetBuyerOrderProducts(t *testing.T) {
	initConfig()
	client := NewOrdersClient(core.NewClient(token))
	resp, err := client.GetBuyerOrderProducts(2479030)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(resp)
}

// 获取用户订单列表（卖出）
func TestOrdersClient_GetSellerOrders(t *testing.T) {
	initConfig()
	client := NewOrdersClient(core.NewClient(token))

	resp, err := client.GetSellerOrders("1")
	if err != nil {
		logrus.Errorln(err)
	}
	for _, data := range resp.Data {
		logrus.Infoln(data.OrderID)
	}
}

func TestOrdersClient_GetSellerOrderProducts(t *testing.T) {
	initConfig()
	client := NewOrdersClient(core.NewClient(token))
	resp, err := client.GetSellerOrderProducts(2475268)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(resp)
}
