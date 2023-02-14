package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community/models"
	"github.com/sirupsen/logrus"
)

const (
	BaseAPI = "https://dtcg-api.moecard.cn"
)

type Client struct {
	Token string
	Retry int
}

func NewClient(token string, retry int) *Client {
	return &Client{
		Token: token,
		Retry: retry,
	}
}

type RequestOption struct {
	Method   string
	ReqQuery map[string]string
	ReqBody  interface{}
}

func (c *Client) Request(uri string, wantResp interface{}, reqOpts *RequestOption) error {
	logrus.WithFields(logrus.Fields{
		"URI":   uri,
		"请求体":   reqOpts.ReqBody,
		"URL参数": reqOpts.ReqQuery,
	}).Debugf("检查请求")

	statusCode, body, err := c.request(uri, reqOpts)
	if err != nil {
		return err
	}

	// DTCGDB 的部分接口在 token 失效时，没有 json 格式的响应体，直接返回 500，所以需要特殊处理
	if statusCode >= 500 {
		logrus.Errorf("DTCGDB 服务器异常，响应码：%v，重新获取 token", statusCode)
		cfg := config.NewConfig("", "")
		c.Token = Login(cfg.DtcgDB.Username, cfg.DtcgDB.Password)
		statusCode, body, err = c.request(uri, reqOpts)
		if err != nil {
			return err
		}
	}

	if statusCode != 200 {
		return fmt.Errorf("响应异常，响应码：%v", statusCode)
	}

	err = json.Unmarshal(body, wantResp)
	if err != nil {
		return fmt.Errorf("解析 %v 异常: %v", string(body), err)
	}

	return nil
}

func (c *Client) request(api string, reqOpts *RequestOption) (int, []byte, error) {
	var (
		originalBody []byte
		req          *http.Request
		resp         *http.Response
		err          error
	)

	url := fmt.Sprintf("%s%s", BaseAPI, api)

	if reqOpts.ReqBody != nil {
		rb, err := json.Marshal(reqOpts.ReqBody)
		if err != nil {
			return 0, nil, fmt.Errorf("解析请求体失败：%v", err)
		}
		req, err = http.NewRequest(reqOpts.Method, url, bytes.NewBuffer(rb))
		if err != nil {
			return 0, nil, fmt.Errorf("构建请求失败：%v", err)
		}
	} else {
		req, err = http.NewRequest(reqOpts.Method, url, nil)
		if err != nil {
			return 0, nil, fmt.Errorf("构建请求失败：%v", err)
		}
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authority", "dtcg-api.moecard.cn")
	if c.Token != "" {
		req.Header.Add("authorization", fmt.Sprintf("Bearer %v", c.Token))
	}

	// 如果有 URL 的 Query 则逐一添加
	if len(reqOpts.ReqQuery) > 0 {
		q := req.URL.Query()
		for key, value := range reqOpts.ReqQuery {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}

	// HTTP 重试，请求过多会被限流返回 429
	// TODO: 限流重试逻辑需要优化
	// time.Sleep(1200 * time.Millisecond)
	for i := 0; i < c.Retry; i++ {
		resp, err = client.Do(req)
		if err != nil {
			logrus.Errorf("获取 HTTP 响应异常：%v。准备重试", err)
		} else if err == nil && resp.StatusCode != 429 && resp.StatusCode < 500 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return 0, nil, fmt.Errorf("读取响应体异常：%v", err)
			}

			return resp.StatusCode, body, nil
		} else if resp.StatusCode >= 500 {
			return resp.StatusCode, nil, nil
		}

		if resp != nil {
			resp.Body.Close()
		}

		logrus.Errorf("第 %v 次 HTTP 请求异常【%v】，等待 10 秒后重试", i+1, resp.Status)

		if req.Body != nil {
			resetBody(req, originalBody)
		}

		time.Sleep(10 * time.Second)
	}

	return 0, nil, req.Context().Err()
}

func resetBody(request *http.Request, originalBody []byte) {
	request.Body = io.NopCloser(bytes.NewBuffer(originalBody))
	request.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBuffer(originalBody)), nil
	}
}

// 开始执行登录前，先检查当前 token 是否可用
func CheckToken(token string) bool {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://dtcg-api.moecard.cn/api/community/user/session", nil)
	if err != nil {
		logrus.Fatalf("TOKEN 检查失败: %v", err)
	}
	req.Header.Add("authority", "dtcg-api.moecard.cn")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %v", token))

	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("TOKEN 检查失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func Login(username, password string) string {
	var loginPostResp models.LoginPostResp

	reqBody, _ := json.Marshal(&models.LoginReqBody{
		Username: username,
		Email:    "",
		Password: password,
	})

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", BaseAPI, "/api/community/user/login"), bytes.NewBuffer(reqBody))
	if err != nil {
		logrus.Errorf("登录失败，创建请求异常: %v", err)
	}
	req.Header.Add("authority", "dtcg-api.moecard.cn")
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("登录失败，HTTP请求异常: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("登录失败，读取响应体异常: %v", err)
	}

	err = json.Unmarshal(body, &loginPostResp)
	if err != nil {
		logrus.Errorf("登录失败，解析响应体异常: %v", err)
	}

	// TODO: 将刚取得的 TOKEN 写入文件中

	return loginPostResp.Data.Token
}
