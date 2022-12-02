package cardprice

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	UpdateMethod   UpdateMethod
	UpdateImageURL bool
	FromWhere      string
}

type UpdateMethod struct {
	SetPrefix     []string
	StartAt       int
	CardIDFromDBs []int
}

var updateFlags UpdateFlags

func UpdateCardPriceCommand() *cobra.Command {
	updateCardPriceCmd := &cobra.Command{
		Use:   "update",
		Short: "更新卡牌价格",
		Run:   updateCardPrice,
	}

	updateCardPriceCmd.Flags().StringSliceVar(&updateFlags.UpdateMethod.SetPrefix, "sets-name", nil, "更新哪些卡包的价格，使用 card-set list 子命令获取卡包名称。若不指定则更新所有")
	updateCardPriceCmd.Flags().IntVar(&updateFlags.UpdateMethod.StartAt, "start-at", 0, "从哪张卡牌开始更新")
	updateCardPriceCmd.Flags().IntSliceVar(&updateFlags.UpdateMethod.CardIDFromDBs, "id", nil, "更新哪几张张卡牌的价格")
	updateCardPriceCmd.Flags().BoolVar(&updateFlags.UpdateImageURL, "update-image", false, "是否更新卡牌的图片URL")
	updateCardPriceCmd.Flags().StringVar(&updateFlags.FromWhere, "from-where", "dtcgdb", "从哪里获取卡牌价格，目前支持 dtcgdb 和 jhs")

	return updateCardPriceCmd
}

func updateCardPrice(cmd *cobra.Command, args []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("获取全部卡牌描述失败: %v", err)
	}

	// ##### 三种更新卡牌价格的方式 #####
	switch expression := updateFlags.UpdateMethod; {
	// 更新指定卡牌ID的价格
	case expression.CardIDFromDBs != nil:
		for _, cardIDFromDB := range updateFlags.UpdateMethod.CardIDFromDBs {
			cardDesc, _ := database.GetCardDescByCardIDFromDB(fmt.Sprint(cardIDFromDB))
			updateRun(cardDesc)
		}
	// 更新指定卡牌集合的卡牌价格
	case expression.SetPrefix != nil:
		for _, cardDesc := range cardsDesc.Data {
			updateCardPriceBaseonCardSet(cardDesc, updateFlags.UpdateMethod.SetPrefix)
		}
	// 从指定的卡牌ID开始更新
	case expression.StartAt != 0:
		updateCardPriceBaseonStartAt(cardsDesc, updateFlags.UpdateMethod.StartAt)
	default:
		// 从头开始更新
		updateCardPriceBaseonStartAt(cardsDesc, 0)
	}
}

// 从指定的卡牌开始更新
func updateCardPriceBaseonStartAt(cardsDesc *models.CardsDesc, startAtCardIDFromDB int) {
	var startAt int
	if startAtCardIDFromDB != 0 {
		for i, cardDesc := range cardsDesc.Data {
			if cardDesc.CardIDFromDB == startAtCardIDFromDB {
				startAt = i
			}
		}
	} else {
		startAt = 0
	}

	for i := startAt; i < len(cardsDesc.Data); i++ {
		updateRun(&cardsDesc.Data[i])
	}
}

// 只更新指定卡牌集合的卡牌价格
func updateCardPriceBaseonCardSet(cardDesc models.CardDesc, setsPrefix []string) {
	for _, setPrefix := range setsPrefix {
		if cardDesc.SetPrefix == setPrefix {
			updateRun(&cardDesc)
		}
	}
}

func updateRun(cardDesc *models.CardDesc) {
	var (
		cardVersionID int
		minPrice      float64
		avgPrice      float64
	)
	switch updateFlags.FromWhere {
	case "dtcgdb":
		cardVersionID, minPrice, avgPrice = GetPriceFromDtcgdb(cardDesc)
	case "jhs":
		cardVersionID, minPrice, avgPrice = GetPriceFromJhs(cardDesc)
	}

	// if updateFlags.UpdateImageURL {
	// 	imageUrl := GetImageURL(cardVersionID)
	// 	database.UpdateCardPrice(&models.CardPrice{
	// 		CardIDFromDB:   cardDesc.CardIDFromDB,
	// 		SetID:          cardDesc.SetID,
	// 		SetPrefix:      cardDesc.SetPrefix,
	// 		Serial:         cardDesc.Serial,
	// 		ScName:         cardDesc.ScName,
	// 		AlternativeArt: cardDesc.AlternativeArt,
	// 		Rarity:         cardDesc.Rarity,
	// 		CardVersionID:  cardVersionID,
	// 		MinPrice:       minPrice,
	// 		AvgPrice:       avgPrice,
	// 		ImageUrl:       imageUrl,
	// 	}, map[string]interface{}{})
	// } else {
	// 	database.UpdateCardPrice(&models.CardPrice{
	// 		CardIDFromDB: cardDesc.CardIDFromDB,
	// 		MinPrice:     minPrice,
	// 		AvgPrice:     avgPrice,
	// 	}, map[string]interface{}{})
	// }
	if updateFlags.UpdateImageURL {
		imageUrl := GetImageURL(cardVersionID)
		database.UpdateCardPrice(&models.CardPrice{CardIDFromDB: cardDesc.CardIDFromDB}, map[string]interface{}{
			"set_id":          cardDesc.SetID,
			"set_prefix":      cardDesc.SetPrefix,
			"serial":          cardDesc.Serial,
			"sc_name":         cardDesc.ScName,
			"alternative_art": cardDesc.AlternativeArt,
			"rarity":          cardDesc.Rarity,
			"card_version_id": cardVersionID,
			"min_price":       minPrice,
			"avg_price":       avgPrice,
			"image_url":       imageUrl,
		})
	} else {
		database.UpdateCardPrice(&models.CardPrice{
			CardIDFromDB: cardDesc.CardIDFromDB,
		}, map[string]interface{}{
			"min_price": minPrice,
			"avg_price": avgPrice,
		})
	}
}
