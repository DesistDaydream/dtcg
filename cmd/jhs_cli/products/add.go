package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/internal/database"
	databasemodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type AddFlags struct {
	SetPrefix []string
	isAdd     bool
	AddPolicy AddPolicy
}

type AddPolicy struct {
	PriceRange  []float64
	PriceChange float64
	isArt       string
}

var addFlags AddFlags

func AddCommand() *cobra.Command {
	long := `
根据策略添加商品。
比如：
  jhs_cli products add -s BTC-03 -r 0,2.99 表示将所有价格在 0-2.99 之间卡牌的价格不增加，以集换价售卖。
  jhs_cli products add -s BTC-03 -r 3,9.99 -c 0.5 --art="否" 表示将所有价格在 3-9.99 之间的非异画卡牌的价格增加 0.5 元。
  jhs_cli products add -s BTC-03 -r 10,50 -c 5 --art="是" 表示将所有价格在 0-1000 之间的异画卡牌的价格增加 5 元。
  jhs_cli products add -s BTC-03 -r 50.01,1000 -c 10 表示将所有价格在 50.01-1000 之间的异画卡牌的价格增加 10 元。
`
	addCardSetCmd := &cobra.Command{
		Use:   "add",
		Short: "添加商品",
		Long:  long,
		Run:   addProducts,
	}

	addCardSetCmd.Flags().StringSliceVarP(&addFlags.SetPrefix, "sets-name", "s", nil, "要上架哪些卡包的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	addCardSetCmd.Flags().BoolVarP(&addFlags.isAdd, "yes", "y", false, "是否真实更新卡牌信息，默认值只检查更新目标并列出将要调整的价格。")
	addCardSetCmd.Flags().Float64SliceVarP(&addFlags.AddPolicy.PriceRange, "price-range", "r", nil, "更新策略，卡牌价格区间。")
	addCardSetCmd.Flags().Float64VarP(&addFlags.AddPolicy.PriceChange, "price-change", "c", 0, "卡牌需要变化的价格。")
	addCardSetCmd.Flags().StringVar(&addFlags.AddPolicy.isArt, "art", "", "是否更新异画，可用的值有两个：是、否。空值为更新所有卡牌")

	return addCardSetCmd
}

// 添加商品
func addProducts(cmd *cobra.Command, args []string) {
	if addFlags.SetPrefix == nil {
		logrus.Error("请指定要添加的卡牌集合，使用 dtcg_cli card-set list 子命令获取卡包名称。")
		return
	}

	if addFlags.AddPolicy.PriceRange == nil {
		logrus.Error("请指定要添加的卡牌价格区间。比如 -r 0,2.99 表示将所有价格在 0-2.99 之间卡牌。")
		return
	}

	genNeedAddProducts(addFlags.AddPolicy.PriceRange, addFlags.AddPolicy.isArt, addFlags.AddPolicy.PriceChange)
}

// 生成需要添加的商品信息
func genNeedAddProducts(avgPriceRange []float64, alternativeArt string, priceChange float64) {
	var (
		cards *databasemodels.CardsPrice
		err   error
	)

	// 生成需要更新的卡牌信息
	cards, err = database.GetCardPriceByCondition(300, 1, &databasemodels.CardPriceQuery{
		SetsPrefix:     addFlags.SetPrefix,
		AlternativeArt: alternativeArt,
		MinPriceRange:  "",
		AvgPriceRange:  fmt.Sprintf("%v-%v", avgPriceRange[0], avgPriceRange[1]),
	})
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}

	logrus.Infof("%v 价格区间中共有 %v 张卡牌需要添加", avgPriceRange, len(cards.Data))

	for _, card := range cards.Data {
		// TODO: 从集换社获取一下 card.CardVersionID 是否已上架。只上架那些还没有上架的卡牌。但是每个卡牌都要向集换社发一个请求，这样是不是没必要？有必要进行这种判断吗？~

		var price string

		cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(card.CardVersionID))
		if err != nil {
			logrus.Errorln("获取卡牌价格详情失败", err)
		}

		if cardPrice.AvgPrice == 0 {
			price = fmt.Sprint(cardPrice.MinPrice + priceChange)
		} else {
			price = fmt.Sprint(cardPrice.AvgPrice + priceChange)
		}

		logrus.WithFields(logrus.Fields{
			"原始价格": cardPrice.AvgPrice,
			"上架价格": cardPrice.AvgPrice + priceChange,
		}).Infof("将要上架的【%v】【%v %v】调整 %v 元", card.AlternativeArt, card.Serial, card.ScName, priceChange)

		if addFlags.isAdd {
			addRun(cardPrice, fmt.Sprint(card.CardVersionID), price)
		}
	}
}

func addRun(cardPrice *databasemodels.CardPrice, cardVersionID string, price string) {
	// 开始上架
	resp, err := handler.H.JhsServices.Products.Add(&models.ProductsAddReqBody{
		CardVersionID:        cardVersionID,
		Price:                price,
		Quantity:             "4",
		Condition:            "1",
		Remark:               "",
		GameKey:              "dgm",
		UserCardVersionImage: cardPrice.ImageUrl,
	})
	if err != nil {
		logrus.Errorf("%v 上架失败：%v", cardPrice.ScName, err)
	} else {
		logrus.Infof("%v 上架成功：%v", cardPrice.ScName, resp)
	}
}
