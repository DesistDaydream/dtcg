package wishes

import (
	"github.com/spf13/cobra"
)

type AddFlags struct {
	Test string
}

var addFlags AddFlags

func AddWishListCmd() *cobra.Command {
	long := `
从 dtcg db 的 云卡组 获取卡牌列表，添加到我想收清单中
`
	addWishListCmd := &cobra.Command{
		Use:   "add",
		Short: "从 dtcg db 的 云卡组 获取卡牌列表，添加到我想收清单中",
		Long:  long,
	}

	addWishListCmd.AddCommand(
		addFromMoecardCmd(),
		addFromWishListCmd(),
	)

	addWishListCmd.Flags().StringVar(&addFlags.Test, "testAddWishList", "", "保留标志")

	return addWishListCmd
}
