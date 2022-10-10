package services

import (
	"encoding/json"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/models"
)

type SearchClient struct {
	client *core.Client
}

func NewSearchClient(client *core.Client) *SearchClient {
	return &SearchClient{
		client: client,
	}
}

func (s *SearchClient) PostDeckSearch(reqBody *models.DeckSearchRequestBody, reqQuery *models.SearchReqQuery) (*models.DeckSearchPostResponse, error) {
	var deckSearchResp models.DeckSearchPostResponse
	uri := "/api/community/deck/search"

	reqOpts := &core.RequestOption{
		Method:   "POST",
		ReqQuery: core.StructToMapStr(reqQuery),
		ReqBody:  reqBody,
	}

	respBody, err := s.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &deckSearchResp)
	if err != nil {
		return nil, err
	}

	return &deckSearchResp, nil
}

func (s *SearchClient) PostCardSearch(cardPack int) (*models.CardSearchPostResponse, error) {
	var deckSearchResp models.CardSearchPostResponse
	uri := "/api/cdb/cards/search"

	reqBody := &models.CardSearchRequestBody{
		CardPack:   cardPack,
		ClassInput: false,
		Keyword:    "",
		Language:   "chs",
		OrderType:  "default",
		TagsLogic:  "or",
		Type:       "",
	}

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqQuery: core.StructToMapStr(&models.SearchReqQuery{
			Limit: "20",
			Page:  "1",
		}),
		ReqBody: reqBody,
	}

	body, err := s.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &deckSearchResp)
	if err != nil {
		return nil, err
	}

	return &deckSearchResp, nil
}

// func (s *SearchClient) GetSeries() (*models.SeriesGetResp, error) {
// 	var resp models.SeriesGetResp
// 	uri := "/api/cdb/series"

// 	body, err := s.client.Request(uri, reqOpts)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = json.Unmarshal(body, &deckSearchResp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &resp, nil
// }
