package community

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community/models"
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
		ReqQuery: core.StructToMapStr(reqQuery),
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
