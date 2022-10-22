package market

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
)

type MarketClient struct {
	client *core.Client
}

func NewMarketClient(client *core.Client) *MarketClient {
	return &MarketClient{
		client: client,
	}
}

func (p *MarketClient) GetProductSellers(cardVersionID string) (*models.ProductSellersGetResp, error) {
	var productSellers models.ProductSellersGetResp
	uri := "/api/market/card-versions/products"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.ProductSellersGetReqQuery{
			CardVersionID: cardVersionID,
			Condition:     "1",
			GameKey:       "dgm",
			Page:          "1",
		}),
	}

	err := p.client.Request(uri, &productSellers, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productSellers, nil
}
