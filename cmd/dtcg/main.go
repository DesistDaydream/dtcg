package main

import (
	"time"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/flags"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/handler"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/router"
	logging "github.com/DesistDaydream/logging/pkg/logrus_init"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	var (
		f        flags.Flags
		logFlags logging.LogrusFlags
	)
	flags.AddFlags(&f)
	logging.AddFlags(&logFlags)
	pflag.Parse()
	if err := logging.LogrusInit(&logFlags); err != nil {
		logrus.Fatalf("初始化日志失败: %v", err)
	}

	// 初始化配置文件
	config.NewConfig("", "")

	// 连接数据库
	dbInfo := &database.DBInfo{
		FilePath: config.Conf.SQLite.FilePath,
		Server:   config.Conf.Mysql.Server,
		Password: config.Conf.Mysql.Password,
	}
	database.InitDB(dbInfo)

	handler.H = handler.NewHandler(config.Conf.Moecard.Username, config.Conf.Moecard.Password, config.Conf.Moecard.Retry)

	if !f.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	duration, err := time.ParseDuration(config.Conf.JHS.AutoUpdateTokenDuration)
	if err != nil {
		logrus.Fatalf("解析自动更新 Token 时间失败: %v", err)
	}
	ticker := time.NewTicker(duration)
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
	r.Run(config.Conf.Listen)
}
