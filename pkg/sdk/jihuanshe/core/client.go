package core

import (
	"fmt"
	"io/ioutil"
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

func Request(api string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", BaseAPI, api)
	// method := "GET"

	resp, err := http.Get(url)
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
