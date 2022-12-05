package products

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/cmd/jhs_cli/handler"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type AddFlags struct {
	SetPrefix []string
	AddPolicy AddPolicy
}

type AddPolicy struct {
	// 根据最大价格上架
	ControlAddPrice bool
	MaxPrice        float64
}

var addFlags AddFlags

func AddCommand() *cobra.Command {
	addCardSetCmd := &cobra.Command{
		Use:   "add",
		Short: "添加商品",
		Run:   addProducts,
	}

	addCardSetCmd.Flags().StringSliceVarP(&addFlags.SetPrefix, "sets-name", "s", nil, "要上架哪些卡包的卡牌，使用 dtcg_cli card-set list 子命令获取卡包名称。")
	addCardSetCmd.Flags().BoolVarP(&addFlags.AddPolicy.ControlAddPrice, "control-add-price", "c", false, "是否控制上架价格，如果为 true，需要指定 --max-price 标志。")
	addCardSetCmd.Flags().Float64VarP(&addFlags.AddPolicy.MaxPrice, "max-price", "m", 100, "最高价格，低于该价格的卡牌才会被上架。")

	return addCardSetCmd
}

// 生成需要上架的卡牌信息
func genNeedAddedCardInfo(setsPrefix []string) ([]string, error) {
	var cardVersionIDs []string

	for _, setPrefix := range setsPrefix {
		// 根据 ser_prefix 获取 card_version_id
		cardsPrice, err := database.GetCardPriceWhereSetPrefix(setPrefix)
		if err != nil {
			return nil, err
		}

		// 上架策略
		switch policy := addFlags.AddPolicy; {
		// 是否根据价格控制上架
		case policy.ControlAddPrice:
			for _, cardPrice := range cardsPrice.Data {
				if cardPrice.AvgPrice <= addFlags.AddPolicy.MaxPrice {
					cardVersionIDs = append(cardVersionIDs, fmt.Sprint(cardPrice.CardVersionID))
				}
			}
		default:
			for _, cardPrice := range cardsPrice.Data {
				cardVersionIDs = append(cardVersionIDs, fmt.Sprint(cardPrice.CardVersionID))
			}
		}
	}

	logrus.Debugf("当前需要上架 %v 张卡牌：%v", len(cardVersionIDs), cardVersionIDs)

	return cardVersionIDs, nil
}

// 添加商品
func addProducts(cmd *cobra.Command, args []string) {
	if addFlags.SetPrefix == nil {
		logrus.Error("请指定要上架的卡牌集合，使用 dtcg_cli card-set list 子命令获取卡包名称。")
		return
	}

	cards, err := genNeedAddedCardInfo(addFlags.SetPrefix)
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}

	fmt.Println(cards)

	for _, cardVersionID := range cards {
		var price string

		cardPrice, err := database.GetCardPriceWhereCardVersionID(cardVersionID)
		if err != nil {
			logrus.Errorln("获取卡牌价格详情失败", err)
		}

		if cardPrice.AvgPrice == 0 {
			price = fmt.Sprint(cardPrice.MinPrice + float64(2))
		} else {
			price = fmt.Sprint(cardPrice.AvgPrice + float64(2))
		}

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

		// resp, err := client.Add(cardModelToCardVersionID[rows[i][0]], rows[i][1], rows[i][2])
		if err != nil {
			logrus.Errorf("%v 上架失败：%v", cardPrice.ScName, err)
		} else {
			logrus.Infof("%v 上架成功：%v", cardPrice.ScName, resp)
		}
	}
}
