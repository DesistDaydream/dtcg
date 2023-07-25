package main

import (
	"os"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/orders"
	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/products"
	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/wishes"
	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"

	logging "github.com/DesistDaydream/logging/pkg/logrus_init"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Flags struct {
	enable_dtcgdb_auth bool
	ConfigPath         string
	ConfigName         string
}

var (
	flags    Flags
	logFlags logging.LogrusFlags
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
	RootCmd.PersistentFlags().BoolVar(&flags.enable_dtcgdb_auth, "enable-dtcgdb-auth", false, "是否使用卡查网站的 TOKEN 以获取私密信息")
	RootCmd.PersistentFlags().StringVar(&flags.ConfigPath, "config-path", "", "配置文件路径")
	RootCmd.PersistentFlags().StringVar(&flags.ConfigName, "config-name", "", "配置文件名称")

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
	if err := logging.LogrusInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	c, _ := config.NewConfig(flags.ConfigPath, flags.ConfigName)

	// 连接数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	// 实例化一个处理器，包括各种 SDK 的服务能力
	if flags.enable_dtcgdb_auth {
		handler.H = handler.NewHandler(true, "1", c.Moecard.Username, c.Moecard.Password)
	} else {
		handler.H = handler.NewHandler(false, "1", "", "")
	}
}
