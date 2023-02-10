package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	UpdateAllField bool
	FromWhere      string
}

var updateFlags UpdateFlags

func UpdateCardPriceCommand() *cobra.Command {
	updateCardPriceCmd := &cobra.Command{
		Use:              "update",
		Short:            "更新卡牌价格数据表",
		PersistentPreRun: cardPriceUpdatePersistentPreRun,
		Run:              updateCardPrice,
	}

	updateCardPriceCmd.PersistentFlags().BoolVarP(&updateFlags.UpdateAllField, "all-field", "a", false, "是否更新卡牌价格的全部字段")
	updateCardPriceCmd.PersistentFlags().StringVarP(&updateFlags.FromWhere, "from-where", "w", "jhs", "从哪里获取卡牌价格，目前支持 dtcgdb 和 jhs。")

	updateCardPriceCmd.AddCommand(
		UpdateByIDCommand(),
		UpdateByStartAtCommand(),
		UpdateBySetsNameCommand(),
		UpdateByNoImgCommand(),
		UpdateOnlyCardVersionIDCommand(),
	)

	return updateCardPriceCmd
}

func cardPriceUpdatePersistentPreRun(cmd *cobra.Command, args []string) {
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}
}

func updateCardPrice(cmd *cobra.Command, args []string) {

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
