package collector

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/DesistDaydream/dtcg/pkg/fileparse"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// 这三个常量用于给每个 Metrics 名字添加前缀
const (
	name      = "dtcg_exporter"
	Namespace = "dtcg"
)

// Name 用于给前端页面显示 const 常量中定义的内容
func Name() string {
	return name
}

// GetToken 获取 E37 认证所需 Token
func GetToken(opts *JihuansheOpts) (token string, err error) {
	return opts.password, nil
}

// ######## 从此处开始到文件结尾，都是关于配置连接 E37 的代码 ########

// JihuansheClient 连接 E37 所需信息
type JihuansheClient struct {
	Client *http.Client
	Token  string
	Opts   *JihuansheOpts

	JihuansheCardsDesc []models.JihuansheCardDescForPrice
}

// NewE37Client 实例化 E37 客户端
func NewE37Client(opts *JihuansheOpts) *JihuansheClient {
	uri := opts.URL

	u, err := url.Parse(uri)
	if err != nil {
		panic(fmt.Sprintf("invalid E37 URL: %s", err))
	}
	if u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") {
		panic(fmt.Sprintf("invalid E37 URL: %s", uri))
	}

	// ######## 配置 http.Client 的信息 ########
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}
	// 初始化 TLS 相关配置信息
	tlsClientConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    rootCAs,
	}
	// 可以通过命令行选项配置 TLS 的 InsecureSkipVerify
	// 这个配置决定是否跳过 https 协议的验证过程，就是 curl 加不加 -k 选项。默认跳过
	if opts.Insecure {
		tlsClientConfig.InsecureSkipVerify = true
	}
	transport := &http.Transport{
		TLSClientConfig: tlsClientConfig,
	}
	// ######## 配置 http.Client 的信息结束 ########

	// 解析 Excel 文件，获取其中所有卡片信息
	JhsCardsDesc, err := fileparse.NewExcelDataForPrice(opts.File, "", opts.Test)
	if err != nil {
		logrus.Fatalf("从 %v 文件中解析卡牌信息异常：%v", opts.File, err)
	}

	// 第一次启动程序时获取 Token，若无法获取 Token 则程序无法启动
	// token, err := GetToken(opts)
	// if err != nil {
	// 	panic(err)
	// }
	return &JihuansheClient{
		Opts: opts,
		// Token: token,
		Client: &http.Client{
			Timeout:   opts.Timeout,
			Transport: transport,
		},
		JihuansheCardsDesc: JhsCardsDesc.Rows,
	}
}

// Request 建立与 E37 的连接，并返回 Response Body
func (c *JihuansheClient) Request(method string, endpoint string, reqBody io.Reader) (body []byte, err error) {
	// 根据认证信息及 endpoint 参数，创建与 E37 的连接，并返回 Body 给每个 Metric 采集器
	url := c.Opts.URL + endpoint
	logrus.WithFields(logrus.Fields{
		"url":    url,
		"method": method,
	}).Debugf("抓取指标时的请求URL")

	// 创建一个新的 Request
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.Token))

	// 根据新建立的 Request，发起请求，并获取 Response
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error handling request for %s http-statuscode: %s", endpoint, resp.Status)
	}

	// 处理 Response Body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"code": resp.StatusCode,
		"body": string(body),
	}).Tracef("每次请求的响应体以及响应状态码")
	return body, nil
}

// Ping 在 Scraper 接口的实现方法 scrape() 中调用。
// 让 Exporter 每次获取数据时，都检验一下目标设备通信是否正常
func (c *JihuansheClient) Ping() (b bool, err error) {
	return true, nil
}

// 从 Opts 中获取并发数
func (c *JihuansheClient) GetConcurrency() int {
	return c.Opts.Concurrency
}

// 集换社采集器标志
type JihuansheOpts struct {
	URL      string
	Username string
	password string
	// 并发数
	Concurrency int
	// 这俩是关于 http.Client 的选项
	Timeout  time.Duration
	Insecure bool

	// 卡片详情的 Excel 文件
	File string
	// 是否进行测试，若不进行测试，则获取所有卡盒的信息
	Test bool
	// 要采集集换价大于多少的卡的信息
	Price float64
}

// AddFlag use after set Opts
func (o *JihuansheOpts) AddFlag() {
	pflag.StringVar(&o.URL, "e37-server", "https://api.jihuanshe.com", "集换社的 HTTP API 地址。")
	pflag.StringVar(&o.password, "password", "", "token")
	pflag.IntVar(&o.Concurrency, "concurrent", 5, "并发数。")
	pflag.DurationVar(&o.Timeout, "time-out", time.Millisecond*1600, "等待 HTTP 响应的超时时间")
	pflag.BoolVar(&o.Insecure, "insecure", true, "是否禁用 TLS 验证。")
	pflag.StringVar(&o.File, "file", "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/价格统计表.xlsx", "是否进行测试。")
	pflag.BoolVar(&o.Test, "test", false, "是否进行测试。")
	pflag.Float64Var(&o.Price, "price", 1, "要采集集换价大于多少的卡，单位：元")
}
