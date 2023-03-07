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

// 添加我在卖的商品
func (p *ProductsClient) Add(productsAddRequestBody *models.ProductsAddReqBody) (*models.ProductsAddResp, error) {
	var productsAddResp models.ProductsAddResp
	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method:  "POST",
		ReqBody: productsAddRequestBody,
	}

	err := p.client.Request(uri, &productsAddResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsAddResp, nil
}

// 列出我在卖的商品
func (p *ProductsClient) List(page string, keyword string, onSale string) (*models.ProductsListResp, error) {
	var productsResp models.ProductsListResp

	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.ProductsListReqQuery{
			GameKey:    "dgm",
			GameSubKey: "sc",
			Keyword:    keyword,
			OnSale:     onSale,
			Page:       page,
			Sorting:    "published_at_desc",
		}),
	}

	err := p.client.Request(uri, &productsResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsResp, nil
}

// 更新我在卖的商品
func (p *ProductsClient) Update(productsUpdateRequestBody *models.ProductsUpdateReqBody, productID string) (*models.ProductsUpdateResp, error) {
	var productsUpdateResp models.ProductsUpdateResp
	uri := "/api/market/sellers/products/" + productID

	reqOpts := &core.RequestOption{
		Method:  "PUT",
		ReqBody: productsUpdateRequestBody,
	}

	err := p.client.Request(uri, &productsUpdateResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsUpdateResp, nil
}

// 删除我在卖的商品
func (p *ProductsClient) Del(productID string) (*models.ProductsDelResp, error) {
	var productsDelResp models.ProductsDelResp

	uri := "/api/market/sellers/products/" + productID

	reqOpts := &core.RequestOption{
		Method:  "DELETE",
		ReqBody: &models.ProductsDelReqBody{},
	}

	err := p.client.Request(uri, &productsDelResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsDelResp, nil
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
