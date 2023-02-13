package products

import (
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	productsCmd := &cobra.Command{
		Use:   "products",
		Short: "管理我在卖的商品信息",
		// PersistentPreRun: productsPersistentPreRun,
	}

	productsCmd.AddCommand(
		AddCommand(),
		UpdateCommand(),
		DelCommand(),
	)

	return productsCmd
}
