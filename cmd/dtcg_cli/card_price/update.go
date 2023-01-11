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
	UpdateAllField bool
	FromWhere      string
}

type UpdateMethod struct {
	SetPrefix     []string
	StartAt       int
	CardIDFromDBs []int
	UpdateNoImage bool
}

var updateFlags UpdateFlags

func UpdateCardPriceCommand() *cobra.Command {
	updateCardPriceCmd := &cobra.Command{
		Use:              "update",
		Short:            "更新卡牌价格",
		PersistentPreRun: cardPricePersistentPreRun,
		Run:              updateCardPrice,
	}

	updateCardPriceCmd.Flags().StringSliceVar(&updateFlags.UpdateMethod.SetPrefix, "sets-name", nil, "更新哪些卡包的价格，使用 card-set list 子命令获取卡包名称。若不指定则更新所有")
	updateCardPriceCmd.Flags().IntVar(&updateFlags.UpdateMethod.StartAt, "start-at", 0, "从哪张卡牌开始更新。值为card_id_from_db")
	updateCardPriceCmd.Flags().IntSliceVar(&updateFlags.UpdateMethod.CardIDFromDBs, "id", nil, "更新哪几张张卡牌的价格")
	updateCardPriceCmd.Flags().BoolVarP(&updateFlags.UpdateMethod.UpdateNoImage, "no-image", "n", false, "是否只更新没有卡图的卡牌价格")
	updateCardPriceCmd.Flags().BoolVarP(&updateFlags.UpdateAllField, "all-field", "a", false, "是否更新卡牌价格的全部字段")
	updateCardPriceCmd.Flags().StringVarP(&updateFlags.FromWhere, "from-where", "w", "jhs", "从哪里获取卡牌价格，目前支持 dtcgdb 和 jhs。注意：从 dtcgdb 更新将会导致 card_version_id 重置为 0")

	updateCardPriceCmd.AddCommand(
		UpdateNoImgCardPriceCommand(),
	)

	return updateCardPriceCmd
}

func updateCardPrice(cmd *cobra.Command, args []string) {
	cardsDesc, err := database.ListCardDesc()
	if err != nil {
		logrus.Errorf("获取全部卡牌描述失败: %v", err)
	}

	// ##### 多种更新卡牌价格的方式 #####
	switch expression := updateFlags.UpdateMethod; {
	// 1. 更新指定卡牌ID的价格
	case expression.CardIDFromDBs != nil:
		for _, cardIDFromDB := range updateFlags.UpdateMethod.CardIDFromDBs {
			cardDesc, _ := database.GetCardDescByCardIDFromDB(fmt.Sprint(cardIDFromDB))
			updateRun(cardDesc)
		}
	// 2. 更新指定卡牌集合的卡牌价格
	case expression.SetPrefix != nil:
		for _, cardDesc := range cardsDesc.Data {
			updateCardPriceBaseonCardSet(cardDesc, updateFlags.UpdateMethod.SetPrefix)
		}
	// 3. 从指定的卡牌ID开始更新
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
	// 获取卡牌的集换社ID与价格
	switch updateFlags.FromWhere {
	case "dtcgdb":
		cardVersionID, minPrice, avgPrice = GetPriceFromDtcgdb(cardDesc)
	case "jhs":
		cardVersionID, minPrice, avgPrice = GetPriceFromJhs(cardDesc)
	}

	// 从 DTCG DB 处获取到的 card_version_id 可能为 0
	// 为了防止 card_version_id 被重置为 0，当 card_version_id 为 0 时，不再更新卡牌价格，直接返回
	if cardVersionID == 0 {
		return
	}

	// 若 card_version_id 不为 0 时，更新卡牌价格
	// 可以更新全部字段，也可以只更新价格
	if updateFlags.UpdateAllField {
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
