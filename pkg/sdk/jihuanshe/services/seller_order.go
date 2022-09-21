package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/models"
)

func GetSellerOrders(page string, token string) (*models.SellerOrders, error) {
	url := "https://api.jihuanshe.com/api/market/sellers/orders"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	// req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")

	q := req.URL.Query()
	q.Add("page", page)
	q.Add("token", token)
	q.Add("status", "complete,waiting_to_confirm,waiting_to_pay,waiting_to_send,waiting_to_receive,waiting_to_refund,waiting_to_return_goods")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var sellerOrders models.SellerOrders
	err = json.Unmarshal(body, &sellerOrders)
	if err != nil {
		return nil, fmt.Errorf("解析异常：%v", err)
	}

	return &sellerOrders, nil
}
