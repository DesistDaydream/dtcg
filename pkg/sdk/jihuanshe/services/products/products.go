package products

import (
	"encoding/json"
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

func (p *ProductsClient) List(page string) (*models.ProductsListResponse, error) {
	var products models.ProductsListResponse

	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: StructToMapStr(&models.ProductsListRequestQuery{
			GameKey:    "dgm",
			GameSubKey: "sc",
			OnSale:     "1",
			Page:       page,
			Token:      p.client.Token,
		}),
	}

	body, err := p.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &products)

	return &products, nil
}

func (p *ProductsClient) Get(cardVersionID string) (*models.ProductsGetResponse, error) {
	var productsGetresp models.ProductsGetResponse
	uri := "/api/market/products/bySellerCardVersionId"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: StructToMapStr(&models.ProductsGetRequestQuery{
			GameKey:       "dgm",
			SellerUserID:  "609077",
			CardVersionID: cardVersionID,
			Token:         p.client.Token,
		}),
	}

	body, err := p.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &productsGetresp)
	if err != nil {
		return nil, err
	}

	return &productsGetresp, nil
}

func (p *ProductsClient) Add(productsAddRequestBody *models.ProductsAddRequestBody) (*models.ProductsAddResponse, error) {
	var productsAddResponse models.ProductsAddResponse
	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqQuery: StructToMapStr(&models.ProductsAddRequestQuery{
			Token: p.client.Token,
		}),
		ReqBody: StructToMapStr(productsAddRequestBody),
	}

	body, err := p.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &productsAddResponse)
	if err != nil {
		return nil, err
	}

	return &productsAddResponse, nil
}

func (p *ProductsClient) Del(productID string) (*models.ProductsDelResponse, error) {
	var productsDelResponse models.ProductsDelResponse

	uri := "/api/market/sellers/products/" + productID

	reqOpts := &core.RequestOption{
		Method: "DELETE",
		ReqQuery: StructToMapStr(&models.ProductsDelRequestQuery{
			Token: p.client.Token,
		}),
		ReqBody: StructToMapStr(&models.ProductsDelRequestBody{}),
	}

	body, err := p.client.Request(uri, reqOpts)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &productsDelResponse)
	if err != nil {
		return nil, err
	}

	return &productsDelResponse, nil
}

func (p *ProductsClient) Update() {}

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
