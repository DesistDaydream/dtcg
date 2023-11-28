package deckprice

import (
	"encoding/json"
	"fmt"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/dtcg/api/v1/models"
	"github.com/DesistDaydream/dtcg/pkg/handler"

	moecardmodels "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community/models"
)

type price struct {
	allMinPrice float64
	allAvgPrice float64
}

func calculatePrice(p *price, resp *models.PostDeckPriceResp, card *moecardmodels.CardInfo) error {
	cardPrice, err := database.GetCardPrice(fmt.Sprint(card.Cards.CardID))
	if err != nil {
		return fmt.Errorf("获取价格失败")
	}

	minPrice := cardPrice.MinPrice * float64(card.Number)
	avgPrice := cardPrice.AvgPrice * float64(card.Number)

	resp.Data = append(resp.Data, models.PostDeckPriceRespData{
		CardIDFromDB:   cardPrice.CardIDFromDB,
		Count:          int(card.Number),
		Serial:         cardPrice.Serial,
		ScName:         cardPrice.ScName,
		Rarity:         cardPrice.Rarity,
		AlternativeArt: cardPrice.AlternativeArt,
		MinPrice:       fmt.Sprintf("%.2f", minPrice),
		AvgPrice:       fmt.Sprintf("%.2f", avgPrice),
		MinUnitPrice:   fmt.Sprintf("%.2f", cardPrice.MinPrice),
		AvgUnitPrice:   fmt.Sprintf("%.2f", cardPrice.AvgPrice),
		Image:          cardPrice.ImageUrl,
	})

	p.allMinPrice = p.allMinPrice + minPrice
	p.allAvgPrice = p.allAvgPrice + avgPrice

	return nil
}

// 根据 DTCG_DB 导出的 JSON 格式卡组信息获取卡组价格
func GetDeckPriceWithJSON(req *models.PostDeckPriceWithJSONReqBody) (*models.PostDeckPriceResp, error) {
	var (
		resp models.PostDeckPriceResp
		p    price
	)

	decks, err := handler.H.MoecardServices.Community.PostDeckConvert(req.Deck)
	if err != nil {
		return nil, fmt.Errorf("从 dtcg db 网站获取卡组详情失败: %v", err)
	}

	for _, card := range decks.Data.DeckInfo.Eggs {
		err := calculatePrice(&p, &resp, &card)
		if err != nil {
			return nil, err
		}
	}

	for _, card := range decks.Data.DeckInfo.Main {
		err := calculatePrice(&p, &resp, &card)
		if err != nil {
			return nil, err
		}
	}

	resp.MinPrice = fmt.Sprintf("%.2f", p.allMinPrice)
	resp.AvgPrice = fmt.Sprintf("%.2f", p.allAvgPrice)

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

// 根据 card_id_from_db 计算卡组价格
func GenDeckPriceWithMoecardID(req *models.PostDeckPriceWithIDReq) (*models.PostDeckPriceResp, error) {
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
			Rarity:         cardPrice.Rarity,
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
