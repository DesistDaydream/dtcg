package cdb

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb/models"
)

type CdbClient struct {
	client *core.Client
}

func NewCdbClient(client *core.Client) *CdbClient {
	return &CdbClient{
		client: client,
	}
}

// 搜索卡片
func (s *CdbClient) PostCardSearch(cardPack int) (*models.CardSearchPostResp, error) {
	var cardSearchResp models.CardSearchPostResp
	uri := "/api/cdb/cards/search"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqQuery: core.StructToMapStr(&models.CardSearchReqQuery{
			Limit: "300",
			Page:  "1",
		}),
		ReqBody: &models.CardSearchReqBody{
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
func (s *CdbClient) GetSeries() (*models.SeriesGetResp, error) {
	var resp models.SeriesGetResp
	uri := "/api/cdb/series"

	reqOpts := &core.RequestOption{}

	err := s.client.Request(uri, &resp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// 获取卡片价格
func (s *CdbClient) GetCardPrice(cardID string) (*models.CardPriceGetResp, error) {
	var resp models.CardPriceGetResp
	uri := "/api/cdb/jhs/price"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.CardsPriceGetReq{
			CardID: cardID,
		}),
		ReqBody: nil,
	}

	err := s.client.Request(uri, &resp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
