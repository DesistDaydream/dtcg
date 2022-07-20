package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/models"
	"github.com/sirupsen/logrus"
)

type FilterCondition struct {
	Page             string `json:"page"`
	Limit            string `json:"limit"`
	Name             string `json:"name"`
	State            string `json:"state"`
	CardGroup        string `json:"cardGroup"`
	RareDegree       string `json:"rareDegree"`
	BelongsType      string `json:"belongsType"`
	CardLevel        string `json:"cardLevel"`
	Form             string `json:"form"`
	Attribute        string `json:"attribute"`
	Type             string `json:"type"`
	Color            string `json:"color"`
	EnvolutionEffect string `json:"envolutionEffect"`
	SafeEffect       string `json:"safeEffect"`
	ParallCard       string `json:"parallCard"`
	KeyEffect        string `json:"keyEffect"`
}

// 根据过滤条件获取卡片详情
func GetCardsDesc(c *FilterCondition) (*models.CardDesc, error) {
	req, err := http.NewRequest("GET", "https://dtcgweb-api.digimoncard.cn/gamecard/gamecardmanager/weblist", nil)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	// 添加参数
	q := req.URL.Query()
	q.Add("page", c.Page)
	q.Add("limit", c.Limit)
	q.Add("name", c.Name)
	q.Add("state", c.State)
	q.Add("cardGroup", c.CardGroup)
	q.Add("rareDegree", c.RareDegree)
	q.Add("belongsType", c.BelongsType)
	q.Add("cardLevel", c.CardLevel)
	q.Add("form", c.Form)
	q.Add("attribute", c.Attribute)
	q.Add("type", c.Type)
	q.Add("color", c.Color)
	q.Add("envolutionEffect", c.EnvolutionEffect)
	q.Add("safeEffect", c.SafeEffect)
	q.Add("parallCard", c.ParallCard)
	q.Add("keyEffect", c.KeyEffect)
	req.URL.RawQuery = q.Encode()

	logrus.Debugln(req.URL.String())
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
	var cardDesc *models.CardDesc
	err = json.Unmarshal(body, &cardDesc)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	return cardDesc, nil
}
