package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
)

// 获取卡组列表
func GetCardGroups() (*models.CardGroupsResponse, error) {
	url := "https://dtcgweb-api.digimoncard.cn/game/gamecard/weblist"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("获取卡牌系列列表失败: %v", err)
	}
	defer resp.Body.Close()

	var cardGroups models.CardGroupsResponse
	err = json.NewDecoder(resp.Body).Decode(&cardGroups)
	if err != nil {
		return nil, fmt.Errorf("解析卡牌系列列表失败: %v", err)
	}

	return &cardGroups, nil
}
