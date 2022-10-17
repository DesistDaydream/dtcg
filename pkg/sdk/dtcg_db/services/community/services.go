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

func (c *CommunityClient) PostConvertDeck(decks string) (*models.DecksConvertPostResponse, error) {
	var decksConvertPostResponse models.DecksConvertPostResponse
	uri := "/api/community/decks/convert"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqBody: &models.DecksConvertPostRequestBody{
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
