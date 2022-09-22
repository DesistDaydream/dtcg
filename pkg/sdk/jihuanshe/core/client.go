package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseAPI = "https://api.jihuanshe.com"
)

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

type RequestOption struct {
	Method   string
	ReqQuery map[string]string
	ReqBody  map[string]string
}

func (c *Client) Request(api string, reqOpts *RequestOption) ([]byte, error) {
	var (
		rb  *bytes.Buffer
		req *http.Request
		err error
	)

	url := fmt.Sprintf("%s%s", BaseAPI, api)

	//  如果有请求体则添加
	if len(reqOpts.ReqBody) > 0 {
		requestBody, err := json.Marshal(reqOpts.ReqBody)
		if err != nil {
			return nil, err
		}
		rb = bytes.NewBuffer(requestBody)
	}

	if rb != nil {
		req, err = http.NewRequest(reqOpts.Method, url, rb)
	} else {
		req, err = http.NewRequest(reqOpts.Method, url, nil)
	}
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	// 如果有 URL 的 Query 则逐一添加
	if len(reqOpts.ReqQuery) > 0 {
		q := req.URL.Query()
		for key, value := range reqOpts.ReqQuery {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
