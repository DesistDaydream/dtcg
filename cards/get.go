package cards

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
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
