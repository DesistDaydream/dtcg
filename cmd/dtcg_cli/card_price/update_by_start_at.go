package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// 从指定的卡牌开始更新，指定 card_id_from_db
func UpdateByStartAtCommand() *cobra.Command {
	updateStartAtCmd := &cobra.Command{
		Use:   "start-at",
		Short: "从哪张卡牌开始更新。参数为 card_id_from_db 的值",
		Run:   updateStartAt,
	}

	return updateStartAtCmd
}

func updateStartAt(cmd *cobra.Command, args []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("获取全部卡牌描述失败: %v", err)
	}

	// startAt 是一个数组中元素的索引。也就是从数据库获取的 cardsDesc.data 这个数组的索引。
	// 首先根据 card_id_from_db 的号推导出 id 号，然后 startAt 就是 id-1
	var startAt int

	switch len(args) {
	case 0:
		startAt = 0
	case 1:
		cardDesc, err := database.GetCardDescByCardIDFromDB(args[0])
		if err != nil {
			logrus.Errorf("获取 card_id_from_db 为 %v 的卡牌信息失败: %v", args[0], err)
		}
		startAt = cardDesc.ID - 1
	default:
		logrus.Fatalln("参数异常，使用一个参数")
	}

	for i := startAt; i < len(cardsDesc.Data); i++ {
		updateRun(&cardsDesc.Data[i])
	}
}
