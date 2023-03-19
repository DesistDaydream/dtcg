package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ProductsFlags struct {
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

	productsCmd.PersistentFlags().BoolVarP(&productsFlags.isRealRun, "yes", "y", false, "是否真正执行处理商品的逻辑，默认值只检查商品的增删改查而不真的去调用集换社接口。")

	return productsCmd
}

// 生成待处理的卡牌信息
func GenNeedHandleCards(avgPriceRange []float64, alternativeArt string) (*dbmodels.CardsPrice, error) {
	var (
		cards *dbmodels.CardsPrice
		err   error
	)
	cards, err = database.GetCardPriceByCondition(300, 1, &dbmodels.CardPriceQuery{
		SetsPrefix:     updateFlags.SetPrefix,
		AlternativeArt: alternativeArt,
		MinPriceRange:  "",
		AvgPriceRange:  fmt.Sprintf("%v-%v", avgPriceRange[0], avgPriceRange[1]),
	})
	if err != nil {
		return nil, err
	}

	logrus.Infof("%v 价格区间中共有 %v 张卡牌需要更新", avgPriceRange, len(cards.Data))

	return cards, nil
}
