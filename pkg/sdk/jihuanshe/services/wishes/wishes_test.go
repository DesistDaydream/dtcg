package wishes

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/pkg/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

// var sellerUserID string = "609077"
var (
	token      string = ""
	wishListID string = "2610132"
	client     *WishesClient
	table      *tablewriter.Table
)

// var cardVersionID string = "3982"

func initConfig() {
	// 初始化配置文件
	c, _ := config.NewConfig("../../../../../config", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	token = c.JHS.Token
}

func init() {
	initConfig()
	client = NewWishesClient(core.NewClient(token))
	table = tablewriter.NewWriter(os.Stdout)
}

// 创建清单测试
func TestWishesClient_CreateList(t *testing.T) {
	resp, err := client.CreateWashList("测试清单")
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Infoln(resp)

	wishListID = fmt.Sprint(resp.WishListID)
}

// 向清单中添加卡牌测试
func TestWishesClient_Add(t *testing.T) {
	wishListID = "1794222"
	resp, err := client.Add("3850", "0", "4", "", wishListID)
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Infoln(resp)
}

// 列出官方推荐的清单
func TestWishesClient_GetRecommendList(t *testing.T) {
	table.SetHeader([]string{"名称", "清单ID"})

	resp, err := client.ListWishListRecommend()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, wishListRecommendData := range resp.Data {
		table.Append([]string{wishListRecommendData.Name, strconv.FormatInt(wishListRecommendData.WishListID, 10)})
	}

	table.Render()
}

// 获取清单详情
func TestWishesClient_Get(t *testing.T) {
	resp, err := client.Get(wishListID)
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, data := range resp.Data {
		table.Append([]string{data.NameCN, data.Number, strconv.FormatInt(data.Quantity, 10), data.MinPrice})
	}

	table.Render()
}

// 一键匹配清单
func TestWishesClient_WishListMatch(t *testing.T) {
	resp, err := client.WishListMatch(wishListID)
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, card := range resp[0].MatchCards {
		table.Append([]string{card.CardName, card.Number, card.Price, strconv.FormatInt(card.Quantity, 10)})
	}

	table.Render()
}

// 通用测试
func TestCommon(t *testing.T) {

}
