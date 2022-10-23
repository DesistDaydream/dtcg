package products

import (
	"os"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/sirupsen/logrus"
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

var productsClient *products.ProductsClient

func productsPersistentPreRun(cmd *cobra.Command, args []string) {
	// 执行父命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}

	// 自身执行前操作
	file, err := os.ReadFile("pkg/sdk/jihuanshe/services/token.txt")
	if err != nil {
		logrus.Fatal(err)
	}
	token := string(file)
	productsClient = products.NewProductsClient(core.NewClient(token))
}
