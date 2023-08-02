package wishes

import (
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	wishesCmd := &cobra.Command{
		Use:   "wishes",
		Short: "管理我想收信息",
		// PersistentPreRun: productsPersistentPreRun,
	}

	wishesCmd.AddCommand(
		AddWishListCmd(),
		CompareCmd(),
		GetWishListPriceCmd(),
	)

	return wishesCmd
}
