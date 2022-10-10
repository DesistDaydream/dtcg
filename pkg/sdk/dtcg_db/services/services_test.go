package services

import (
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/sirupsen/logrus"
)

func TestSearchClient_CardDeckSearch(t *testing.T) {
	client := NewSearchClient(core.NewClient(""))
	got, err := client.PostCardSearch(50)
	if err != nil {
		logrus.Fatalln(err)
	}

	fmt.Println(got)
}

// 获取卡包信息
func TestSearchClient_GetSeries(t *testing.T) {
	client := NewSearchClient(core.NewClient(""))
	series, err := client.GetSeries()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, serie := range series.Data {
		for _, pack := range serie.SeriesPack {
			if pack.Language == "chs" {
				logrus.WithFields(logrus.Fields{
					"前缀": pack.PackPrefix,
					"名称": pack.PackName,
					"ID": pack.PackID,
				}).Infof("%v 中的卡包信息", serie.SeriesName)
			}
		}
	}
}
