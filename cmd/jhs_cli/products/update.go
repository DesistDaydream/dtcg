package products

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	SellerUserID string   // 集换社卖家 ID
	SetPrefix    []string // 要更新哪些卡牌集合中的卡
	CurSaleState string   // 当前商品的售卖状态
	ExpSaleState string   // 期望商品变成哪种售卖状态
}

var updateFlags UpdateFlags

func UpdateCommand() *cobra.Command {

	updateProductsCmd := &cobra.Command{
		Use:   "update",
		Short: "更新商品",
		// Run:   updatePrice,
		PersistentPreRun: updatePersistentPreRun,
	}

	updateProductsCmd.PersistentFlags().StringVarP(&updateFlags.SellerUserID, "seller-user-id", "i", "934972", "卖家用户ID。")
	updateProductsCmd.PersistentFlags().StringSliceVarP(&updateFlags.SetPrefix, "sets-name", "s", nil, "要上架哪些卡包的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	updateProductsCmd.PersistentFlags().StringVar(&updateFlags.CurSaleState, "cur-sale-state", "1", "当前售卖状态。即获取什么状态的商品。1: 售卖。0: 下架")
	updateProductsCmd.PersistentFlags().StringVar(&updateFlags.ExpSaleState, "exp-sale-state", "1", "期望的售卖状态。")

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
