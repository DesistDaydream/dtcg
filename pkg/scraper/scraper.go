package scraper

import (
	"io"

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
	Scrape(client CommonClient, ch chan<- prometheus.Metric) error
}

// CommonClient 是连接 Server 的客户端接口，不同的 Server，客户端的信息不同。但是至少需要两种行为
// 第一:根据给定的 API 与 Server 建立连接，并获取响应体
// 第二:判断 Server 是否存活
// 这个接口主要是用于为 CommonScraper 接口的 Scrape 方法提供连接 Server 所需的信息。
type CommonClient interface {
	// Request 获取指定 API 下的 响应 Body，并提供给 CommonScraper 接口下的 Scrape() 函数处理这些信息，以便展示 Metrics
	Request(method string, endpoint string, reqBody io.Reader) (body []byte, err error)
	// Ping 每次获取 Metric 时，都会执行健康检查，检查 Server 端是否健康
	Ping() (bool, error)
	// GetConcurrency 获取当前 Server 的并发数
	GetConcurrency() int
}
