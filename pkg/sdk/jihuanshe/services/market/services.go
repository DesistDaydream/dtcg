package market

import (
	"strconv"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/utils"
)

type MarketClient struct {
	client *core.Client
}

func NewMarketClient(client *core.Client) *MarketClient {
	return &MarketClient{
		client: client,
	}
}

// 更新 Token
func (m *MarketClient) AuthUpdateTokenPost() (*models.CommonSuccessResp, error) {
	var commonSuccessResp models.CommonSuccessResp
	uri := "/api/market/auth/update-push-token"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqBody: &models.UpdateTokenPostReqBody{
			PushDevice: "ios",
			Token:      m.client.Token,
		},
	}

	err := m.client.Request(uri, &commonSuccessResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &commonSuccessResp, nil
}

// 添加我在卖的商品
func (m *MarketClient) SellersProductsAdd(productsAddRequestBody *models.ProductsAddReqBody) (*models.CommonSuccessResp, error) {
	var productsAddResp models.CommonSuccessResp
	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method:  "POST",
		ReqBody: productsAddRequestBody,
	}

	err := m.client.Request(uri, &productsAddResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsAddResp, nil
}

// 列出我在卖的商品
func (m *MarketClient) SellersProductsList(page, keyword, onSale, sorting string) (*models.ProductsListResp, error) {
	var productsResp models.ProductsListResp

	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.ProductsListReqQuery{
			GameKey:    "dgm",
			GameSubKey: "sc",
			Keyword:    keyword,
			OnSale:     onSale,
			Page:       page,
			Sorting:    "published_at_desc",
			Rarity:     "",
		}),
	}

	err := m.client.Request(uri, &productsResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsResp, nil
}

// 更新我在卖的商品
func (m *MarketClient) SellersProductsUpdate(productsUpdateRequestBody *models.ProductsUpdateReqBody, productID string) (*models.ProductsUpdateResp, error) {
	var productsUpdateResp models.ProductsUpdateResp
	uri := "/api/market/sellers/products/" + productID

	reqOpts := &core.RequestOption{
		Method:  "PUT",
		ReqBody: productsUpdateRequestBody,
	}

	err := m.client.Request(uri, &productsUpdateResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsUpdateResp, nil
}

// 删除我在卖的商品
func (m *MarketClient) SellersProductsDel(productID string) (*models.CommonSuccessResp, error) {
	var productsDelResp models.CommonSuccessResp

	uri := "/api/market/sellers/products/" + productID

	reqOpts := &core.RequestOption{
		Method:  "DELETE",
		ReqBody: &models.ProductsDelReqBody{},
	}

	err := m.client.Request(uri, &productsDelResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsDelResp, nil
}

// 获取提现日志
func (m *MarketClient) SellersWithdrawLogGet(page string) (*models.WithdrawResp, error) {
	var withdrawResp models.WithdrawResp
	uri := "/api/market/sellers/withdraw/withdrawLogs"

	reqOpts := &core.RequestOption{
		Method:   "GET",
		ReqQuery: map[string]string{"page": page},
	}

	err := m.client.Request(uri, &withdrawResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &withdrawResp, nil
}

// 获取用户订单列表(买入)
func (m *MarketClient) OrderList(page string) (*models.BuyerOrdersListResp, error) {
	var buyerOrders models.BuyerOrdersListResp

	uri := "/api/market/orders"
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.OrderListReqQuery{
			Page:   page,
			Status: "complete",
			Token:  m.client.Token,
		}),
	}

	err := m.client.Request(uri, &buyerOrders, reqOpts)
	if err != nil {
		return nil, err
	}

	return &buyerOrders, nil
}

// 获取用户订单详情(买入)
func (m *MarketClient) OrderGet(orderID int) (*models.OrderByBuyerGetResp, error) {
	var orderProducts models.OrderByBuyerGetResp

	uri := "/api/market/orders/" + strconv.Itoa(orderID)
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.OrderGetReqQuery{
			Token: m.client.Token,
		}),
	}

	err := m.client.Request(uri, &orderProducts, reqOpts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}

// 获取用户订单列表（卖出）
func (m *MarketClient) SellerOrderList(page string) (*models.SellerOrderListResp, error) {
	var sellerOrders models.SellerOrderListResp

	uri := "/api/market/sellers/orders"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.OrderListReqQuery{
			Page:   page,
			Status: "complete",
			Token:  m.client.Token,
		}),
	}

	err := m.client.Request(uri, &sellerOrders, reqOpts)
	if err != nil {
		return nil, err
	}

	return &sellerOrders, nil
}

// 获取用户订单详情（卖出）
func (m *MarketClient) SellerOrderGet(orderID int) (*models.OrderBySellerGetResp, error) {
	var orderProducts models.OrderBySellerGetResp

	uri := "/api/market/sellers/orders/" + strconv.Itoa(orderID)
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.OrderGetReqQuery{
			Token: m.client.Token,
		}),
	}

	err := m.client.Request(uri, &orderProducts, reqOpts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}

// 获取商品的“在售”列表
func (m *MarketClient) CardVersionsProductsGet(cardVersionID string, page string) (*models.ProductSellersGetResp, error) {
	var productSellers models.ProductSellersGetResp
	uri := "/api/market/card-versions/products"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.ProductSellersGetReqQuery{
			CardVersionID: cardVersionID,
			Condition:     "1",
			GameKey:       "dgm",
			Page:          page,
		}),
	}

	err := m.client.Request(uri, &productSellers, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productSellers, nil
}
