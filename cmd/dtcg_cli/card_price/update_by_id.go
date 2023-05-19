package cardprice

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/spf13/cobra"
)

// 更新指定卡牌 ID 的价格
func UpdateByIDCommand() *cobra.Command {
	updateIDCmd := &cobra.Command{
		Use:   "id",
		Short: "更新指定卡牌 ID 的价格。这里面的 ID 是 card_id_from_db 的值，多个 ID 以逗号分隔",
		Run:   updateID,
	}

	return updateIDCmd
}

func updateID(cmd *cobra.Command, args []string) {
	for _, cardIDFromDB := range args {
		cardDesc, _ := database.GetCardDescByCardIDFromDB(fmt.Sprint(cardIDFromDB))
		updateRun(cardDesc)
	}
}
