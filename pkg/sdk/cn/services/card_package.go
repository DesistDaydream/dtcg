package services

import (
	"encoding/json"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/sirupsen/logrus"
)

// 获取卡组列表
func GetCardGroups() (*models.CardPackage, error) {
	url := "https://dtcgweb-api.digimoncard.cn/game/gamecard/weblist"
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
