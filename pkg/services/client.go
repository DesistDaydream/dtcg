package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/models"
	"github.com/sirupsen/logrus"
)

// 获取卡牌系列列表
func GetCardPackage() (*models.CardPackage, error) {
	url := "https://dtcgweb-api.digimoncard.cn/game/gamecard/cachelist"
	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("获取卡牌系列列表失败: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var gameCard *models.CardPackage
	err = json.NewDecoder(resp.Body).Decode(&gameCard)
	if err != nil {
		logrus.Errorf("解析卡牌系列列表失败: %v", err)
		return nil, err
	}

	return gameCard, nil
}

// 根据过滤条件获取卡片详情
func GetCardDesc(cardPackage string) (*models.CardDesc, error) {
	req, err := http.NewRequest("GET", "https://dtcgweb-api.digimoncard.cn/gamecard/gamecardmanager/weblist", nil)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	// 添加参数
	q := req.URL.Query()
	q.Add("page", "")
	q.Add("limit", "300")
	q.Add("name", "")
	q.Add("state", "0")
	q.Add("cardGroup", cardPackage)
	q.Add("rareDegree", "")
	q.Add("belongsType", "")
	q.Add("cardLevel", "")
	q.Add("form", "")
	q.Add("attribute", "")
	q.Add("type", "")
	q.Add("color", "")
	q.Add("envolutionEffect", "")
	q.Add("safeEffect", "")
	q.Add("parallCard", "")
	q.Add("keyEffect", "")
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
	var cardDesc *models.CardDesc
	err = json.Unmarshal(body, &cardDesc)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	return cardDesc, nil
}
