package main

import (
	"time"

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
	Debug                   bool
	AutoUpdateTokenDuration time.Duration
}

func AddFlags(f *Flags) {
	pflag.BoolVarP(&f.Debug, "debug", "d", false, "是否开启 Gin 的 debug 模式")
	pflag.DurationVar(&f.AutoUpdateTokenDuration, "duration", 60*time.Minute, "每次更新集换社 Token 的间隔时间")
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

	handler.H = handler.NewHandler(c.Moecard.Username, c.Moecard.Password, c.Moecard.Retry)

	if !flags.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	ticker := time.NewTicker(flags.AutoUpdateTokenDuration)
	go func() {
		for range ticker.C {
			_, err := handler.H.JhsServices.Market.AuthUpdateTokenPost()
			if err != nil {
				logrus.Errorf("更新集换社 Token 失败，原因: %v", err)
			}
			logrus.Infof("更新集换社 Token 成功")
		}
	}()

	r := router.InitRouter()
	r.Run(c.Listen)
}
