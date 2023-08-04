package cdb

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/utils"
)

type CdbClient struct {
	client *core.Client
}

func NewCdbClient(client *core.Client) *CdbClient {
	return &CdbClient{
		client: client,
	}
}

// 列出卡牌集合
func (s *CdbClient) GetSeries() (*models.SeriesGetResp, error) {
	var seriesGetResp models.SeriesGetResp
	uri := "/api/cdb/series"

	reqOpts := &core.RequestOption{}

	err := s.client.Request(uri, &seriesGetResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &seriesGetResp, nil
}

// 获取卡牌集合详情(包括卡集中包含的所有卡牌)
func (s *CdbClient) GetPackage(setName string) (*models.PackageGetResp, error) {
	var packageGetResp models.PackageGetResp
	uri := fmt.Sprintf("/api/cdb/package/%s?extend_cards=1", setName)

	reqOpts := &core.RequestOption{
		Method: "GET",
	}

	err := s.client.Request(uri, &packageGetResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &packageGetResp, nil
}

// TODO: 获取卡牌上下文(整个卡牌游戏中的所有颜色、所有稀有度等等)

// 搜索卡牌
func (s *CdbClient) PostCardSearch(cardPack int, limit string, language string, isParallel string) (*models.CardSearchPostResp, error) {
	var cardSearchResp models.CardSearchPostResp
	uri := "/api/cdb/cards/search"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqQuery: utils.StructToMapStr(&models.CardSearchReqQuery{
			Limit: limit,
			Page:  "1",
		}),
		ReqBody: &models.CardSearchReqBody{
			CardPack:   cardPack,
			ClassInput: false,
			Keyword:    "",
			Language:   language,
			OrderType:  "default",
			TagsLogic:  "or",
			Type:       "",
			IsParallel: isParallel,
		},
	}

	err := s.client.Request(uri, &cardSearchResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &cardSearchResp, nil
}

// TODO: 获取卡牌详情

// 获取卡牌价格
func (s *CdbClient) GetCardPrice(cardID string) (*models.CardPriceGetResp, error) {
	var resp models.CardPriceGetResp
	uri := "/api/cdb/jhs/price"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.CardsPriceGetReq{
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
