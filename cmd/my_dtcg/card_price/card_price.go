package cardprice

import (
	"fmt"
	"strconv"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	"github.com/sirupsen/logrus"
)

func AddCardPrice() {
	cardsDesc, err := database.ListCardDescFromDtcgDB()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	for i := 390; i < len(cardsDesc.Data); i++ {
		client := services.NewSearchClient(core.NewClient(""))

		cardPrice, err := client.GetCardPrice(fmt.Sprint(cardsDesc.Data[i].CardID))
		if err != nil {
			logrus.Errorf("获取卡片价格失败: %v", err)
		}

		var f1 float64
		if len(cardPrice.Data.Products) == 0 {
			f1 = 0
		} else {
			f1, _ = strconv.ParseFloat(cardPrice.Data.Products[0].MinPrice, 64)
		}

		f2, _ := strconv.ParseFloat(cardPrice.Data.AvgPrice, 64)

		d := &database.CardPrice{
			CardID:        cardsDesc.Data[i].CardID,
			CardVersionID: int(cardPrice.Data.Products[0].CardVersionID),
			MinPrice:      f1,
			AvgPrice:      f2,
		}

		database.AddCardPirce(d)
	}

	// for _, cardDesc := range cardsDesc.Data {
	// 	client := services.NewSearchClient(core.NewClient(""))

	// 	cardPrice, err := client.GetCardPrice(fmt.Sprint(cardDesc.CardID))
	// 	if err != nil {
	// 		logrus.Errorf("获取卡片价格失败: %v", err)
	// 	}

	// 	var f1 float64
	// 	if len(cardPrice.Data.Products) == 0 {
	// 		f1 = 0
	// 	} else {
	// 		f1, _ = strconv.ParseFloat(cardPrice.Data.Products[0].MinPrice, 64)
	// 	}

	// 	f2, _ := strconv.ParseFloat(cardPrice.Data.AvgPrice, 64)

	// 	d := &database.CardPrice{
	// 		CardID:        cardDesc.CardID,
	// 		CardVersionID: int(cardPrice.Data.Products[0].CardVersionID),
	// 		MinPrice:      f1,
	// 		AvgPrice:      f2,
	// 	}

	// 	database.AddCardPirce(d)
	// }
}
