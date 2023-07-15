package products

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
)

type ProductsClient struct {
	client *core.Client
}

func NewProductsClient(client *core.Client) *ProductsClient {
	return &ProductsClient{
		client: client,
	}
}

// 获取我在卖的商品详情。注意：这也是整个集换社获取一个商品详情的接口
// TODO: 这个接口其实不应该算在我在卖的商品服务里面，但是暂时也不知道应该放在哪里。
// 注意：该接口只能获取到商家在售的商品，已经下架的商品无法通过该接口获取到。
func (p *ProductsClient) Get(cardVersionID string, sellerUserID string) (*models.ProductsGetResp, error) {
	var productsGetResp models.ProductsGetResp
	uri := "/api/market/products/bySellerCardVersionId"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.ProductsGetReqQuery{
			GameKey:       "dgm",
			SellerUserID:  sellerUserID,
			CardVersionID: cardVersionID,
		}),
	}

	err := p.client.Request(uri, &productsGetResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsGetResp, nil
}
