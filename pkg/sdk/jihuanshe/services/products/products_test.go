package products

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
)

var sellerUserID string = "609077"
var token string = ""
var cardVersionID string = "4282"
var productID string = ""

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

// 测试将结构体转为 map
func TestStructToMapStr(t *testing.T) {
	obj := models.ProductsGetReqQuery{
		GameKey:       "dgm",
		SellerUserID:  sellerUserID,
		CardVersionID: cardVersionID,
	}

	got := core.StructToMapStr(&obj)

	fmt.Println(len(got))

	gotByte, _ := json.Marshal(got)
	fmt.Println(string(gotByte))
	for k, v := range got {
		fmt.Println(k, v)
	}
}

// 获取商品信息
func TestProductsClientGet(t *testing.T) {
	initConfig()
	client := NewProductsClient(core.NewClient(token))

	got, err := client.Get(cardVersionID, sellerUserID)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(got)
}
