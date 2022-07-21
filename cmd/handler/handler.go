package handler

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// 统计下载失败和下载成功的次数
var (
	Total        int
	SuccessCount int
	FailCount    int
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

// 下载图片
func (i *ImageHandler) DownloadImage(url string) {
	// 判断目录中是否有这张图片
	if _, err := os.Stat(i.FilePath); err == nil {
		logrus.Errorf("%v 图片已存在", i.FileName)
		FailCount++
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

	logrus.Debugln("下载完成")
	SuccessCount++
}
