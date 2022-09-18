package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetCardColor(t *testing.T) {
	got, err := GetCardColor()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	fmt.Println(got)
}

func TestGetCardLevel(t *testing.T) {
	cardLevelResp, err := GetCardLevel()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	fmt.Println(cardLevelResp)

	jsonByte, _ := json.Marshal(cardLevelResp)
	fileName := filepath.Join("../../../../cards", "card_level.json")
	os.WriteFile(fileName, jsonByte, 0666)
}

func TestGetCardGetway(t *testing.T) {
	got, err := GetCardGetway()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	fmt.Println(got)
}
