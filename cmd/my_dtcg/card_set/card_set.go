package cardset

import (
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cardSetCmd := &cobra.Command{
		Use:   "card-set",
		Short: "控制卡片集合信息",
		// PersistentPreRun: vpcPersistentPreRun,
	}

	cardSetCmd.AddCommand(
		AddCardSetCommand(),
	)

	return cardSetCmd
}
