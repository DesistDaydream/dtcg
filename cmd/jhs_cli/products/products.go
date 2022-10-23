package products

import (
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	productsCmd := &cobra.Command{
		Use:              "products",
		Short:            "控制卡片集合信息",
		PersistentPreRun: productsPersistentPreRun,
	}

	productsCmd.AddCommand(
		AddCommand(),
		UpdateCommand(),
	)

	return productsCmd
}

func productsPersistentPreRun(cmd *cobra.Command, args []string) {
	// 执行父命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}
}
