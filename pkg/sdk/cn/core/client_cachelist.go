package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BaseAPIPrefix = "https://dtcgweb-api.digimoncard.cn/card/"
	BaseAPISuffix = "/cachelist"
)

func RequestCachelist(api string) ([]byte, error) {
	url := fmt.Sprintf("%s%s%s", BaseAPIPrefix, api, BaseAPISuffix)

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
