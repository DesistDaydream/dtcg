package cardgroup

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/sirupsen/logrus"
)

func AddCardGroup(wirteToJSON bool) {
	cardPackageResp, err := services.GetCardGroups()
	if err != nil {
		logrus.Fatalln(err)
	}

	sort.Slice(cardPackageResp.List, func(i, j int) bool {
		return cardPackageResp.List[i].CreateTime < cardPackageResp.List[j].CreateTime
	})

	if wirteToJSON {
		jsonByte, _ := json.Marshal(cardPackageResp)

		fileName := filepath.Join("cards", "card_package.json")
		os.WriteFile(fileName, jsonByte, 0666)
	}

	for _, cardGroup := range cardPackageResp.List {
		g := &database.CardGroup{
			OfficialID: cardGroup.ID,
			Name:       cardGroup.Name,
			Image:      cardGroup.Image,
			State:      cardGroup.State,
			Position:   cardGroup.Position,
			CreateTime: cardGroup.CreateTime,
			UpdateTime: cardGroup.UpdateTime,
		}
		database.AddCardGroup(g)
	}
}
