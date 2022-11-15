package database

import (
	"testing"

	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
)

func TestGetCardDescByCondition(t *testing.T) {
	initDB()
	got, err := GetCardDescByCondition(5, 1, &models.QueryCardDesc{
		CardPack:   0,
		ClassInput: false,
		Color:      []string{"红", "白"},
		EvoCond:    []models.EvoCond{},
		Keyword:    "奥米加",
		Language:   "",
		OrderType:  "",
		QField: []string{
			"effect",
			"sc_name",
			"evo_cover_effect",
		},
		Rarity:    []string{},
		Tags:      []string{},
		TagsLogic: "",
		Type:      "",
	})
	if err != nil {
		logrus.Errorln(err)
	}

	logrus.Infof("共 %v 条数据，当前第 %v 页，共 %v 页，每页最多 %v 条数据", got.Count, got.PageCurrent, got.PageSize, got.PageTotal)

	for _, g := range got.Data {
		logrus.WithFields(logrus.Fields{
			"名称":  g.ScName,
			"颜色":  g.Color,
			"稀有度": g.Rarity,
			"DP":  g.DP,
		}).Infof("")
	}
}
