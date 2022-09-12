package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
)

// 根据过滤条件获取卡片详情
func GetCardDescs(r *models.FilterConditionReq) (*models.CardListResponse, error) {
	req, err := http.NewRequest("GET", "https://dtcgweb-api.digimoncard.cn/gamecard/gamecardmanager/weblist", nil)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	// 添加参数
	q := req.URL.Query()
	q.Add("page", r.Page)
	q.Add("limit", r.Limit)
	q.Add("name", r.Name)
	q.Add("state", r.State)
	q.Add("cardGroup", r.CardGroup)
	q.Add("rareDegree", r.RareDegree)
	q.Add("belongsType", r.BelongsType)
	q.Add("cardLevel", r.CardLevel)
	q.Add("form", r.Form)
	q.Add("attribute", r.Attribute)
	q.Add("type", r.Type)
	q.Add("color", r.Color)
	q.Add("envolutionEffect", r.EnvolutionEffect)
	q.Add("safeEffect", r.SafeEffect)
	q.Add("parallCard", r.ParallCard)
	q.Add("keyEffect", r.KeyEffect)
	req.URL.RawQuery = q.Encode()

	// 发起请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	// 解析 JSON 到 struct 中
	var cardDesc *models.CardListResponse
	err = json.Unmarshal(body, &cardDesc)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	return cardDesc, nil
}
