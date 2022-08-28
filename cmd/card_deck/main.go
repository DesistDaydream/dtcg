package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/models"
)

func main() {
	// 红5,蓝6,黄7,绿8,紫9,黑10,混色11
	cardColors := []string{"5", "6", "7", "8", "9", "10", "11"}
	// 卡组分享3，国内比赛13
	cardType := "3"
	// 无限制：空，简中：chs，日文：ja
	gameEnv := "chs"
	// 需要统计晚于该日志的卡组
	createAt := "2022-08-25 00:00:00"

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

		req := &models.DeckSearchReq{}
		req.Body.Tags = []string{cardType, c}
		req.Body.Kw = ""
		req.Body.Envir = gameEnv
		req.Query.Limit = "50"
		req.Query.Page = "1"

		resp, err := services.PostDeckSearch(req)
		if err != nil {
			logrus.Fatalln(err)
		}

		var count int = 0
		loc, _ := time.LoadLocation("Asia/Shanghai")
		// 需要晚于 BTC2 发售日期
		btc2time, _ := time.ParseInLocation("2006-01-02 15:04:05", createAt, loc)

		for _, deck := range resp.Data.Decks.List {
			tt, _ := time.ParseInLocation("2006-01-02 15:04:05", deck.CreatedAt, loc)
			if tt.After(btc2time) {
				count++
			}
		}

		logrus.WithFields(logrus.Fields{
			"颜色": cardColor,
			"总数": count,
		}).Infoln("卡组统计信息")
	}
}
