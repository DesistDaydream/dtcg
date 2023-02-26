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
	Remark    string
}

type AddPolicy struct {
	PriceRange  []float64
	PriceChange float64
	isArt       string
}

var (
	addFlags     AddFlags
	successCount int = 0
	failCount    int = 0
	failList     []string
)

func AddCommand() *cobra.Command {
	long := `
根据策略添加商品。
比如：
  jhs_cli products add -s BTC-03 -r 0,1000 -c 20 表示将所有价格在 0-1000 之间卡牌的价格增加 20 块售卖。
`
	addProdcutCmd := &cobra.Command{
		Use:   "add",
		Short: "添加商品",
		Long:  long,
		Run:   addProducts,
	}

	addProdcutCmd.Flags().StringSliceVarP(&addFlags.SetPrefix, "sets-name", "s", nil, "要上架哪些卡包的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	addProdcutCmd.Flags().BoolVarP(&addFlags.isAdd, "yes", "y", false, "是否真实更新卡牌信息，默认值只检查更新目标并列出将要调整的价格。")
	addProdcutCmd.Flags().Float64SliceVarP(&addFlags.AddPolicy.PriceRange, "price-range", "r", nil, "更新策略，卡牌价格区间。")
	addProdcutCmd.Flags().Float64VarP(&addFlags.AddPolicy.PriceChange, "price-change", "c", 0, "卡牌需要变化的价格。")
	addProdcutCmd.Flags().StringVar(&addFlags.AddPolicy.isArt, "art", "", "是否更新异画，可用的值有两个：是、否。空值为更新所有卡牌")
	addProdcutCmd.Flags().StringVar(&addFlags.Remark, "remark", "拍之前请联系确认库存", "商品备注信息")

	return addProdcutCmd
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

	logrus.WithFields(logrus.Fields{
		"总数": cards.Count,
		"成功": successCount,
		"失败": failCount,
	}).Infof("上架结果")
}

func addRun(cardPrice *databasemodels.CardPrice, cardVersionID string, price string) {
	// 开始上架
	resp, err := handler.H.JhsServices.Products.Add(&models.ProductsAddReqBody{
		AuthenticatorID:         "",
		Grading:                 "",
		CardVersionID:           cardVersionID,
		Condition:               "1",
		GameKey:                 "dgm",
		Price:                   price,
		ProductCardVersionImage: cardPrice.ImageUrl,
		Quantity:                "4",
		Remark:                  addFlags.Remark,
	})
	if err != nil {
		logrus.Errorf("%v 上架失败：%v", cardPrice.ScName, err)
		failCount++
		failList = append(failList, fmt.Sprintf("%v-%v", cardVersionID, cardPrice.ScName))
	} else {
		logrus.Infof("%v 上架成功：%v", cardPrice.ScName, resp)
		successCount++
	}
}
