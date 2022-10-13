package main

import (
	"os"

	carddesc "github.com/DesistDaydream/dtcg/cmd/my_dtcg/card_desc"
	cardprice "github.com/DesistDaydream/dtcg/cmd/my_dtcg/card_price"
	cardset "github.com/DesistDaydream/dtcg/cmd/my_dtcg/card_set"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Flags struct {
	Add string
}

func AddFlsgs(f *Flags) {
	pflag.StringVarP(&f.Add, "add", "a", "", "向数据库添加数据的内容")

}

var (
	flags    Flags
	logFlags logging.LoggingFlags
)

func main() {
	app := newApp()
	err := app.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newApp() *cobra.Command {
	long := `
数码宝贝集换式卡牌游戏，简称 DTCG。该工具可以从如下几个地方获取卡牌数据并保存到自己的数据库中
	[官网](https://www.digimoncard.cn/)
	[数码兽卡片游戏数据库](https://digimon.card.moe/)
还可以从集换社获取卡牌的价格
`

	var RootCmd = &cobra.Command{
		Use:              "mydtcg",
		Short:            "我的 DTCG 管理工具",
		Long:             long,
		PersistentPreRun: rootPersistentPreRun,
	}

	AddFlsgs(&flags)
	logging.AddFlags(&logFlags)

	// 添加子命令
	RootCmd.AddCommand(
		cardset.CreateCommand(),
		carddesc.CreateCommand(),
		cardprice.CreateCommand(),
	)

	return RootCmd
}

// 执行每个 root 下的子命令时，都需要执行的函数
func rootPersistentPreRun(cmd *cobra.Command, args []string) {
	if err := logging.LogInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	dbInfo := &database.DBInfo{
		FilePath: "internal/database/my_dtcg.db",
	}

	database.InitDB(dbInfo)
}
