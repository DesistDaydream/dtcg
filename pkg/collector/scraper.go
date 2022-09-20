package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

//https://github.com/prometheus/mysqld_exporter/blob/master/collector/scraper.go

// CommonScraper 接口是抓取器的最小接口，可让我们向本 exporter 添加新的 Prometheus 指标。
// 可以这么理解，每个抓取 Metric 的行为，都会抽象成一个 **Scraper(抓取器)**。
// 并且，可以通过命令行标志来控制开启或关闭哪个抓取器
// 注意：抓取器的 Name和Help 与 Metric 的 Name和Help 不是一个概念
type CommonScraper interface {
	// Name 是抓取器的名称. Should be unique.
	Name() string

	// Help 是抓取器的帮助信息，这里的 Help 的内容将会作为命令行标志的帮助信息。
	Help() string

	// Scrape 是抓取器的具体行为。从客户端采集数据，并将其作为 Metric 通过 channel(通道) 发送。
	Scrape(client *JihuansheClient, ch chan<- prometheus.Metric) error
}
