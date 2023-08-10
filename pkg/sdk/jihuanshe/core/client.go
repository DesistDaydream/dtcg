package core

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/utils"
	"github.com/sirupsen/logrus"
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
	ReqBody  interface{}
}

func (c *Client) Request(uri string, wantResp interface{}, reqOpts *RequestOption) error {
	var errorResp models.ErrorResp

	logrus.WithFields(logrus.Fields{
		"URI":   uri,
		"请求体":   reqOpts.ReqBody,
		"URL参数": reqOpts.ReqQuery,
	}).Debugf("检查请求")

	statusCode, body, err := c.request(uri, reqOpts)
	if err != nil {
		return err
	}

	if statusCode != 200 {
		return fmt.Errorf("响应异常，响应码：%v", statusCode)
	}

	// 如果响应体中包含 "error" 这些字符串，则返回错误信息
	if strings.Contains(string(body), "\"error\"") {
		err = json.Unmarshal(body, &errorResp)
		if err != nil {
			return fmt.Errorf("解析 %v 异常: %v", string(body), err)
		}
		return fmt.Errorf("请求失败，错误码：%v，错误信息：%v", errorResp.Code, errorResp.Msg)
	}

	err = json.Unmarshal(body, wantResp)
	if err != nil {
		return fmt.Errorf("解析 %v 异常: %v", string(body), err)
	}

	return nil
}

func (c *Client) request(api string, reqOpts *RequestOption) (int, []byte, error) {
	var (
		// originalBody []byte
		req  *http.Request
		resp *http.Response
		err  error
	)

	url := fmt.Sprintf("%s%s", BaseAPI, api)

	//  如果有请求体则添加
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
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.jihuanshe.com")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

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
		return 0, nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, body, nil
}

func (c *Client) RequestWithEncrypt(uri string, wantResp interface{}, reqOpts *RequestOption) error {
	// var errorResp models.ErrorResp

	logrus.WithFields(logrus.Fields{
		"URI":   uri,
		"请求体":   reqOpts.ReqBody,
		"URL参数": reqOpts.ReqQuery,
	}).Debugf("检查请求")

	// 生成 16 个字符的随机数作为对称加密所需的密码
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	key := hex.EncodeToString(bytes)

	encryptedKey, err := utils.EncryptWithRsaPublicKey(key, utils.JhsRsaPublicKey)
	if err != nil {
		return fmt.Errorf("解析公钥失败: %v", err)
	}
	logrus.Debugf("key: %v", base64.StdEncoding.EncodeToString(encryptedKey))

	a := utils.NewAesCrypto([]byte(key))
	reqBodyByte, _ := json.Marshal(reqOpts.ReqBody)
	encryptedData, err := a.AesEncryptECB(reqBodyByte)
	if err != nil {
		return fmt.Errorf("加密请求体失败: %v", err)
	}
	logrus.Debugf("data: %v", base64.StdEncoding.EncodeToString(encryptedData))

	req, err := http.NewRequest(reqOpts.Method, fmt.Sprintf("%s%s", BaseAPI, uri), nil)
	if err != nil {
		return fmt.Errorf("构建请求失败：%v", err)
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.jihuanshe.com")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Add("key", base64.StdEncoding.EncodeToString(encryptedKey))

	q := req.URL.Query()
	q.Add("data", base64.StdEncoding.EncodeToString(encryptedData))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 || string(body) == "" {
		return fmt.Errorf("响应体为空或响应异常，响应码：%v", resp.StatusCode)
	}

	logrus.Debugf("检查响应体: %s", string(body))

	type respData struct {
		Data string `json:"data"`
	}

	var r respData

	err = json.Unmarshal(body, &r)
	if err != nil {
		return fmt.Errorf("解析加密的响应体异常:%v", err)
	}

	dataByte, err := base64.StdEncoding.DecodeString(r.Data)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	decryptedData, err := a.AesDecryptECB(dataByte)
	if err != nil {
		return fmt.Errorf("解密响应体失败：%v", resp.StatusCode)
	}

	// 用于在不知道响应体的情况下，看看响应体结构以便写 struct
	// fmt.Println(string(decryptedData))

	err = json.Unmarshal(decryptedData, wantResp)
	if err != nil {
		return fmt.Errorf("解析已解密的响应体异常: %v", err)
	}

	return nil
}
