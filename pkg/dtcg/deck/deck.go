package deck

import (
	"encoding/json"
	"fmt"

	"github.com/DesistDaydream/dtcg/pkg/dtcg/handler"
	"github.com/sirupsen/logrus"
)

func GetResp(hID string) (string, error) {
	decks, err := handler.H.MoecardServices.Community.GetDeck(hID)
	if err != nil {
		logrus.Errorln(err)
	}

	var cardsID []string

	for _, card := range decks.Data.DeckInfo.Eggs {
		for i := 0; i < card.Number; i++ {
			cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
		}
	}

	for _, card := range decks.Data.DeckInfo.Main {
		for i := 0; i < card.Number; i++ {
			cardsID = append(cardsID, fmt.Sprint(card.Cards.CardID))
		}
	}

	cardsIDString, _ := json.Marshal(&cardsID)

	return string(cardsIDString), nil
}
