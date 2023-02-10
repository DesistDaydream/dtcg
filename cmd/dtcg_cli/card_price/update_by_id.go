package cardprice

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/spf13/cobra"
)

func UpdateByIDCommand() *cobra.Command {
	updateIDCmd := &cobra.Command{
		Use:   "id",
		Short: "更新指定卡牌ID的价格。这里面的ID是 card_id_from_db 的值",
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
