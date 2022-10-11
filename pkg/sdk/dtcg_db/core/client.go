package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	BaseAPI = "https://dtcg-api.moecard.cn"
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
	ReqBody  interface{}
}

func (c *Client) Request(uri string, wantResp interface{}, reqOpts *RequestOption) error {
	body, err := c.request(uri, reqOpts)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"uri":   uri,
		"请求体":   reqOpts.ReqBody,
		"url参数": reqOpts.ReqQuery,
	}).Debugf("检查请求")

	err = json.Unmarshal(body, wantResp)
	if err != nil {
		return fmt.Errorf("解析 %v 异常: %v", string(body), err)
	}

	return nil
}

func (c *Client) request(api string, reqOpts *RequestOption) ([]byte, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	url := fmt.Sprintf("%s%s", BaseAPI, api)

	if reqOpts.ReqBody != nil {
		rb, err := json.Marshal(reqOpts.ReqBody)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(reqOpts.Method, url, bytes.NewBuffer(rb))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(reqOpts.Method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authority", "dtcg-api.moecard.cn")

	// 如果有 URL 的 Query 则逐一添加
	if len(reqOpts.ReqQuery) > 0 {
		q := req.URL.Query()
		for key, value := range reqOpts.ReqQuery {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: 限流重试逻辑
	// 请求过多会被限流返回 429
	time.Sleep(500 * time.Millisecond)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("请求失败，响应码：%v", resp.StatusCode)
	}

	return body, nil
}

func StructToMapStr(obj interface{}) map[string]string {
	data := make(map[string]string)

	v := reflect.ValueOf(obj).Elem()
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
