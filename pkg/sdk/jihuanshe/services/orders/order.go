package orders

import (
	"strconv"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/orders/models"
)

type OrdersClient struct {
	client *core.Client
}

func NewOrdersClient(client *core.Client) *OrdersClient {
	return &OrdersClient{
		client: client,
	}
}

// 获取用户订单列表(买入)
func (o *OrdersClient) GetBuyerOrders(page string) (*models.BuyerOrdersResponse, error) {
	var buyerOrders models.BuyerOrdersResponse

	uri := "/api/market/orders"
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.BuyerOrdersQuery{
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
func (o *OrdersClient) GetBuyerOrderProducts(orderID int) (*models.BuyerOrderProductsResponse, error) {
	var orderProducts models.BuyerOrderProductsResponse

	orderIDStr := strconv.Itoa(orderID)
	uri := "/api/market/orders/" + orderIDStr
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.BuyerOrderProductsRequest{
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
func (o *OrdersClient) GetSellerOrders(page string) (*models.SellerOrdersResponse, error) {
	var sellerOrders models.SellerOrdersResponse

	uri := "/api/market/sellers/orders"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.BuyerOrdersQuery{
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
func (o *OrdersClient) GetSellerOrderProducts(orderID int) (*models.SellerOrderProductsResponse, error) {
	var orderProducts models.SellerOrderProductsResponse

	orderIDStr := strconv.Itoa(orderID)
	uri := "/api/market/sellers/orders/" + orderIDStr
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.SellerOrderProductsRequest{
			Token: o.client.Token,
		}),
	}

	err := o.client.Request(uri, &orderProducts, reqOpts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}
