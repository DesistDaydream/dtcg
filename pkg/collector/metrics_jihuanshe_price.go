package collector

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var (
	// check interface
	_ CommonScraper = ScrapePrice{}

	// 最低价格
	minPrice = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "jihuanshe", "min_price"),
		"最低价格",
		[]string{"alternative_art", "card_set", "serial", "sc_name", "rarity", "card_version_id"}, nil,
	)
	// 平均价格
	avgPrice = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "jihuanshe", "avg_price"),
		"平均价格",
		[]string{"alternative_art", "card_set", "serial", "sc_name", "rarity", "card_version_id"}, nil,
	)
)

// ScrapePrice 是将要实现 Scraper 接口的一个 Metric 结构体
type ScrapePrice struct {
}

// Name 指定自己定义的 抓取器 的名字，与 Metric 的名字不是一个概念，但是一般保持一致
// 该方法用于为 ScrapeAvgPrice 结构体实现 Scraper 接口
func (s ScrapePrice) Name() string {
	return "jihuanshe_price_info"
}

// Help 指定自己定义的 抓取器 的帮助信息，这里的 Help 的内容将会作为命令行标志的帮助信息。与 Metric 的 Help 不是一个概念。
// 该方法用于为 ScrapeAvgPrice 结构体实现 Scraper 接口
func (s ScrapePrice) Help() string {
	return "Jihuanshe price Info"
}

// Scrape 从客户端采集数据，并将其作为 Metric 通过 channel(通道) 发送。主要就是采集 E37 集群信息的具体行为。
// 该方法用于为 ScrapeAvgPrice 结构体实现 Scraper 接口
func (s ScrapePrice) Scrape(client *JihuansheClient, ch chan<- prometheus.Metric) (err error) {
	cardsPrice, err := database.ListCardsPrice()
	if err != nil {
		logrus.Fatalf("获取卡片价格信息失败: %v", err)
	}

	logrus.Debugf("已获取 %v 条卡片信息", cardsPrice.Count)

	var wg sync.WaitGroup
	defer wg.Wait()

	// 用来控制并发数量
	concurrenceyControl := make(chan bool, client.GetConcurrency())

	c := products.NewProductsClient(core.NewClient(""))

	for _, cardPrice := range cardsPrice.Data {
		concurrenceyControl <- true
		wg.Add(1)
		go func(cardPrice models.CardPrice) {
			defer wg.Done()

			// 只采集集换价在指定价格区间内的卡的信息
			if cardPrice.AvgPrice >= client.Opts.minPrice && cardPrice.AvgPrice <= client.Opts.maxPrice {
				priceInfo, err := c.Get(fmt.Sprint(cardPrice.CardVersionID), "1")
				if err != nil {
					logrus.Errorf("获取卡片 %v-%v %v 信息异常：%v", cardPrice.Serial, cardPrice.ScName, cardPrice.CardVersionID, err)
					<-concurrenceyControl
					return
				}

				fMin, _ := strconv.ParseFloat(priceInfo.MinPrice, 64)
				fAvg, _ := strconv.ParseFloat(priceInfo.AvgPrice, 64)

				// 最低价格
				ch <- prometheus.MustNewConstMetric(minPrice, prometheus.GaugeValue, float64(fMin),
					cardPrice.AlternativeArt,
					cardPrice.SetPrefix,
					cardPrice.Serial,
					cardPrice.ScName,
					cardPrice.Rarity,
					fmt.Sprint(cardPrice.CardVersionID),
				)

				// 平均价格
				ch <- prometheus.MustNewConstMetric(avgPrice, prometheus.GaugeValue, float64(fAvg),
					cardPrice.AlternativeArt,
					cardPrice.SetPrefix,
					cardPrice.Serial,
					cardPrice.ScName,
					cardPrice.Rarity,
					fmt.Sprint(cardPrice.CardVersionID),
				)
			}

			<-concurrenceyControl
		}(cardPrice)
	}

	return nil
}
