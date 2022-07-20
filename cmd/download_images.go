package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/DesistDaydream/dtcg/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type ImageHandler struct {
	Lang        string
	CardPackage string

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
	}
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

// 获取需要保存图片的目录
func (i *ImageHandler) GetDir() *ImageHandler {
	i.Dir = "./images/" + i.Lang + "/" + i.CardPackage
	return i
}

// 3.获取图片名称
func (i *ImageHandler) GetFileName(url string) *ImageHandler {
	// 提取 url 中的文件名
	i.FileName = url[strings.LastIndex(url, "/")+1:]
	return i
}

// 4.获取图片保存路径
func (i *ImageHandler) GetFilePath() *ImageHandler {
	i.FilePath = i.Dir + "/" + i.FileName
	return i
}

// 统计下载失败和下载成功的次数
var (
	total        int
	successCount int
	failCount    int
)

// 下载图片
func (i *ImageHandler) downloadImage(url string) {
	// 判断目录中是否有这张图片
	if _, err := os.Stat(i.FilePath); err == nil {
		logrus.Errorf("%v 图片已存在", i.FileName)
		failCount++
		return
	}

	// 下载图片
	resp, err := http.Get(url)
	if err != nil {
		logrus.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	// 创建文件
	file, err := os.Create(i.FilePath)
	if err != nil {
		logrus.Fatalln(err)
		return
	}
	defer file.Close()

	// 写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		logrus.Fatalln(err)
		return
	}

	logrus.Infof("下载完成")
	successCount++
}

func GetImagesURL(cardPackage string) ([]string, error) {
	var urls []string
	// 根据过滤条件获取卡片详情
	cardDesc, err := services.GetCardDesc(cardPackage)
	if err != nil {
		return nil, err
	}

	for _, mon := range cardDesc.Page.List {
		logrus.Debugln(mon.ImageCover)
		urls = append(urls, mon.ImageCover)
	}

	return urls, nil
}

func main() {
	logging := logging.LoggingFlags{}
	logging.AddFlags()
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

		// 获取图片保存目录
		dir := imageHandler.GetLang(*lang).GetCardPackage(cardPackage.Name).GetDir().Dir
		// 如果目录不存在则创建
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			// 递归创建目录
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				logrus.Fatalf("创建【%v】 目录失败: %v", dir, err)
			}
		}

		// 获取下载图片的 URL
		urls, err := GetImagesURL(cardPackage.Name)
		if err != nil {
			panic(err)
		}

		// 统计需要下载的图片总量
		total = total + len(urls)

		// 下载图片
		for _, url := range urls {
			// 前面已经获取到了图片保存的目录，使用 imageHandler.Dir 即可生成图标保存路径
			imageHandler.
				GetFileName(url).
				GetFilePath().
				downloadImage(url)
		}

	}

	logrus.WithFields(logrus.Fields{
		"total":   total,
		"success": successCount,
		"fail":    failCount,
	}).Infof("统计下载结果")
}
