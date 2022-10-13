package cardprice

import (
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cardSetCmd := &cobra.Command{
		Use:   "card-price",
		Short: "控制卡片价格信息",
	}

	cardSetCmd.AddCommand(
		AddCardPriceCommand(),
		UpdateCardPriceCommand(),
	)

	return cardSetCmd
}
