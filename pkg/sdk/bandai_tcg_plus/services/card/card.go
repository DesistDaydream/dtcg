package card

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/bandai_tcg_plus/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/bandai_tcg_plus/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/utils"
)

type CardClient struct {
	client *core.Client
}

func NewCardClient(client *core.Client) *CardClient {
	return &CardClient{
		client: client,
	}
}

// 显示卡牌元信息
func (c *CardClient) GetCardMetadata(gameTitleID string) (*models.CardMetadataGetResp, error) {
	var cardListResp models.CardMetadataGetResp
	uri := "/api/user/card"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.CardMetadataGetReqQuery{
			GameTitleID:  gameTitleID,
			LanguageCode: "EN",
		}),
		ReqBody: nil,
	}

	err := c.client.Request(uri, &cardListResp, reqOpts)
	if err != nil {
		return nil, err
	}
	return &cardListResp, nil
}

// 列出卡牌
func (c *CardClient) GetCardList(cardSet string) (*models.CardListResp, error) {
	var cardListResp models.CardListResp
	uri := "/api/user/card/list"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.CardListReqQuery{
			CardSet:     cardSet,
			GameTitleID: "2",
			Limit:       "400",
			Offset:      "0",
		}),
		ReqBody: nil,
	}

	err := c.client.Request(uri, &cardListResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &cardListResp, nil
}
