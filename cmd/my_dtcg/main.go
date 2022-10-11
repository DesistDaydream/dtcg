package main

import (
	carddesc "github.com/DesistDaydream/dtcg/cmd/my_dtcg/card_desc"
	cardgroup "github.com/DesistDaydream/dtcg/cmd/my_dtcg/card_group"
	cardprice "github.com/DesistDaydream/dtcg/cmd/my_dtcg/card_price"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	Add string
}

func AddFlsgs(f *Flags) {
	pflag.StringVarP(&f.Add, "add", "a", "", "向数据库添加数据的内容")

}

func main() {
	var flags Flags
	AddFlsgs(&flags)
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()
	pflag.Parse()

	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	i := &database.DBInfo{
		FilePath: "internal/database/my_dtcg.db",
	}
	database.InitDB(i)

	switch flags.Add {
	case "cardgroup":
		cardgroup.AddCardGroup(false)
	case "carddesc":
		carddesc.AddCardDesc()
	case "cardgroupdtcgdb":
		cardgroup.AddCardGroupFromDtcgDB()
	case "carddescdtcgdb":
		carddesc.AddCardDescFromDtcgDB()
	case "cardprice":
		cardprice.AddCardPrice("2780")
	default:
		logrus.Errorln("使用 --add 指定要添加的数据")
	}
}
