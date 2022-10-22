package main

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/router"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Config struct {
	Mysql MySQL `json:"mysql"`
}
type MySQL struct {
	Server   string `json:"server"`
	Password string `json:"password"`
}

type Flags struct {
	Debug      bool
	ListenAddr string
	MySQL      MySQL
}

func AddFlags(f *Flags) {
	pflag.BoolVarP(&f.Debug, "debug", "d", false, "是否开启 debug 模式")
	pflag.StringVarP(&f.ListenAddr, "listen", "l", ":2205", "程序监听地址")
	pflag.StringVar(&f.MySQL.Server, "server", "", "数据库连接地址")
	pflag.StringVar(&f.MySQL.Password, "password", "", "数据库连接密码")
}

func main() {
	var (
		flags    Flags
		logFlags logging.LoggingFlags
	)
	AddFlags(&flags)
	logging.AddFlags(&logFlags)
	pflag.Parse()
	if err := logging.LogInit(&logFlags); err != nil {
		logrus.Fatalf("初始化日志失败: %v", err)
	}

	// 初始化配置文件
	// viper.SetConfigName("my_dtcg.yaml")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("./config_file")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	logrus.Fatalf("读取配置文件失败: %v", err)
	// }

	dbInfo := &database.DBInfo{
		FilePath: "internal/database/my_dtcg.db",
		// Server:   viper.GetString("mysql.server"),
		// Password: viper.GetString("mysql.password"),
		Server:   flags.MySQL.Server,
		Password: flags.MySQL.Password,
	}
	database.InitDB(dbInfo)

	if !flags.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()
	r.Run(flags.ListenAddr)
}
