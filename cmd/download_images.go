package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/DesistDaydream/dtcg/pkg/models"
	"github.com/DesistDaydream/dtcg/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type ImageHandler struct {
	Lang        string
	CardPackage string
	URL         string

	// 文件保存路径相关信息
	Dir      string
	FileName string
	FilePath string
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{
		Lang:        "",
		CardPackage: "",
		FileName:    "",
		URL:         "",
	}
}

// 获取需要保存图片的目录
func (i *ImageHandler) GetDir() *ImageHandler {
	i.Dir = "./images/" + i.Lang + "/" + i.CardPackage
	return i
}

// 获取图片保存路径
// 1.获取图片语言
func (i *ImageHandler) GetLang(lang string) *ImageHandler {
	i.Lang = lang
	return i
}

// 2.获取卡包名称
func (i *ImageHandler) GetCardPackage(cardPackageName string) *ImageHandler {
	i.CardPackage = cardPackageName
	return i
}

// 3.获取图片名称
func (i *ImageHandler) GetFileName(url string) *ImageHandler {
	// 提取 url 中的文件名
	i.FileName = url[strings.LastIndex(url, "/")+1:]
	return i
}

// 4.获取图片保存路径
func (i *ImageHandler) GetPath() *ImageHandler {
	i.FilePath = "./images/" + i.Lang + "/" + i.CardPackage + "/" + i.FileName
	return i
}

// 下载图片
func (i *ImageHandler) downloadImage() {
	// 判断目录中是否有这张图片
	if _, err := os.Stat(i.FilePath); err == nil {
		logrus.Infof("%v 图片已存在", i.FileName)
		return
	}

	// 下载图片
	resp, err := http.Get(i.URL)
	if err != nil {
		logrus.Fatalln(err)
		return
	}

	// 创建文件
	file, err := os.Create(i.FilePath)
	if err != nil {
		logrus.Fatalln(err)
		return
	}

	// 写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		logrus.Fatalln(err)
		return
	}

	// 关闭文件
	file.Close()

	// 关闭响应
	resp.Body.Close()

	logrus.Infof("下载完成")
}

func (i *ImageHandler) GetImagesURL(lang string, cardPackage string) ([]string, error) {
	var urls []string
	// 建立 Request
	req, err := http.NewRequest("GET", "https://dtcgweb-api.digimoncard.cn/gamecard/gamecardmanager/weblist", nil)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	// 添加参数
	q := req.URL.Query()
	q.Add("page", "")
	q.Add("limit", "300")
	q.Add("name", "")
	q.Add("state", "0")
	q.Add("cardGroup", cardPackage)
	q.Add("rareDegree", "")
	q.Add("belongsType", "")
	q.Add("cardLevel", "")
	q.Add("form", "")
	q.Add("attribute", "")
	q.Add("type", "")
	q.Add("color", "")
	q.Add("envolutionEffect", "")
	q.Add("safeEffect", "")
	q.Add("parallCard", "")
	q.Add("keyEffect", "")
	req.URL.RawQuery = q.Encode()

	// 发起请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	// 解析 JSON 到 struct 中
	var cardDesc models.CardDesc
	err = json.Unmarshal(body, &cardDesc)
	if err != nil {
		logrus.Fatalln(err)
		return nil, err
	}

	for _, mon := range cardDesc.Page.List {
		logrus.Debugln(mon.ImageCover)
		urls = append(urls, mon.ImageCover)
		// i.downloadImage(mon.ImageCover, lang, cardPackage)
	}

	return urls, nil
}

func main() {
	// 指定要下载的图片的语言
	lang := pflag.String("lang", "cn", "language of images")
	pflag.Parse()

	// 获取 cardGroup 列表
	cardPackages, err := services.GetCardPackage()
	if err != nil {
		logrus.Errorf("GetGameCard error: %v", err)
	}
	for _, cardPackage := range cardPackages.List {
		imageHandler := NewImageHandler()

		// 获取图片保存路径
		dir := imageHandler.GetLang(*lang).GetCardPackage(cardPackage.Name).GetDir().Dir

		// 如果目录不存在的话，为每个 cardGroup 的名字按语言创建目录
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, os.ModePerm)
		}

		// 获取下载图片的 URL
		urls, err := imageHandler.GetImagesURL(*lang, cardPackage.Name)
		if err != nil {
			panic(err)
		}

		for _, url := range urls {
			imageHandler.GetLang(*lang).
				GetCardPackage(cardPackage.Name).
				GetFileName(url).
				GetPath().
				downloadImage()
		}
	}
}
