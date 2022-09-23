package services

import (
	"encoding/json"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
	"github.com/sirupsen/logrus"
)

// 获取卡牌效果关键字
func GetCardGetway() (*models.CacheListResp, error) {
	client := core.NewClient()
	body, err := client.RequestCachelist("cardgetway")
	if err != nil {
		logrus.Errorf("获取卡牌效果关键字列表失败: %v", err)
		return nil, err
	}

	var cardColorResp *models.CacheListResp

	err = json.Unmarshal(body, &cardColorResp)
	if err != nil {
		logrus.Errorf("解析卡牌效果关键字列表失败: %v", err)
		return nil, err
	}

	return cardColorResp, nil
}

// 获取卡牌颜色
func GetCardColor() (*models.CacheListResp, error) {
	client := core.NewClient()
	body, err := client.RequestCachelist("cardcolor")
	if err != nil {
		logrus.Errorf("获取卡牌颜色列表失败: %v", err)
		return nil, err
	}

	var cardColorResp *models.CacheListResp

	err = json.Unmarshal(body, &cardColorResp)
	if err != nil {
		logrus.Errorf("解析卡牌颜色列表失败: %v", err)
		return nil, err
	}

	return cardColorResp, nil
}

// 获取卡牌等级
func GetCardLevel() (*models.CacheListResp, error) {
	client := core.NewClient()
	body, err := client.RequestCachelist("cardlevels")
	if err != nil {
		logrus.Errorf("获取卡牌等级列表失败: %v", err)
		return nil, err
	}

	var cardColorResp *models.CacheListResp

	err = json.Unmarshal(body, &cardColorResp)
	if err != nil {
		logrus.Errorf("解析卡牌等级列表失败: %v", err)
		return nil, err
	}

	return cardColorResp, nil
}
