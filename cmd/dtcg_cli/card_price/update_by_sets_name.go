package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func UpdateSetsNameCommand() *cobra.Command {
	updateSetsNameCmd := &cobra.Command{
		Use:   "sets-name",
		Short: "更新指定卡牌集合的卡牌价格，使用 card-set list 子命令获取卡包名称。若不指定则更新所有",
		Run:   updateSetsName,
	}

	return updateSetsNameCmd
}

func updateSetsName(cmd *cobra.Command, args []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("获取全部卡牌描述失败: %v", err)
	}

	for _, cardDesc := range cardsDesc.Data {
		for _, setPrefix := range args {
			if cardDesc.SetPrefix == setPrefix {
				updateRun(&cardDesc)
			}
		}
	}
}
