package main

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/router"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	ListenAddr string
}

func (f *Flags) AddFlags() {
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

	i := &database.DBInfo{
		FilePath: "internal/database/my_dtcg.db",
	}
	database.InitDB(i)

	r := router.InitRouter()
	r.Run(flags.ListenAddr)
}
