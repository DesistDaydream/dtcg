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

	// 卡牌的信息
	CardVersionID int
	SetPrefix     []string  // 要处理哪些卡集中的卡
	Rarity        []string  // 卡牌的稀有度。U、R、C、SR、SEC
	PriceRange    []float64 // 卡牌的价格范围
	IsArt         string
}

type Policy struct {
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
	productsCmd.PersistentFlags().IntVarP(&productsFlags.CardVersionID, "card-version-id", "i", 0, "卡牌在集换社中的 ID。指定该标志时，将只筛选到这一张卡牌")
	productsCmd.PersistentFlags().StringSliceVarP(&productsFlags.SetPrefix, "sets-name", "s", nil, "要处理哪些卡集的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	productsCmd.PersistentFlags().StringSliceVar(&productsFlags.Rarity, "rarity", nil, "卡牌的稀有度。")
	productsCmd.PersistentFlags().Float64SliceVarP(&productsFlags.PriceRange, "price-range", "r", nil, "卡牌价格区间。")
	productsCmd.PersistentFlags().StringVar(&productsFlags.IsArt, "art", "", "是否异画，可用的值有两个：是、否。空值为更新所有卡牌")

	return productsCmd
}

// 生成待处理的卡牌信息
func GenNeedHandleCards() (*dbmodels.CardsPrice, error) {
	var (
		cards         *dbmodels.CardsPrice
		err           error
		avgPriceRange string
	)

	if productsFlags.PriceRange != nil {
		avgPriceRange = fmt.Sprintf("%v-%v", productsFlags.PriceRange[0], productsFlags.PriceRange[1])
	}

	cards, err = database.GetCardPriceByCondition(300, 1, &dbmodels.CardPriceQuery{
		CardVersionID:  productsFlags.CardVersionID,
		SetsPrefix:     productsFlags.SetPrefix,
		Keyword:        "",
		Language:       "",
		QField:         []string{},
		Rarity:         productsFlags.Rarity,
		AlternativeArt: productsFlags.IsArt,
		MinPriceRange:  "",
		AvgPriceRange:  avgPriceRange,
	})
	if err != nil {
		return nil, err
	}

	logrus.Infof("在 %v 卡集中，%v 价格区间共有 %v 张卡牌需要更新", productsFlags.SetPrefix, productsFlags.PriceRange, len(cards.Data))

	return cards, nil
}
