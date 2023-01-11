package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func UpdateStartAtCommand() *cobra.Command {
	updateStartAtCmd := &cobra.Command{
		Use:   "start-at",
		Short: "从哪张卡牌开始更新。参数为 card_id_from_db 的值",
		Run:   updateStartAt,
	}

	return updateStartAtCmd
}

// 从指定的卡牌开始更新
func updateStartAt(cmd *cobra.Command, args []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("获取全部卡牌描述失败: %v", err)
	}

	var startAt int

	switch len(args) {
	case 0:
		startAt = 0

	case 1:
		// 从数据库中获取卡牌的ID
		cardDesc, _ := database.GetCardDescByCardIDFromDB(args[0])
		// 开始更新的位置是id-1，因为数组元素编号是从 0 开始的
		startAt = cardDesc.ID - 1
	}

	for i := startAt; i < len(cardsDesc.Data); i++ {
		updateRun(&cardsDesc.Data[i])
	}
}
