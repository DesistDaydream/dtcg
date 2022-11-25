package deckprice

import (
	"encoding/json"
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

func GetResp(req *models.PostDeckPriceWithJSONReqBody) (*models.PostDeckPriceResp, error) {
	var (
		resp        models.PostDeckPriceResp
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

		resp.Data = append(resp.Data, models.PostDeckPriceRespData{
			Count:          int(card.Number),
			Serial:         cardPrice.Serial,
			ScName:         cardPrice.ScName,
			AlternativeArt: cardPrice.AlternativeArt,
			MinPrice:       fmt.Sprintf("%.2f", minPrice),
			AvgPrice:       fmt.Sprintf("%.2f", avgPrice),
			MinUnitPrice:   fmt.Sprintf("%.2f", cardPrice.MinPrice),
			AvgUnitPrice:   fmt.Sprintf("%.2f", cardPrice.AvgPrice),
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

		resp.Data = append(resp.Data, models.PostDeckPriceRespData{
			Count:          int(card.Number),
			Serial:         cardPrice.Serial,
			ScName:         cardPrice.ScName,
			AlternativeArt: cardPrice.AlternativeArt,
			MinPrice:       fmt.Sprintf("%.2f", minPrice),
			AvgPrice:       fmt.Sprintf("%.2f", avgPrice),
			MinUnitPrice:   fmt.Sprintf("%.2f", cardPrice.MinPrice),
			AvgUnitPrice:   fmt.Sprintf("%.2f", cardPrice.AvgPrice),
		})

		allMinPrice = allMinPrice + minPrice
		allAvgPrice = allAvgPrice + avgPrice
	}

	resp.MinPrice = fmt.Sprintf("%.2f", allMinPrice)
	resp.AvgPrice = fmt.Sprintf("%.2f", allAvgPrice)

	return &resp, nil
}

func transform(ids string) (*models.PostDeckPriceWithIDReqTransform, error) {
	var (
		cards      models.PostDeckPriceWithIDReqTransform
		cardsID    []string
		cardsIndex int
		count      int
	)

	json.Unmarshal([]byte(ids), &cardsID)

	for i := 0; i < len(cardsID); i++ {
		// 第一条数据不需要判断
		if i == 0 {
			count = 1
			cardDesc, _ := database.GetCardDescByCardIDFromDB(cardsID[i])
			cards.CardsInfo = append(cards.CardsInfo, models.CardInfo{
				Count:        count,
				CardIDFromDB: cardsID[i],
				ScName:       cardDesc.ScName,
				Serial:       cardDesc.Serial,
			})
		} else if cardsID[i] != cardsID[i-1] {
			count = 1
			cardDesc, _ := database.GetCardDescByCardIDFromDB(cardsID[i])
			cards.CardsInfo = append(cards.CardsInfo, models.CardInfo{
				Count:        count,
				CardIDFromDB: cardsID[i],
				ScName:       cardDesc.ScName,
				Serial:       cardDesc.Serial,
			})
			cardsIndex++
		} else {
			count++
			cards.CardsInfo[cardsIndex].Count = count
		}
	}

	return &cards, nil
}

func GetRespWithID(req *models.PostDeckPriceWithIDReq) (*models.PostDeckPriceResp, error) {
	var (
		resp        models.PostDeckPriceResp
		allMinPrice float64
		allAvgPrice float64
	)

	cardsInfo, _ := transform(req.IDs)

	for _, card := range cardsInfo.CardsInfo {
		cardPrice, err := database.GetCardPrice(fmt.Sprint(card.CardIDFromDB))
		if err != nil {
			return nil, fmt.Errorf("获取价格失败")
		}

		minPrice := cardPrice.MinPrice * float64(card.Count)
		avgPrice := cardPrice.AvgPrice * float64(card.Count)

		resp.Data = append(resp.Data, models.PostDeckPriceRespData{
			Count:          int(card.Count),
			Serial:         cardPrice.Serial,
			ScName:         cardPrice.ScName,
			AlternativeArt: cardPrice.AlternativeArt,
			MinPrice:       fmt.Sprintf("%.2f", minPrice),
			AvgPrice:       fmt.Sprintf("%.2f", avgPrice),
			MinUnitPrice:   fmt.Sprintf("%.2f", cardPrice.MinPrice),
			AvgUnitPrice:   fmt.Sprintf("%.2f", cardPrice.AvgPrice),
		})

		allMinPrice = allMinPrice + minPrice
		allAvgPrice = allAvgPrice + avgPrice
	}

	resp.MinPrice = fmt.Sprintf("%.2f", allMinPrice)
	resp.AvgPrice = fmt.Sprintf("%.2f", allAvgPrice)

	return &resp, nil
}
