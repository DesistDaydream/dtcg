package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	carddesc "github.com/DesistDaydream/dtcg/cmd/my_dtcg/card_desc"
	cardgroup "github.com/DesistDaydream/dtcg/cmd/my_dtcg/card_group"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	Add string
}

func AddFlsgs(f *Flags) {
	pflag.StringVarP(&f.Add, "add", "a", "", "向数据库添加数据的内容")

}

func main() {
	var flags Flags
	AddFlsgs(&flags)
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()
	pflag.Parse()

	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	i := &database.DBInfo{
		FilePath: "internal/database/my_dtcg.db",
	}
	database.InitDB(i)

	switch flags.Add {
	case "carddesc":
		carddesc.AddCardDesc()
	case "cardgroup":
		cardgroup.AddCardGroup(false)
	case "dtcgdb":
		AddCardDescFromDtcgDB()
	default:
		logrus.Errorln("使用 --add 指定要添加的数据")
	}
}

func AddCardDescFromDtcgDB() {
	url := "https://dtcg-api.moecard.cn/api/cdb/cards/search?page=1&limit=700"
	method := "POST"

	payload := strings.NewReader(`{"keyword":"","language":"chs","class_input":false,"card_pack":"","type":"","color":[],"rarity":[],"tags":[],"tags__logic":"or","order_type":"default","evo_cond":[{}],"qField":[]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "dtcg-api.moecard.cn")
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var a database.CardsDescFromDtcgDB

	err = json.Unmarshal(body, &a)
	if err != nil {
		logrus.Errorln(err)
	}

	for _, l := range a.Data.List {
		database.AddCardDescFromDtcgDB(&l)
	}
}
