package core

import (
	"fmt"
	"io"
	"net/http"
)

const (
	BaseAPIPrefix = "https://dtcgweb-api.digimoncard.cn/card/"
	BaseAPISuffix = "/cachelist"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) RequestCachelist(api string) ([]byte, error) {
	url := fmt.Sprintf("%s%s%s", BaseAPIPrefix, api, BaseAPISuffix)

	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
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
