package adddatatodb

import (
	carddesc "github.com/DesistDaydream/dtcg/cmd/dtcg_cli/add_data_to_db/card_desc"
	cardgroup "github.com/DesistDaydream/dtcg/cmd/dtcg_cli/add_data_to_db/card_group"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type AddDataToDBFlags struct {
	Target string
}

var addDataToDBFlags AddDataToDBFlags

func AddFlags(f *AddDataToDBFlags) {
	pflag.StringVar(&f.Target, "target", "", "向数据库添加数据的内容")

}
func CreateCommand() *cobra.Command {
	AddDataToDBCmd := &cobra.Command{
		Use:   "add-data-to-db",
		Short: "向我的数据库中添加 DTCG 中文官网或者 Moecard 卡查的数据",
		Run:   AddDataToDB,
	}

	return AddDataToDBCmd
}

func AddDataToDB(cmd *cobra.Command, args []string) {

	switch addDataToDBFlags.Target {
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
