package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/en/models"
)

func GetCardList(r *models.CardListReq) (*models.CardList, error) {
	url := "https://api.bandai-tcg-plus.com/api/user/card/list"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if r.CardSet != "" {
		q.Add("card_set[]", r.CardSet)
	}
	q.Add("game_title_id", r.GameTitleID)
	q.Add("limit", r.Limit)
	q.Add("offset", r.Offset)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cardList *models.CardList
	err = json.Unmarshal(body, &cardList)
	if err != nil {
		return nil, err
	}

	return cardList, nil

}
