package main

import (
	"os"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/orders"
	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/products"
	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/wishes"
	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Flags struct {
	enable_dtcgdb_auth bool
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
集换社是一个国内的垃圾卡牌交易市场。但是奈何没有竞争对手，并且起步较早，导致其摆烂但是又有很多人用。
该工具用来对集换社进行控制，可以批量上架、更新、下架自己的商品。
`

	var RootCmd = &cobra.Command{
		Use:   "myjhs",
		Short: "我的集换社管理工具",
		Long:  long,
		// PersistentPreRun: rootPersistentPreRun,
	}

	cobra.OnInitialize(initConfig)

	logging.AddFlags(&logFlags)
	RootCmd.PersistentFlags().BoolVar(&flags.enable_dtcgdb_auth, "enable-dtcgdb-auth", false, "DTCG DB 中我的卡组的 ID")

	// 添加子命令
	RootCmd.AddCommand(
		products.CreateCommand(),
		orders.CreateCommand(),
		wishes.CreateCommand(),
	)

	return RootCmd
}

// 执行每个 root 下的子命令时，都需要执行的函数
func initConfig() {
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

	// 实例化一个处理器，包括各种 SDK 的服务能力
	if flags.enable_dtcgdb_auth {
		handler.H = handler.NewHandler(true, c.JHS.Token, c.DtcgDB.Username, c.DtcgDB.Password, c.DtcgDB.Token)
	} else {
		handler.H = handler.NewHandler(false, c.JHS.Token, "", "", "")
	}
}
