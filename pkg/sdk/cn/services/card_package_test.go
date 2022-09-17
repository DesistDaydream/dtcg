package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetCardPackage(t *testing.T) {
	cardPackageResp, err := GetCardGroups()
	if err != nil {
		logrus.Fatalln(err)
	}

	jsonByte, _ := json.Marshal(cardPackageResp)

	fileName := filepath.Join("../../../../cards", "card_package.json")
	os.WriteFile(fileName, jsonByte, 0666)
	// for _, p := range cardPackageResp.List {
	// 	logrus.WithFields(logrus.Fields{
	// 		"名称": p.Name,
	// 		"ID": p.ID,
	// 		"状态": p.State,
	// 	}).Info("卡包信息")
	// }
}
