package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/spf13/cobra"
)

func UpdateOnlyCardVersionIDCommand() *cobra.Command {
	updateCardVersionIDCmd := &cobra.Command{
		Use:   "cid",
		Short: "更新 card_version_id 的值",
		Run:   updateCardVersionID,
	}

	return updateCardVersionIDCmd
}

// 只更新数据库的 card_price 表中的 card_version_id 字段
func updateCardVersionID(cmd *cobra.Command, args []string) {
	cardVersionID := 3833
	for id := 1302; id <= 1315; id++ {
		database.DB.Exec("UPDATE `my_dtcg`.`card_prices` SET `card_version_id`=? WHERE `id`=?", cardVersionID, id).Debug()
		cardVersionID += 1
	}
}
