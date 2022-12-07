package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community/models"
	"github.com/sirupsen/logrus"
)

type Services struct {
	CoreClient *core.Client
	Cdb        *cdb.CdbClient
	Community  *community.CommunityClient
}

func NewServices(isLogin bool, username, password string, retry int) *Services {
	var token string

	if isLogin {
		token = login(username, password)
	} else {
		token = ""
	}

	s := new(Services)
	s.init(token, retry)
	return s
}

func (s *Services) init(token string, retry int) {
	s.CoreClient = core.NewClient(token, retry)
	s.Cdb = cdb.NewCdbClient(s.CoreClient)
	s.Community = community.NewCommunityClient(s.CoreClient)
}

func login(username, password string) string {
	var loginPostResp models.LoginPostResp

	reqBody, _ := json.Marshal(&models.LoginReqBody{
		Username: username,
		Email:    "",
		Password: password,
	})

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", core.BaseAPI, "/api/community/user/login"), bytes.NewBuffer(reqBody))
	if err != nil {
		logrus.Errorf("登录失败，创建请求异常: %v", err)
	}
	req.Header.Add("authority", "dtcg-api.moecard.cn")
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("登录失败，HTTP请求异常: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("登录失败，读取响应体异常: %v", err)
	}

	err = json.Unmarshal(body, &loginPostResp)
	if err != nil {
		logrus.Fatalf("登录失败，解析响应体异常: %v", err)
	}

	return loginPostResp.Data.Token
}
