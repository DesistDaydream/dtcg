package deckprice

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community"
)

// type fanxing interface {
// 	m2.Egg | m2.Main
// }

// func genData[T fanxing](resp models.PostDeckPriceResponse, card T) {

// }

// func genData(card *communityModels.CardInfo, resp *v1Models.PostDeckPriceResponse) error {
// 	cardPrice, err := database.GetCardPrice(fmt.Sprint(card.Cards.CardID))
// 	if err != nil {
// 		return fmt.Errorf("获取价格失败")
// 	}

// 	minPrice := cardPrice.MinPrice * float64(card.Number)
// 	avgPrice := cardPrice.AvgPrice * float64(card.Number)

// 	resp.Data = append(resp.Data, v1Models.MutCardPrice{
// 		Count:          int(card.Number),
// 		Serial:         cardPrice.Serial,
// 		ScName:         cardPrice.ScName,
// 		AlternativeArt: cardPrice.AlternativeArt,
// 		MinPrice:       minPrice,
// 		AvgPrice:       avgPrice,
// 	})

// 	resp.MinPrice = resp.MinPrice + minPrice
// 	resp.AvgPrice = resp.AvgPrice + avgPrice

// 	return nil
// }

func GetResp(req *models.PostDeckPriceRequest) (*models.PostDeckPriceResponse, error) {
	var (
		resp        models.PostDeckPriceResponse
		allMinPrice float64
		allAvgPrice float64
	)

	client := community.NewCommunityClient(core.NewClient(""))
	decks, err := client.PostDeckConvert(req.Deck)
	if err != nil {
		return nil, fmt.Errorf("从 dtcg db 网站获取卡组详情失败: %v", err)
	}

	// TODO: 假设 Eggs 和 Main 是两种类型的话，怎么用泛型？
	for _, card := range decks.Data.DeckInfo.Eggs {
		cardPrice, err := database.GetCardPrice(fmt.Sprint(card.Cards.CardID))
		if err != nil {
			return nil, fmt.Errorf("获取价格失败")
		}

		minPrice := cardPrice.MinPrice * float64(card.Number)
		avgPrice := cardPrice.AvgPrice * float64(card.Number)

		resp.Data = append(resp.Data, models.MutCardPrice{
			Count:          int(card.Number),
			Serial:         cardPrice.Serial,
			ScName:         cardPrice.ScName,
			AlternativeArt: cardPrice.AlternativeArt,
			MinPrice:       fmt.Sprintf("%.2f", minPrice),
			AvgPrice:       fmt.Sprintf("%.2f", avgPrice),
		})

		allMinPrice = allMinPrice + minPrice
		allAvgPrice = allAvgPrice + avgPrice
	}

	// TODO: 假设 Eggs 和 Main 是两种类型的话，怎么用泛型？
	for _, card := range decks.Data.DeckInfo.Main {
		cardPrice, err := database.GetCardPrice(fmt.Sprint(card.Cards.CardID))
		if err != nil {
			return nil, fmt.Errorf("获取价格失败")
		}

		minPrice := cardPrice.MinPrice * float64(card.Number)
		avgPrice := cardPrice.AvgPrice * float64(card.Number)

		resp.Data = append(resp.Data, models.MutCardPrice{
			Count:          int(card.Number),
			Serial:         cardPrice.Serial,
			ScName:         cardPrice.ScName,
			AlternativeArt: cardPrice.AlternativeArt,
			MinPrice:       fmt.Sprintf("%.2f", minPrice),
			AvgPrice:       fmt.Sprintf("%.2f", avgPrice),
		})

		allMinPrice = allMinPrice + minPrice
		allAvgPrice = allAvgPrice + avgPrice
	}

	resp.MinPrice = fmt.Sprintf("%.2f", allMinPrice)
	resp.AvgPrice = fmt.Sprintf("%.2f", allAvgPrice)

	return &resp, nil
}

func GetRespWithID(req *models.PostDeckPriceWithIDReq) (*models.PostDeckPriceResponse, error) {
	var (
		resp models.PostDeckPriceResponse
		// allMinPrice float64
		// allAvgPrice float64
	)

	return &resp, nil
}
