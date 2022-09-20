package collector

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/bitly/go-simplejson"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/xuri/excelize/v2"
)

// 这三个常量用于给每个 Metrics 名字添加前缀
const (
	name      = "dtcg_exporter"
	Namespace = "dtcg"
	//Subsystem(s).
	// exporter = "exporter"
)

// Name 用于给前端页面显示 const 常量中定义的内容
func Name() string {
	return name
}

// GetToken 获取 E37 认证所需 Token
func GetToken(opts *JihuansheOpts) (token string, err error) {
	// 设置 json 格式的 request body
	jsonReqBody := []byte("{\"username\":\"" + opts.Username + "\",\"password\":\"" + opts.password + "\"}")
	// 设置 URL
	url := fmt.Sprintf("%v/api/auth", opts.URL)
	// 设置 Request 信息
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-Type", "application/json")
	// 忽略 TLS 的证书验证
	ts := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 发送 Request 并获取 Response
	resp, err := (&http.Client{Transport: ts}).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 处理 Response Body,并获取 Token
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	jsonRespBody, err := simplejson.NewJson(respBody)
	if err != nil {
		return
	}
	token, err = jsonRespBody.Get("token").String()
	if err != nil {
		return "", fmt.Errorf("GetToken Error：%v", err)
	}
	logrus.WithFields(logrus.Fields{
		"Token": token,
	}).Debugf("Get Token Successed!")
	return
}

// 从 Excel 中获取卡片详情写入到结构体中以便后续使用，其中包括适用于集换社的 card_version_id
func FileToJson(file string) ([]models.JihuansheCardDesc, error) {
	var jihuansheCardsDesc []models.JihuansheCardDesc
	// cardGroups, err := cards.GetCardGroups()
	// if err != nil {
	// 	logrus.Fatalln(err)
	// }

	cardGroups := []string{"STC-01"}

	opts := excelize.Options{}
	f, err := excelize.OpenFile(file, opts)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			logrus.Errorln(err)
			return
		}
	}()

	for _, cardGroup := range cardGroups {
		// 逐行读取Excel文件
		rows, err := f.GetRows(cardGroup)
		if err != nil {
			return nil, fmt.Errorf("读取sheet页%v异常：%v", cardGroup, err)
		}

		// 跳过第一行的标题，从第二行开始，所以 i := 1
		for i := 1; i < len(rows); i++ {
			logrus.WithFields(logrus.Fields{
				"行号":  i,
				"行数据": rows[i],
			}).Debugf("检查每一条需要处理的解析记录")

			// 将每一行中的的每列数据赋值到结构体重
			var erd models.JihuansheCardDesc
			erd.CardGroup = rows[i][1]
			erd.Model = rows[i][2]
			erd.Name = rows[i][9]
			erd.CardVersionID = rows[i][25]

			jihuansheCardsDesc = append(jihuansheCardsDesc, erd)
		}
	}

	return jihuansheCardsDesc, nil
}

// ######## 从此处开始到文件结尾，都是关于配置连接 E37 的代码 ########

// JihuansheClient 连接 E37 所需信息
type JihuansheClient struct {
	Client *http.Client
	Token  string
	Opts   *JihuansheOpts

	JihuansheCardsDesc []models.JihuansheCardDesc
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

	//
	file := "/mnt/e/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/价格统计表.xlsx"
	JhsCardsDesc, err := FileToJson(file)
	if err != nil {
		logrus.Fatalf("从 %v 文件中解析卡牌信息异常：%v", file, err)
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
		JihuansheCardsDesc: JhsCardsDesc,
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

// 验证 Token 时所用的请求体
type token struct {
	Token string `json:"token"`
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

// JihuansheOpts 登录 E37 所需属性
type JihuansheOpts struct {
	URL      string
	Username string
	password string
	// 并发数
	Concurrency int
	// 这俩是关于 http.Client 的选项
	Timeout  time.Duration
	Insecure bool
}

// AddFlag use after set Opts
func (o *JihuansheOpts) AddFlag() {
	pflag.StringVar(&o.URL, "e37-server", "https://api.jihuanshe.com", "集换社的 HTTP API 地址。")
	pflag.StringVar(&o.Username, "e37-user", "admin", "e37 username")
	pflag.StringVar(&o.password, "e37-pass", "admin", "e37 password")
	pflag.IntVar(&o.Concurrency, "concurrent", 5, "并发数。")
	pflag.DurationVar(&o.Timeout, "time-out", time.Millisecond*1600, "等待 HTTP 响应的超时时间")
	pflag.BoolVar(&o.Insecure, "insecure", true, "是否禁用 TLS 验证。")
}
