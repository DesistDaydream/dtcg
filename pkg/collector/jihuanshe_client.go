package collector

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/DesistDaydream/dtcg/internal/database"
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

// GetToken 获取 集换社 认证所需 Token
func GetToken(opts *JihuansheOpts) (token string, err error) {
	return opts.token, nil
}

// ######## 从此处开始到文件结尾，都是关于配置连接 集换社 的代码 ########

// JihuansheClient 连接 集换社 所需信息
type JihuansheClient struct {
	Client *http.Client
	Token  string
	Opts   *JihuansheOpts

	// CardsPrice *models.CardsPrice
}

// 实例化 HTTP 客户端
func NewJihuansheClient(opts *JihuansheOpts) *JihuansheClient {
	uri := opts.url

	u, err := url.Parse(uri)
	if err != nil {
		panic(fmt.Sprintf("invalid URL: %s", err))
	}
	if u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") {
		panic(fmt.Sprintf("invalid URL: %s", uri))
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
	if opts.insecure {
		tlsClientConfig.InsecureSkipVerify = true
	}
	transport := &http.Transport{
		TLSClientConfig: tlsClientConfig,
	}
	// ######## 配置 http.Client 的信息结束 ########

	dbInfo := &database.DBInfo{
		DBType:   "mysql",
		FilePath: opts.FilePath,
		Server:   opts.MySQLServer,
		Password: opts.MySQLPassword,
	}
	database.InitDB(dbInfo)

	// 第一次启动程序时获取 Token，若无法获取 Token 则程序无法启动
	// token, err := GetToken(opts)
	// if err != nil {
	// 	panic(err)
	// }
	return &JihuansheClient{
		Client: &http.Client{
			Timeout:   opts.timeout,
			Transport: transport,
		},
		Token: "",
		Opts:  opts,
	}
}

type RequestOption struct {
	Method   string
	ReqQuery map[string]string
	ReqBody  interface{}
}

// 发起 HTTP 请求，并返回响应体
func (c *JihuansheClient) Request(endpoint string, reqOpts *RequestOption) (body []byte, err error) {
	var req *http.Request
	// 根据认证信息及 endpoint 参数，创建与 集换社 的连接，并返回 Body 给每个 Metric 采集器
	url := c.Opts.url + endpoint
	logrus.WithFields(logrus.Fields{
		"url":    url,
		"method": reqOpts.Method,
	}).Debugf("抓取指标时的请求URL")

	// 创建一个新的 Request
	if reqOpts.ReqBody != nil {
		rb, err := json.Marshal(reqOpts.ReqBody)
		if err != nil {
			return nil, fmt.Errorf("解析请求体失败：%v", err)
		}
		req, err = http.NewRequest(reqOpts.Method, url, bytes.NewBuffer(rb))
		if err != nil {
			return nil, fmt.Errorf("构建请求失败：%v", err)
		}
	} else {
		req, err = http.NewRequest(reqOpts.Method, url, nil)
		if err != nil {
			return nil, fmt.Errorf("构建请求失败：%v", err)
		}
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
	return c.Opts.concurrency
}

// 集换社采集器标志
type JihuansheOpts struct {
	url   string
	token string
	// 并发数
	concurrency int
	// 这俩是关于 http.Client 的选项
	timeout  time.Duration
	insecure bool

	// 可以用 SQLite 或 MySQL 获取卡片详情
	// SQLite 信息
	FilePath string
	// MySQL 信息
	MySQLServer   string
	MySQLPassword string

	// 是否进行测试，若不进行测试，则获取所有卡盒的信息
	test bool
	// 要采集集换价大于多少的卡的信息
	minPrice float64
	// 要采集集换价小于多少的卡的信息
	maxPrice float64
}

// AddFlag use after set Opts
func (o *JihuansheOpts) AddFlag() {
	pflag.StringVar(&o.url, "jhs-server", "https://api.jihuanshe.com", "集换社的 HTTP API 地址。")
	pflag.StringVar(&o.token, "token", "", "token")
	pflag.IntVar(&o.concurrency, "concurrent", 5, "并发数。")
	pflag.DurationVar(&o.timeout, "time-out", time.Millisecond*1600, "等待 HTTP 响应的超时时间")
	pflag.BoolVar(&o.insecure, "insecure", true, "是否禁用 TLS 验证。")
	pflag.StringVar(&o.FilePath, "filepath", "internal/database/my_dtcg.db", "连接 SQLite 的路径。")
	pflag.StringVar(&o.MySQLServer, "server", "", "数据库连接地址")
	pflag.StringVar(&o.MySQLPassword, "password", "", "数据库连接密码")
	pflag.BoolVar(&o.test, "test", false, "是否进行测试。")
	pflag.Float64Var(&o.minPrice, "min-price", 1, "要采集的卡的集换价的最低价，单位：元")
	pflag.Float64Var(&o.maxPrice, "max-price", 500, "要采集的卡的集换价的最高价，单位：元")
}
