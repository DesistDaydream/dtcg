package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community/models"
)

func main() {
	// 红5,蓝6,黄7,绿8,紫9,黑10,混色11
	cardColors := []string{"5", "6", "7", "8", "9", "10", "11"}
	// 卡组分享3，国内比赛13
	cardType := "13"
	// 无限制：空，简中：chs，日文：ja
	gameEnv := "chs"
	// 需要统计晚于该日志的卡组
	createAt := "2023-07-01 00:00:00"

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

		reqBody := &models.DeckSearchReqBody{
			Tags:  []string{cardType, c},
			Kw:    "",
			Envir: gameEnv,
		}
		reqQuery := &models.DeckSearchReqQuery{
			Limit: "100",
			Page:  "1",
		}

		client := community.NewCommunityClient(core.NewClient("", 1))

		resp, err := client.PostDeckSearch(reqBody, reqQuery)
		if err != nil {
			logrus.Fatalf("获取卡组列表异常：%v", err)
		}

		var count int = 0
		loc, _ := time.LoadLocation("Asia/Shanghai")
		// 需要晚于指定卡包的发售日期
		btctime, _ := time.ParseInLocation("2006-01-02 15:04:05", createAt, loc)
		logrus.Debugf("指定的要统计的时间：%v", btctime)

		for _, deck := range resp.Data.Decks.List {
			tt, _ := time.ParseInLocation("2006-01-02 15:04:05", deck.CreatedAt, loc)
			if tt.After(btctime) {
				count++
			}
		}

		logrus.WithFields(logrus.Fields{
			"颜色": cardColor,
			"总数": count,
		}).Infoln("卡组统计信息")
	}
}
