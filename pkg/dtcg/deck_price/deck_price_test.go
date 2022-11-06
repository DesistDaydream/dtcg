package deckprice

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/sirupsen/logrus"
)

func initDB() {
	// 初始化配置文件
	c := config.NewConfig("../../../config", "")

	// 连接数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}
	database.InitDB(dbInfo)
}

func Test_transform(t *testing.T) {
	initDB()
	ids := "[\"1896\",\"1896\",\"1896\",\"1897\",\"1898\",\"1898\"]"

	got, err := transform(ids)
	if err != nil {
		logrus.Errorln(err)
	}

	fmt.Println(got)
}
