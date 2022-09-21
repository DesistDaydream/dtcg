package collector

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
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
		[]string{"card_group", "model", "name", "parall_card", "card_version_id"}, nil,
	)
	// 平均价格
	avgPrice = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "jihuanshe", "avg_price"),
		"平均价格",
		[]string{"card_group", "model", "name", "parall_card", "card_version_id"}, nil,
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
	var wg sync.WaitGroup
	defer wg.Wait()

	// 用来控制并发数量
	concurrenceyControl := make(chan bool, client.GetConcurrency())

	for _, jhsCardDesc := range client.JihuansheCardsDesc {
		concurrenceyControl <- true
		wg.Add(1)
		url := fmt.Sprintf("/api/market/products/bySellerCardVersionId?seller_user_id=550420&card_version_id=%v&game_key=dgm", jhsCardDesc.CardVersionID)
		method := "GET"
		go func(url string, jhsCardDesc models.JihuansheCardDesc) {
			defer wg.Done()
			var cardInfo CardInfo
			respBodyUser, err := client.Request(method, url, nil)
			if err != nil {
				logrus.Errorf("获取 %v 用户数据失败，原因:%v", url, err)
				<-concurrenceyControl
				return
			}
			err = json.Unmarshal(respBodyUser, &cardInfo)
			if err != nil {
				logrus.Errorf("解析 %v 用户数据失败，原因:%v", url, err)
				<-concurrenceyControl
				return
			}
			fMin, _ := strconv.ParseFloat(cardInfo.MinPrice, 64)
			fAvg, _ := strconv.ParseFloat(cardInfo.AvgPrice, 64)
			// 最低价格
			ch <- prometheus.MustNewConstMetric(minPrice, prometheus.GaugeValue, float64(fMin),
				jhsCardDesc.CardGroup,
				jhsCardDesc.Model,
				jhsCardDesc.Name,
				jhsCardDesc.ParallCard,
				jhsCardDesc.CardVersionID,
			)

			// 平均价格
			ch <- prometheus.MustNewConstMetric(avgPrice, prometheus.GaugeValue, float64(fAvg),
				jhsCardDesc.CardGroup,
				jhsCardDesc.Model,
				jhsCardDesc.Name,
				jhsCardDesc.ParallCard,
				jhsCardDesc.CardVersionID,
			)

			<-concurrenceyControl
		}(url, jhsCardDesc)
	}

	return nil
}

// 用户统计信息
type CardInfo struct {
	CardID               int         `json:"card_id"`
	CardNameCn           string      `json:"card_name_cn"`
	CardSubName          string      `json:"card_sub_name"`
	CardVersionNumber    string      `json:"card_version_number"`
	CardVersionRarity    string      `json:"card_version_rarity"`
	CardVersionImage     string      `json:"card_version_image"`
	UserCardVersionImage interface{} `json:"user_card_version_image"`
	Language             interface{} `json:"language"`
	ProductLanguage      string      `json:"product_language"`
	MinPrice             string      `json:"min_price"`
	AvgPrice             string      `json:"avg_price"`
	Warehouse            bool        `json:"warehouse"`
	Products             []Products  `json:"products"`
}
type Products struct {
	ProductID       int         `json:"product_id"`
	CardNameCn      string      `json:"card_name_cn"`
	Price           float32     `json:"price"`
	Quantity        int         `json:"quantity"`
	Condition       int         `json:"condition"`
	Remark          string      `json:"remark"`
	PublishLocation interface{} `json:"publish_location"`
}
