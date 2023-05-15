package main

import (
	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/handler"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/router"
	logging "github.com/DesistDaydream/logging/pkg/logrus_init"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	Debug bool
}

func AddFlags(f *Flags) {
	pflag.BoolVarP(&f.Debug, "debug", "d", false, "是否开启 debug 模式")
}

func main() {
	var (
		flags    Flags
		logFlags logging.LogrusFlags
	)
	AddFlags(&flags)
	logging.AddFlags(&logFlags)
	pflag.Parse()
	if err := logging.LogrusInit(&logFlags); err != nil {
		logrus.Fatalf("初始化日志失败: %v", err)
	}

	// 初始化配置文件
	c, _ := config.NewConfig("", "")

	// 连接数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}
	database.InitDB(dbInfo)

	if !flags.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handler.H = handler.NewHandler(c.DtcgDB.Username, c.DtcgDB.Password, c.DtcgDB.Token, c.DtcgDB.Retry)

	r := router.InitRouter()
	r.Run(c.Listen)
}
