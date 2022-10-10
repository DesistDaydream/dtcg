package main

import (
	"github.com/DesistDaydream/dtcg/cmd/my_dtcg/card"
	cardgroup "github.com/DesistDaydream/dtcg/cmd/my_dtcg/card_group"
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

	database.InitDB()

	switch flags.Add {
	case "card":
		card.AddCard()
	case "cardgroup":
		cardgroup.AddCardGroup(false)
	default:
		logrus.Errorln("使用 --add 指定要添加的数据")
	}
}
