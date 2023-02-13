package products

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	SellerUserID string
	SetPrefix    []string
	isUpdate     bool
}

var updateFlags UpdateFlags

func UpdateCommand() *cobra.Command {

	updateProductsCmd := &cobra.Command{
		Use:   "update",
		Short: "更新商品",
		// Run:   updatePrice,
		PersistentPreRun: updatePersistentPreRun,
	}

	updateProductsCmd.PersistentFlags().StringVarP(&updateFlags.SellerUserID, "seller-user-id", "i", "609077", "卖家用户ID。")
	updateProductsCmd.PersistentFlags().StringSliceVarP(&updateFlags.SetPrefix, "sets-name", "s", nil, "要上架哪些卡包的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	updateProductsCmd.PersistentFlags().BoolVarP(&updateFlags.isUpdate, "yes", "y", false, "是否真实更新卡牌信息，默认值只检查更新目标并列出将要调整的价格。")

	updateProductsCmd.AddCommand(
		UpdatePriceCommand(),
		UpdateImageCommand(),
		UpdateSaleStateCommand(),
	)

	return updateProductsCmd
}

func updatePersistentPreRun(cmd *cobra.Command, args []string) {
	if updateFlags.SetPrefix == nil {
		logrus.Fatalln("请指定要更新的卡牌集合，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	}
}
