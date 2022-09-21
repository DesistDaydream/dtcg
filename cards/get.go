package cards

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
)

func GetCardGroups() ([]string, error) {
	var (
		cardGroups     *models.CacheListResp
		cardGroupsName []string
	)

	file, err := os.ReadFile("cards/card_package.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &cardGroups)
	if err != nil {
		return nil, err
	}

	for _, cardGroup := range cardGroups.List {
		cardGroupsName = append(cardGroupsName, cardGroup.Name)
	}

	return cardGroupsName, nil
}

func GetCardLevel() ([]string, error) {
	var cardLevelResp *models.CacheListResp
	var cardLevels []string
	cardLevelFile := filepath.Join("cards", "card_level.json")
	cardLevelFileByte, err := os.ReadFile(cardLevelFile)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(cardLevelFileByte, &cardLevelResp)
	for _, cardLevel := range cardLevelResp.List {
		if cardLevel.Name != "-" && cardLevel.Name != "Lv.2" {
			cardLevels = append(cardLevels, cardLevel.Name)
		}
	}

	return cardLevels, nil
}

func GetCardDesc(cardGroup string) ([]models.CardDesc, error) {
	var cardDescs []models.CardDesc

	cardsDescFile := filepath.Join("cards", cardGroup+".json")

	cardsDescFileByte, err := os.ReadFile(cardsDescFile)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(cardsDescFileByte, &cardDescs)

	return cardDescs, nil
}

// 将所有卡盒中的卡片描述合并为一个整体
func MergeCardsDesc(cardGroups []string) ([]models.CardDesc, error) {
	var allCardsDesc []models.CardDesc
	for _, cardGroup := range cardGroups {
		cardsDesc, err := GetCardDesc(cardGroup)
		if err != nil {
			return nil, err
		}

		allCardsDesc = append(allCardsDesc, cardsDesc...)
	}

	return allCardsDesc, nil
}
