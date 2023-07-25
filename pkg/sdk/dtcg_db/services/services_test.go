package services

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/sirupsen/logrus"
)

func TestNewServices(t *testing.T) {
	// 初始化配置文件
	c, _ := config.NewConfig("../../../../config", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	user, err := database.GetUser("1")
	if err != nil {
		logrus.Fatalf("获取用户信息异常，原因: %v", err)
	}

	services := NewServices(true, c.Moecard.Username, c.Moecard.Password, user.MoecardToken, 2)

	fmt.Println(services.Cdb, services.Community)
}
