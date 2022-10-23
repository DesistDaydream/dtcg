package main

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/router"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Listen string `yaml:"listen"`
	Mysql  MySQL  `yaml:"mysql"`
	SQLite SQLite `yaml:"sqlite"`
}
type MySQL struct {
	Server   string `yaml:"server"`
	Password string `yaml:"password"`
}

type SQLite struct {
	FilePath string `yaml:"file_path"`
}

type Flags struct {
	Debug bool
}

func AddFlags(f *Flags) {
	pflag.BoolVarP(&f.Debug, "debug", "d", false, "是否开启 debug 模式")
}

func main() {
	var (
		config   Config
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
	viper.AddConfigPath("/etc/dtcg")
	viper.AddConfigPath("./config_file")
	viper.SetConfigName("my_dtcg.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("读取配置文件失败: %v", err)
	}
	viper.Unmarshal(&config)

	// 连接数据库
	dbInfo := &database.DBInfo{
		FilePath: config.SQLite.FilePath,
		Server:   config.Mysql.Server,
		Password: config.Mysql.Password,
	}
	database.InitDB(dbInfo)

	if !flags.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()
	r.Run(config.Listen)
}
