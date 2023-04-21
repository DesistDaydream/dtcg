package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/spf13/cobra"
)

type ProductsFlags struct {
	SetPrefix []string // 要处理哪些卡集中的卡
	isRealRun bool
}

var productsFlags ProductsFlags

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

	productsCmd.PersistentFlags().StringSliceVarP(&productsFlags.SetPrefix, "sets-name", "s", nil, "要处理哪些卡集的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	productsCmd.PersistentFlags().BoolVarP(&productsFlags.isRealRun, "yes", "y", false, "是否真正执行处理商品的逻辑，默认值只检查商品的增删改查而不真的去调用集换社接口。")

	return productsCmd
}

// 生成待处理的卡牌信息
func GenNeedHandleCards(avgPriceRange []float64, alternativeArt string, cardVersionID int) (*dbmodels.CardsPrice, error) {
	var (
		cards *dbmodels.CardsPrice
		err   error
	)
	cards, err = database.GetCardPriceByCondition(300, 1, &dbmodels.CardPriceQuery{
		CardVersionID:  cardVersionID,
		SetsPrefix:     productsFlags.SetPrefix,
		AlternativeArt: alternativeArt,
		MinPriceRange:  "",
		AvgPriceRange:  fmt.Sprintf("%v-%v", avgPriceRange[0], avgPriceRange[1]),
	})
	if err != nil {
		return nil, err
	}

	return cards, nil
}
