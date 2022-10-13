package main

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/router"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	Debug      bool
	ListenAddr string
}

func (f *Flags) AddFlags() {
	pflag.BoolVarP(&f.Debug, "debug", "d", false, "是否开启 debug 模式")
	pflag.StringVarP(&f.ListenAddr, "listen", "l", ":2205", "程序监听地址")
}

func main() {
	flags := Flags{}
	flags.AddFlags()
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()
	pflag.Parse()
	if err := logging.LogInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	dbInfo := &database.DBInfo{
		FilePath: "internal/database/my_dtcg.db",
	}
	database.InitDB(dbInfo)

	if !flags.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()
	r.Run(flags.ListenAddr)
}
