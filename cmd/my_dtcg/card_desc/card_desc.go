package carddesc

import (
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cardDescCmd := &cobra.Command{
		Use:   "card-desc",
		Short: "控制卡片描述信息",
		// PersistentPreRun: vpcPersistentPreRun,
	}

	cardDescCmd.AddCommand(
		AddCardDescCommand(),
	)

	return cardDescCmd
}
