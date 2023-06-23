package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/wishes"
	"github.com/sirupsen/logrus"
)

var token string = ""
var wishListID string = "2610301"
var wishesClient *wishes.WishesClient
var productsClient *products.ProductsClient

func initConfig() {
	// 初始化配置文件
	c, _ := config.NewConfig("./config", "")

	// 初始化数据库
	// dbInfo := &database.DBInfo{
	// 	FilePath: c.SQLite.FilePath,
	// 	Server:   c.Mysql.Server,
	// 	Password: c.Mysql.Password,
	// }

	// database.InitDB(dbInfo)

	token = c.JHS.Token
}

func init() {
	initConfig()
	wishesClient = wishes.NewWishesClient(core.NewClient(token))
	productsClient = products.NewProductsClient(core.NewClient(token))
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 4, 8, 4, ' ', 0)
	// fmt.Fprintf(w, "卡牌名称\t我的价格\tta的价格\t我的数量\tta的数量\n")

	resp, err := wishesClient.WishListMatch(wishListID)
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, card := range resp[0].MatchCards {
		products, err := productsClient.List("1", card.Number, "1", "price_asc")
		if err != nil {
			logrus.Fatal(err)
		}

		if len(products.Data) > 0 {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%v\t\n", card.CardName, products.Data[0].Price, card.Price, products.Data[0].Quantity, card.Quantity)
		} else {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%v\t\n", card.CardName, "0", card.Price, 0, card.Quantity)
		}
	}
	w.Flush()
}
