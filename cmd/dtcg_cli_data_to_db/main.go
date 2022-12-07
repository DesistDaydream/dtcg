package main

import (
	carddesc "github.com/DesistDaydream/dtcg/cmd/dtcg_cli_data_to_db/card_desc"
	cardgroup "github.com/DesistDaydream/dtcg/cmd/dtcg_cli_data_to_db/card_group"
	"github.com/DesistDaydream/dtcg/cmd/dtcg_cli_data_to_db/handler"
	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/pkg/database"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	Add string
}

func AddFlags(f *Flags) {
	pflag.StringVarP(&f.Add, "add", "a", "", "向数据库添加数据的内容")

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
		logrus.Fatal("初始化日志失败", err)
	}

	c := config.NewConfig("", "")

	// 连接数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	handler.H = handler.NewHandler(c.DtcgDB.Username, c.DtcgDB.Password, 1)

	switch flags.Add {
	case "cardgroupofficial":
		cardgroup.AddCardGroupFromOfficial()
	case "carddescofficial":
		carddesc.AddCardDescFromOfficial()
	case "cardgroupdtcgdb":
		cardgroup.AddCardGroupFromDtcgDB()
	case "carddescdtcgdb":
		carddesc.AddCardDescFromDtcgDB()
	default:
		logrus.Errorln("使用 --add 指定要添加的数据")
	}
}
