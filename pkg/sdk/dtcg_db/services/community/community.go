package community

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/utils"
)

type CommunityClient struct {
	client *core.Client
}

func NewCommunityClient(client *core.Client) *CommunityClient {
	return &CommunityClient{
		client: client,
	}
}

// 搜索卡组
func (c *CommunityClient) PostDeckSearch(reqBody *models.DeckSearchReqBody, reqQuery *models.DeckSearchReqQuery) (*models.DeckSearchPostResp, error) {
	var deckSearchResp models.DeckSearchPostResp
	uri := "/api/community/deck/search"

	reqOpts := &core.RequestOption{
		Method:   "POST",
		ReqQuery: utils.StructToMapStr(reqQuery),
		ReqBody:  reqBody,
	}

	err := c.client.Request(uri, &deckSearchResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &deckSearchResp, nil
}

// 导入卡组
func (c *CommunityClient) PostDeckConvert(decks string) (*models.DeckConvertPostResp, error) {
	var decksConvertPostResponse models.DeckConvertPostResp
	uri := "/api/community/decks/convert"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqBody: &models.DeckConvertPostReqBody{
			Deck:  decks,
			Envir: "chs",
		},
	}

	err := c.client.Request(uri, &decksConvertPostResponse, reqOpts)
	if err != nil {
		return nil, err
	}

	return &decksConvertPostResponse, nil
}

// 获取卡组广场中的卡组
func (c *CommunityClient) GetDeck(deckHID string) (*models.DeckGetResp, error) {
	var deckGetResp models.DeckGetResp

	uri := fmt.Sprintf("/api/community/decks/%s", deckHID)

	reqOpts := &core.RequestOption{
		Method: "GET",
	}

	err := c.client.Request(uri, &deckGetResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &deckGetResp, nil
}

// 获取云卡组详情(云卡组是个人页面中创建的卡组)
func (c *CommunityClient) GetDeckCloud(deckID string) (*models.CloudDeckGetResp, error) {
	var cloudDeckGetResp models.CloudDeckGetResp
	uri := fmt.Sprintf("/api/community/cloud_deck/%s", deckID)

	reqOpts := &core.RequestOption{
		Method: "GET",
	}

	err := c.client.Request(uri, &cloudDeckGetResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &cloudDeckGetResp, nil
}
