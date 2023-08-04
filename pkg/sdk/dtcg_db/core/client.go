package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DesistDaydream/dtcg/internal/database"
	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core/models"
	"github.com/sirupsen/logrus"
)

const (
	BaseAPI = "https://dtcg-api.moecard.cn"
)

type Client struct {
	UserID int
	Token  string
	Retry  int
}

func NewClient(userID int, token string, retry int) *Client {
	return &Client{
		UserID: userID,
		Token:  token,
		Retry:  retry,
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
	}).Debugf("Moecard API 检查请求")

	statusCode, body, err := c.request(uri, reqOpts)
	if err != nil {
		return err
	}

	// DTCGDB 的部分接口在 token 失效时，没有 json 格式的响应体，直接返回 500，所以需要特殊处理
	if statusCode >= 500 {
		logrus.Errorf("DTCGDB 服务器异常，响应码：%v，重新获取 token", statusCode)
		// TODO: 这里不应该写死 1
		user, err := database.GetUser("1")
		if err != nil {
			logrus.Fatalf("获取用户信息异常，原因: %v", err)
		}
		c.Token = Login(c.UserID, user.MoecardUsername, user.MoecardPassword)
		_, body, err = c.request(uri, reqOpts)
		if err != nil {
			return err
		}
	}

	var commonResp models.CommonResp
	err = json.Unmarshal(body, &commonResp)
	if err != nil {
		return fmt.Errorf("解析 %v 异常: %v", string(body), err)
	}

	// 这边由于 CommonResp.Data 是 interface{} 类型，想要将数据赋值给传递进来的 wantResp.Data 是不容易的，因为数据类型不一样
	// 曾经试过使用 mapstructure 库将 map[string]interface{} 转为 struct，但是 struct 的层级太深的话(比如 CloudDeckGetRespData.Data.DeckInfo.Main)，数据丢失，并且没找到原因。
	// 后来发现，只需要通过 json 库，将 commonResp 编码为 JSON 二进制流后，再将编码后的结果解码为结构体到 wantResp 即可。
	if commonResp.Success {
		// TODO: 这里要不要直接将 commonResp.Data 编码？就像下面这样，毕竟 Data 才是 interface{} 类型，才需要处理
		// 并且返回的时候，直接返回 Data 即可，错误时，将 Message 返回即可，不用把所有的原始信息都返回。上层调用者本质上也只需要 Data 就够了
		// commonRespDataByte, _ := json.Marshal(commonResp.Data)
		commonRespByte, _ := json.Marshal(commonResp)
		json.Unmarshal(commonRespByte, &wantResp)
		return nil
	} else {
		logrus.Debugf("请求失败: %+v", commonResp)
		return fmt.Errorf("%v", commonResp.Message)
	}
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

func Login(userid int, username, password string) string {
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

	// 将刚取得的 TOKEN 更新到数据库中
	database.UpdateUser(&dbmodels.User{ID: userid}, map[string]interface{}{
		"moecard_token": loginPostResp.Data.Token,
	})

	return loginPostResp.Data.Token
}
