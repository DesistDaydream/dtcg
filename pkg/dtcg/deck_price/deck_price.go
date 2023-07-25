package deckprice

import (
	"encoding/json"
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/handler"
)

// TODO: 如何用泛型改写？
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

func GetRespWithJSON(req *models.PostDeckPriceWithJSONReqBody) (*models.PostDeckPriceResp, error) {
	var (
		resp        models.PostDeckPriceResp
		allMinPrice float64
		allAvgPrice float64
	)

	// client := community.NewCommunityClient(core.NewClient("", 10))
	// decks, err := client.PostDeckConvert(req.Deck)
	decks, err := handler.H.MoecardServices.Community.PostDeckConvert(req.Deck)
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
			CardIDFromDB:   cardPrice.CardIDFromDB,
			Count:          int(card.Number),
			Serial:         cardPrice.Serial,
			ScName:         cardPrice.ScName,
			AlternativeArt: cardPrice.AlternativeArt,
			MinPrice:       fmt.Sprintf("%.2f", minPrice),
			AvgPrice:       fmt.Sprintf("%.2f", avgPrice),
			MinUnitPrice:   fmt.Sprintf("%.2f", cardPrice.MinPrice),
			AvgUnitPrice:   fmt.Sprintf("%.2f", cardPrice.AvgPrice),
			Image:          cardPrice.ImageUrl,
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
			CardIDFromDB:   cardPrice.CardIDFromDB,
			Count:          int(card.Number),
			Serial:         cardPrice.Serial,
			ScName:         cardPrice.ScName,
			AlternativeArt: cardPrice.AlternativeArt,
			MinPrice:       fmt.Sprintf("%.2f", minPrice),
			AvgPrice:       fmt.Sprintf("%.2f", avgPrice),
			MinUnitPrice:   fmt.Sprintf("%.2f", cardPrice.MinPrice),
			AvgUnitPrice:   fmt.Sprintf("%.2f", cardPrice.AvgPrice),
			Image:          cardPrice.ImageUrl,
		})

		allMinPrice = allMinPrice + minPrice
		allAvgPrice = allAvgPrice + avgPrice
	}

	resp.MinPrice = fmt.Sprintf("%.2f", allMinPrice)
	resp.AvgPrice = fmt.Sprintf("%.2f", allAvgPrice)

	return &resp, nil
}

// 根据从 DTCG DB 中获取到的所有卡牌的 card_id_from_db，从本地数据库查找对应的卡牌，并生成结构化的卡牌信息。
// 类似于：从 ["A","A","A","B","B"] 这种结构转换为类似 [{"number":3,"name":"A"},{"number":2,"name":"B"}] 这种。
// 这种结构化的数据对于前端或者其他地方使用起来，更加方便和友好。
func transform(ids string) (*models.PostDeckPriceWithIDReqTransform, error) {
	var (
		cards         models.PostDeckPriceWithIDReqTransform
		cardIDFromDBs []string
		cardsIndex    int
		count         int = 1
	)

	json.Unmarshal([]byte(ids), &cardIDFromDBs)

	for i := 0; i < len(cardIDFromDBs); i++ {
		// 第一条数据不需要判断
		if i == 0 {
			count = 1
			cardDesc, err := database.GetCardDescByCardIDFromDB(cardIDFromDBs[i])
			if err != nil {
				return nil, fmt.Errorf("card_id_from_db 为 %v 的卡牌在数据库中未找到，错误原因: %v", cardIDFromDBs[i], err)
			}
			cards.CardsInfo = append(cards.CardsInfo, models.CardInfo{
				Count:        count,
				CardIDFromDB: cardIDFromDBs[i],
				ScName:       cardDesc.ScName,
				Serial:       cardDesc.Serial,
			})
		} else if cardIDFromDBs[i] != cardIDFromDBs[i-1] {
			count = 1
			cardDesc, err := database.GetCardDescByCardIDFromDB(cardIDFromDBs[i])
			if err != nil {
				return nil, fmt.Errorf("card_id_from_db 为 %v 的卡牌在数据库中未找到，错误原因: %v", cardIDFromDBs[i], err)
			}
			cards.CardsInfo = append(cards.CardsInfo, models.CardInfo{
				Count:        count,
				CardIDFromDB: cardIDFromDBs[i],
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

// 根据卡牌 ID 计算卡组价格
func GetRespWithID(req *models.PostDeckPriceWithIDReq) (*models.PostDeckPriceResp, error) {
	var (
		resp        models.PostDeckPriceResp
		allMinPrice float64
		allAvgPrice float64
	)

	cardsInfo, err := transform(req.IDs)
	if err != nil {
		return nil, err
	}

	for _, card := range cardsInfo.CardsInfo {
		cardPrice, err := database.GetCardPrice(fmt.Sprint(card.CardIDFromDB))
		if err != nil {
			return nil, fmt.Errorf("获取价格失败")
		}

		minPrice := cardPrice.MinPrice * float64(card.Count)
		avgPrice := cardPrice.AvgPrice * float64(card.Count)

		resp.Data = append(resp.Data, models.PostDeckPriceRespData{
			CardIDFromDB:   cardPrice.CardIDFromDB,
			Count:          int(card.Count),
			Serial:         cardPrice.Serial,
			ScName:         cardPrice.ScName,
			AlternativeArt: cardPrice.AlternativeArt,
			MinPrice:       fmt.Sprintf("%.2f", minPrice),
			AvgPrice:       fmt.Sprintf("%.2f", avgPrice),
			MinUnitPrice:   fmt.Sprintf("%.2f", cardPrice.MinPrice),
			AvgUnitPrice:   fmt.Sprintf("%.2f", cardPrice.AvgPrice),
			Image:          cardPrice.ImageUrl,
		})

		allMinPrice = allMinPrice + minPrice
		allAvgPrice = allAvgPrice + avgPrice
	}

	resp.MinPrice = fmt.Sprintf("%.2f", allMinPrice)
	resp.AvgPrice = fmt.Sprintf("%.2f", allAvgPrice)

	return &resp, nil
}
