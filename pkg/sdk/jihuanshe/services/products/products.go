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

// 获取我在卖的商品列表
func (p *ProductsClient) List(page string) (*models.ProductsListResponse, error) {
	var productsResp models.ProductsListResponse

	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.ProductsListRequestQuery{
			GameKey:    "dgm",
			GameSubKey: "sc",
			OnSale:     "1",
			Page:       page,
			Token:      p.client.Token,
		}),
	}

	err := p.client.Request(uri, &productsResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsResp, nil
}

// 获取我在卖的商品详情。注意：这也是整个集换社获取一个商品详情的接口
func (p *ProductsClient) Get(cardVersionID string) (*models.ProductsGetResponse, error) {
	var productsGetResp models.ProductsGetResponse
	uri := "/api/market/products/bySellerCardVersionId"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.ProductsGetRequestQuery{
			GameKey:       "dgm",
			SellerUserID:  "609077",
			CardVersionID: cardVersionID,
			// Token:         p.client.Token,
		}),
	}

	err := p.client.Request(uri, &productsGetResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsGetResp, nil
}

// 添加我在买的商品
func (p *ProductsClient) Add(productsAddRequestBody *models.ProductsAddRequestBody) (*models.ProductsAddResponse, error) {
	var productsAddResp models.ProductsAddResponse
	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqQuery: core.StructToMapStr(&models.ProductsAddRequestQuery{
			Token: p.client.Token,
		}),
		ReqBody: productsAddRequestBody,
	}

	err := p.client.Request(uri, &productsAddResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsAddResp, nil
}

// 删除我在卖的商品
func (p *ProductsClient) Del(productID string) (*models.ProductsDelResponse, error) {
	var productsDelResp models.ProductsDelResponse

	uri := "/api/market/sellers/products/" + productID

	reqOpts := &core.RequestOption{
		Method: "DELETE",
		ReqQuery: core.StructToMapStr(&models.ProductsDelRequestQuery{
			Token: p.client.Token,
		}),
		ReqBody: &models.ProductsDelRequestBody{},
	}

	err := p.client.Request(uri, &productsDelResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsDelResp, nil
}

// 更新我在卖的商品
func (p *ProductsClient) Update(productsUpdateRequestBody *models.ProductsUpdateRequestBody, productID string) (*models.ProductsUpdateResponse, error) {
	var productsUpdateResp models.ProductsUpdateResponse
	uri := "/api/market/sellers/products/" + productID

	reqOpts := &core.RequestOption{
		Method: "PUT",
		ReqQuery: core.StructToMapStr(&models.ProductsUpdateRequestQuery{
			Token: p.client.Token,
		}),
		ReqBody: productsUpdateRequestBody,
	}

	err := p.client.Request(uri, &productsUpdateResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsUpdateResp, nil
}
