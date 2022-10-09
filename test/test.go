package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/DesistDaydream/dtcg/cards"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func marshal() {
	mapData := make(map[string]string)

	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/价格统计表.xlsx"
	f, err := excelize.OpenFile(file)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			logrus.Errorln(err)
			return
		}
	}()

	cardGroups, _ := cards.GetCardGroups("")
	for _, cardGroup := range cardGroups {
		rows, err := f.GetRows(cardGroup)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"file":  file,
				"sheet": cardGroup,
			}).Fatalf("读取中sheet页异常: %v", err)
		}

		for i := 1; i < len(rows); i++ {
			key := rows[i][2]
			value := rows[i][25]
			mapData[key] = value
		}

		fmt.Println(mapData)
	}

	byteData, _ := json.Marshal(mapData)
	os.WriteFile("cards/card_version_id_and_card_modle.json", byteData, 0666)

}

func unmarshal() {
	mapData := make(map[string]string)
	file, _ := os.ReadFile("cards/card_version_id_and_card_modle.json")
	err := json.Unmarshal(file, &mapData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mapData)
}

func main() {
	marshal()
}
