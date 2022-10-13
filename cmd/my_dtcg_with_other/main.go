package main

import (
	carddesc "github.com/DesistDaydream/dtcg/cmd/my_dtcg_with_other/card_desc"
	cardgroup "github.com/DesistDaydream/dtcg/cmd/my_dtcg_with_other/card_group"
	"github.com/DesistDaydream/dtcg/internal/database"
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

	i := &database.DBInfo{
		FilePath: "internal/database/my_dtcg.db",
	}
	database.InitDB(i)

	switch flags.Add {
	case "cardgroupofficial":
		cardgroup.AddCardGroupOfficial(false)
	case "carddescofficial":
		carddesc.AddCardDescOfficial()
	case "cardgroupdtcgdb":
		cardgroup.AddCardGroupFromDtcgDB()
	case "carddescdtcgdb":
		carddesc.AddCardDescFromDtcgDB()
	default:
		logrus.Errorln("使用 --add 指定要添加的数据")
	}
}
