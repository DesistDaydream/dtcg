package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/internal/database"
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type AddFlags struct {
	AddPolicy AddPolicy
	Remark    string
}

type AddPolicy struct {
	PriceChange float64
	Operator    string
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

	addProdcutCmd.Flags().Float64VarP(&addFlags.AddPolicy.PriceChange, "price-change", "c", 0, "卡牌需要变化的价格。")
	addProdcutCmd.Flags().StringVarP(&addFlags.AddPolicy.Operator, "operator", "o", "+", "卡牌价格变化的计算方式，乘法还是加法。")
	addProdcutCmd.Flags().StringVar(&addFlags.Remark, "remark", "拍之前请联系确认库存", "商品备注信息")

	return addProdcutCmd
}

// 添加商品
func addProducts(cmd *cobra.Command, args []string) {
	if productsFlags.SetPrefix == nil && productsFlags.CardVersionID == 0 {
		logrus.Fatalln("请指定要更新的卡牌，可以使用 dtcg_cli card-set list 子命令获取卡包名称；或者直接指定卡牌的 card_version_id。")
	}

	// 生成待处理的卡牌信息
	cards, err := GenNeedHandleCards()
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}

	genNeedAddProducts(cards, addFlags.AddPolicy.PriceChange)
}

// 生成待处理的商品信息
func genNeedAddProducts(cards *dbmodels.CardsPrice, priceChange float64) {
	for _, card := range cards.Data {
		// TODO: 从集换社获取一下 card.CardVersionID 是否已上架。只上架那些还没有上架的卡牌。但是每个卡牌都要向集换社发一个请求，这样是不是没必要？有必要进行这种判断吗？~

		cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(card.CardVersionID))
		if err != nil {
			logrus.Errorf("获取 %v 价格失败：%v", card.ScName, err)
			failCount++
			failList = append(failList, fmt.Sprintf("%v-%v", cardPrice.CardVersionID, cardPrice.ScName))
			continue
		}

		var newPrice string

		if addFlags.AddPolicy.Operator == "*" {
			newPrice = fmt.Sprintf("%.2f", cardPrice.AvgPrice*priceChange)
		} else if addFlags.AddPolicy.Operator == "+" {
			newPrice = fmt.Sprintf("%.2f", cardPrice.AvgPrice+priceChange)
		}

		logrus.WithFields(logrus.Fields{
			"原始价格": cardPrice.AvgPrice,
			"更新价格": newPrice,
			"调整价格": fmt.Sprintf("%v %v", addFlags.AddPolicy.Operator, priceChange),
		}).Infof("上架前检查【%v】【%v %v】卡牌", card.AlternativeArt, card.Serial, card.ScName)

		if productsFlags.isRealRun {
			addRun(cardPrice, fmt.Sprint(card.CardVersionID), newPrice)
		}
	}

	logrus.WithFields(logrus.Fields{
		"总数": cards.Count,
		"成功": successCount,
		"失败": failCount,
	}).Infof("上架结果")

	if len(failList) > 0 {
		logrus.Errorf("%v", failList)
	}
}

func addRun(cardPrice *dbmodels.CardPrice, cardVersionID string, newPrice string) {
	// 开始上架
	resp, err := handler.H.JhsServices.Products.Add(&models.ProductsAddReqBody{
		AuthenticatorID:         "",
		Grading:                 "",
		CardVersionID:           cardVersionID,
		Condition:               "1",
		GameKey:                 "dgm",
		Price:                   newPrice,
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
