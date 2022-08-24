package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/models"
	"github.com/sirupsen/logrus"
)

func PostDeckSearch(request *models.DeckSearchReq) (*models.DeckSearchResp, error) {
	url := "https://dtcg-api.moecard.cn/api/community/deck/search"
	method := "POST"
	client := &http.Client{}

	bodyBtyes, err := json.Marshal(request.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyBtyes))
	if err != nil {
		return nil, err
	}

	req.Header.Add("authority", "dtcg-api.moecard.cn")
	req.Header.Add("content-type", "application/json")

	q := req.URL.Query()
	q.Add("limit", request.Query.Limit)
	q.Add("page", request.Query.Page)
	req.URL.RawQuery = q.Encode()

	logrus.Debugln(req.URL, req.Body)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("响应异常，响应码：%v，错误信息：%v", resp.Status, string(body))
	}

	var deckSearchResp *models.DeckSearchResp
	err = json.Unmarshal(body, &deckSearchResp)
	if err != nil {
		return nil, err
	}

	return deckSearchResp, nil
}
