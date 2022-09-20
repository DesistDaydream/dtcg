package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/models"
)

func GetOrderProducts(orderID int, token string) (*models.OrderProducts, error) {
	orderIDStr := strconv.Itoa(orderID)
	url := "https://api.jihuanshe.com/api/market/orders/" + orderIDStr
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	// req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")

	q := req.URL.Query()
	q.Add("token", token)
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

	var orderProducts models.OrderProducts
	err = json.Unmarshal(body, &orderProducts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}
