package products

import (
	"github.com/spf13/cobra"
)

func UpdateCommand() *cobra.Command {
	UpdateProductsCmd := &cobra.Command{
		Use:              "update",
		Short:            "更新我在卖的卡片",
		Run:              updateProducts,
		PersistentPreRun: updatePersistentPreRun,
	}

	UpdateProductsCmd.AddCommand(
		UpdateImageCommand(),
	)

	return UpdateProductsCmd
}

func updatePersistentPreRun(cmd *cobra.Command, args []string) {
	// 执行父命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}
}

func updateProducts(cmd *cobra.Command, args []string) {

}
