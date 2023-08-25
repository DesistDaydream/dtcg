package products

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/sirupsen/logrus"
)

var sellerUserID string = "934972"
var token string = ""
var cardVersionID string = "4510"

func initConfig() {
	// 初始化配置文件
	c, _ := config.NewConfig("../../../../../config", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		DBType:   c.DBType,
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	user, _ := database.GetUser("1")
	token = user.JhsToken
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
