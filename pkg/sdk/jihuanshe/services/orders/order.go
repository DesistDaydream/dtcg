package orders

import (
	"encoding/json"
	"fmt"
	"reflect"
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

func (o *OrdersClient) GetBuyerOrders(page string) (*models.BuyerOrdersResponse, error) {
	var buyerOrders models.BuyerOrdersResponse

	uri := "/api/market/orders"
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: StructToMapStr(&models.BuyerOrdersQuery{
			Page:   page,
			Status: "complete,waiting_to_confirm,waiting_to_pay,waiting_to_send,waiting_to_receive,waiting_to_refund,waiting_to_return_goods",
			Token:  o.client.Token,
		}),
	}

	body, err := o.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &buyerOrders)
	if err != nil {
		return nil, fmt.Errorf("解析异常：%v", err)
	}

	return &buyerOrders, nil
}

func (o *OrdersClient) GetBuyerOrderProducts(orderID int) (*models.BuyerOrderProductsResponse, error) {
	var orderProducts models.BuyerOrderProductsResponse

	orderIDStr := strconv.Itoa(orderID)
	uri := "/api/market/orders/" + orderIDStr
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: StructToMapStr(&models.BuyerOrderProductsRequest{
			Token: o.client.Token,
		}),
	}

	body, err := o.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &orderProducts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}

func (o *OrdersClient) GetSellerOrders(page string) (*models.SellerOrdersResponse, error) {
	var sellerOrders models.SellerOrdersResponse

	uri := "/api/market/sellers/orders"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: StructToMapStr(&models.BuyerOrdersQuery{
			Page:   page,
			Status: "complete,waiting_to_confirm,waiting_to_pay,waiting_to_send,waiting_to_receive,waiting_to_refund,waiting_to_return_goods",
			Token:  o.client.Token,
		}),
	}

	body, err := o.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &sellerOrders)
	if err != nil {
		return nil, fmt.Errorf("解析异常：%v", err)
	}

	return &sellerOrders, nil
}

func (o *OrdersClient) GetSellerOrderProducts(orderID int) (*models.SellerOrderProductsResponse, error) {
	var orderProducts models.SellerOrderProductsResponse

	orderIDStr := strconv.Itoa(orderID)
	uri := "/api/market/sellers/orders/" + orderIDStr
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: StructToMapStr(&models.SellerOrderProductsRequest{
			Token: o.client.Token,
		}),
	}

	body, err := o.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &orderProducts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}

func StructToMapStr(obj interface{}) map[string]string {
	data := make(map[string]string)

	objV := reflect.ValueOf(obj)
	v := objV.Elem()
	typeOfType := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tField := typeOfType.Field(i)
		tFieldTag := string(tField.Tag.Get("query"))
		if len(tFieldTag) > 0 {
			data[tFieldTag] = field.String()
		} else {
			data[tField.Name] = field.String()
		}
	}

	return data
}
