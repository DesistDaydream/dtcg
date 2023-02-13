package products

import (
	"github.com/spf13/cobra"
)

type DelFlags struct {
}

var delFlags DelFlags

func DelCommand() *cobra.Command {
	long := `
根据策略添加商品。
比如：
  jhs_cli products add -s BTC-03 -r 0,1000 -c 20 表示将所有价格在 0-1000 之间卡牌的价格增加 20 块售卖。
`
	delProdcutCmd := &cobra.Command{
		Use:   "del",
		Short: "删除商品",
		Long:  long,
		Run:   delProducts,
	}

	// delProdcutCmd.Flags().StringSliceVarP(&delFlags.SetPrefix, "sets-name", "s", nil, "要上架哪些卡包的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")

	return delProdcutCmd
}

// 添加商品
func delProducts(cmd *cobra.Command, args []string) {

}
