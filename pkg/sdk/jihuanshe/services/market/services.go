package market

import (
	"strconv"

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

// 获取用户订单列表(买入)
func (o *MarketClient) OrderList(page string) (*models.BuyerOrdersListResp, error) {
	var buyerOrders models.BuyerOrdersListResp

	uri := "/api/market/orders"
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.OrderListReqQuery{
			Page:   page,
			Status: "complete",
			Token:  o.client.Token,
		}),
	}

	err := o.client.Request(uri, &buyerOrders, reqOpts)
	if err != nil {
		return nil, err
	}

	return &buyerOrders, nil
}

// 获取用户订单详情(买入)
func (o *MarketClient) OrderGet(orderID int) (*models.OrderByBuyerGetResp, error) {
	var orderProducts models.OrderByBuyerGetResp

	orderIDStr := strconv.Itoa(orderID)
	uri := "/api/market/orders/" + orderIDStr
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.OrderGetReqQuery{
			Token: o.client.Token,
		}),
	}

	err := o.client.Request(uri, &orderProducts, reqOpts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}

// 获取用户订单列表（卖出）
func (o *MarketClient) SellerOrderList(page string) (*models.SellerOrderListResp, error) {
	var sellerOrders models.SellerOrderListResp

	uri := "/api/market/sellers/orders"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.OrderListReqQuery{
			Page:   page,
			Status: "complete",
			Token:  o.client.Token,
		}),
	}

	err := o.client.Request(uri, &sellerOrders, reqOpts)
	if err != nil {
		return nil, err
	}

	return &sellerOrders, nil
}

// 获取用户订单详情（卖出）
func (o *MarketClient) SellerOrderGet(orderID int) (*models.OrderBySellerGetResp, error) {
	var orderProducts models.OrderBySellerGetResp

	orderIDStr := strconv.Itoa(orderID)
	uri := "/api/market/sellers/orders/" + orderIDStr
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.OrderGetReqQuery{
			Token: o.client.Token,
		}),
	}

	err := o.client.Request(uri, &orderProducts, reqOpts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}

// 获取商品的“在售”列表
func (p *MarketClient) CardVersionsProductsGet(cardVersionID string, page string) (*models.ProductSellersGetResp, error) {
	var productSellers models.ProductSellersGetResp
	uri := "/api/market/card-versions/products"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.ProductSellersGetReqQuery{
			CardVersionID: cardVersionID,
			Condition:     "1",
			GameKey:       "dgm",
			Page:          page,
		}),
	}

	err := p.client.Request(uri, &productSellers, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productSellers, nil
}
