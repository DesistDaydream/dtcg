package sellers

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/sellers/models"
)

type SellersClient struct {
	client *core.Client
}

func NewSellersClient(client *core.Client) *SellersClient {
	return &SellersClient{
		client: client,
	}
}

// 添加我在卖的商品
func (p *SellersClient) ProductAdd(productsAddRequestBody *models.ProductsAddReqBody) (*models.ProductsAddResp, error) {
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
func (p *SellersClient) ProductList(page, keyword, onSale, sorting string) (*models.ProductsListResp, error) {
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
			Rarity:     "",
		}),
	}

	err := p.client.Request(uri, &productsResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsResp, nil
}

// 更新我在卖的商品
func (p *SellersClient) ProductUpdate(productsUpdateRequestBody *models.ProductsUpdateReqBody, productID string) (*models.ProductsUpdateResp, error) {
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
func (p *SellersClient) ProductDel(productID string) (*models.ProductsDelResp, error) {
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

// 提现信息
func (s *SellersClient) Withdraw(page string) (*models.WithdrawResp, error) {
	var withdrawResp models.WithdrawResp
	uri := "/api/market/sellers/withdraw/withdrawLogs"

	reqOpts := &core.RequestOption{
		Method:   "GET",
		ReqQuery: map[string]string{"page": page},
	}

	err := s.client.Request(uri, &withdrawResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &withdrawResp, nil
}
