package orders

import "github.com/spf13/cobra"

func CreateCommand() *cobra.Command {
	ordersCmd := &cobra.Command{
		Use:   "orders",
		Short: "管理订单信息",
	}

	ordersCmd.AddCommand(
		GetAllOrderPriceCommand(),
	)

	return ordersCmd
}
