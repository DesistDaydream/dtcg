package products

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

func (p *ProductsClient) List(page string) (*models.ProductsResponse, error) {
	var products models.ProductsResponse

	url := "https://api.jihuanshe.com/api/market/sellers/products"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("game_key", "dgm")
	q.Add("game_sub_key", "sc")
	q.Add("page", page)
	q.Add("token", p.client.Token)
	q.Add("on_sale", "1")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &products)

	return &products, nil
}

func (p *ProductsClient) Get() {}

func (p *ProductsClient) Add() {}

func (p *ProductsClient) Del() {}

func (p *ProductsClient) Update() {}
