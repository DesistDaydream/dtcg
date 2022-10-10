package services

import (
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

// 搜索卡组
func (s *SearchClient) PostDeckSearch(reqBody *models.DeckSearchRequestBody, reqQuery *models.SearchReqQuery) (*models.DeckSearchPostResponse, error) {
	var deckSearchResp models.DeckSearchPostResponse
	uri := "/api/community/deck/search"

	reqOpts := &core.RequestOption{
		Method:   "POST",
		ReqQuery: core.StructToMapStr(reqQuery),
		ReqBody:  reqBody,
	}

	err := s.client.Request(uri, &deckSearchResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &deckSearchResp, nil
}

// 搜索卡片
func (s *SearchClient) PostCardSearch(cardPack int) (*models.CardSearchPostResponse, error) {
	var cardSearchResp models.CardSearchPostResponse
	uri := "/api/cdb/cards/search"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqQuery: core.StructToMapStr(&models.SearchReqQuery{
			Limit: "20",
			Page:  "1",
		}),
		ReqBody: &models.CardSearchRequestBody{
			CardPack:   cardPack,
			ClassInput: false,
			Keyword:    "",
			Language:   "chs",
			OrderType:  "default",
			TagsLogic:  "or",
			Type:       "",
		},
	}

	err := s.client.Request(uri, &cardSearchResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &cardSearchResp, nil
}

// 获取卡包列表
func (s *SearchClient) GetSeries() (*models.SeriesGetResp, error) {
	var resp models.SeriesGetResp
	uri := "/api/cdb/series"

	reqOpts := &core.RequestOption{}

	err := s.client.Request(uri, &resp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
