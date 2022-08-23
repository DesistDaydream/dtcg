package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/models"
)

func PostDeck(reqBodyByte []byte, createAt string) (int, error) {
	url := "https://dtcg-api.moecard.cn/api/community/deck/search"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBodyByte))
	if err != nil {
		logrus.Fatalln(err)
		return 0, err
	}
	req.Header.Add("authority", "dtcg-api.moecard.cn")
	req.Header.Add("content-type", "application/json")

	q := req.URL.Query()
	q.Add("limit", "500")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		logrus.Fatalln(err)
		return 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Fatalln(err)
		return 0, err
	}

	var deck *models.Deck
	err = json.Unmarshal(body, &deck)
	if err != nil {
		logrus.Errorf("解析 %v 数据失败：%v", string(reqBodyByte), err)
	}

	var count int = 0
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// 需要晚于 BTC2 发售日期
	btc2time, _ := time.ParseInLocation("2006-01-02 15:04:05", createAt, loc)

	for _, deck := range deck.Data.Decks.List {
		tt, _ := time.ParseInLocation("2006-01-02 15:04:05", deck.CreatedAt, loc)
		if tt.After(btc2time) {
			count++
		}
	}

	return count, nil
}

func main() {
	// 红5,蓝6,黄7,绿8,紫9,黑10,混色11
	cardColors := []string{"5", "6", "7", "8", "9", "10", "11"}
	// 卡组分享3，国内比赛13
	cardType := "3"
	// 无限制：空，简中：chs，日文：ja
	gameEnv := "chs"
	// 需要统计晚于该日志的卡组
	createAt := "2022-08-20 00:00:00"

	var cardColor string

	for _, c := range cardColors {
		switch c {
		case "5":
			cardColor = "红"
		case "6":
			cardColor = "蓝"
		case "7":
			cardColor = "黄"
		case "8":
			cardColor = "绿"
		case "9":
			cardColor = "紫"
		case "10":
			cardColor = "黑"
		case "11":
			cardColor = "混色"
		}

		r := &models.DeckReq{
			Tags: []string{
				cardType,
				c,
			},
			Kw:    "",
			Envir: gameEnv,
		}
		reqBody, err := json.Marshal(r)
		if err != nil {
			logrus.Errorln(err)
		}

		count, err := PostDeck(reqBody, createAt)
		if err != nil {
			logrus.Errorln(err)
		}

		logrus.WithFields(logrus.Fields{
			"颜色": cardColor,
			"总数": count,
		}).Infoln("卡组统计信息")
	}
}
