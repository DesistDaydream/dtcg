package services

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/pkg/database"
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

	got := NewServices(true, c.DtcgDB.Username, c.DtcgDB.Password, "", 2)

	fmt.Println(got.CoreClient)

}
