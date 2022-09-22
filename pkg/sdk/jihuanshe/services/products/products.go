package products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

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

const (
	BaseAPI = "https://api.jihuanshe.com"
)

type RequestOption struct {
	Method   string
	ReqQuery map[string]string
}

func request(api string, reqOpts *RequestOption) ([]byte, error) {
	url := fmt.Sprintf("%s%s", BaseAPI, api)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range reqOpts.ReqQuery {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p *ProductsClient) List(page string) (*models.ProductsListResponse, error) {
	var products models.ProductsListResponse

	uri := "/api/market/sellers/products"

	reqOpts := &RequestOption{
		Method: "GET",
		ReqQuery: StructToMapStr(&models.ProductsListRequestQuery{
			GameKey:    "dgm",
			GameSubKey: "sc",
			OnSale:     "1",
			Page:       page,
			Token:      p.client.Token,
		}),
	}

	body, err := request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &products)

	return &products, nil
}

func (p *ProductsClient) Get(cardVersionID string) (*models.ProductsGetResponse, error) {
	var productsGetresp models.ProductsGetResponse
	uri := "/api/market/products/bySellerCardVersionId"

	reqOpts := &RequestOption{
		Method: "GET",
		ReqQuery: StructToMapStr(&models.ProductsGetRequestQuery{
			GameKey:       "dgm",
			SellerUserID:  "609077",
			CardVersionID: cardVersionID,
			Token:         p.client.Token,
		}),
	}

	body, err := request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &productsGetresp)
	if err != nil {
		return nil, err
	}

	return &productsGetresp, nil
}

func (p *ProductsClient) Add() {}

func (p *ProductsClient) Del() {}

func (p *ProductsClient) Update() {}

func StructToMapStr(obj interface{}) map[string]string {
	data := make(map[string]string)

	objV := reflect.ValueOf(obj)
	v := objV.Elem()
	typeOfType := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tField := typeOfType.Field(i)
		tFieldTag := string(tField.Tag)
		if len(tFieldTag) > 0 {
			data[tFieldTag] = field.String()
		} else {
			data[tField.Name] = field.String()
		}
	}
	return data
}
