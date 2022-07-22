package handler

import (
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type ImageHandler interface {
	GetLang(string)
	GetCardPackageList() []*CardInfo
	DownloadCardImage([]*CardInfo)
}

type CardInfo struct {
	Lang  string
	Name  string
	ID    string
	State string
}

// 生成需要保存图片的目录
func GenerateDir(lang string, cardPackageName string) string {
	dir := "./images/" + lang + "/" + cardPackageName
	return dir
}

// 创建图片保存路径
func CreateDir(dir string) error {
	// 如果目录不存在则创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 递归创建目录
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// 统计下载失败和下载成功的次数
var (
	Total        int
	SuccessCount int
	FailCount    int
)

// 下载图片
func DownloadImage(url string, filePath string) error {
	// 判断目录中是否有这张图片
	if _, err := os.Stat(filePath); err == nil {
		// logrus.Errorf("%v 图片已存在", filePath)
		// FailCount++
		return err
	}

	// 下载图片
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {

		return err
	}
	defer file.Close()

	// 写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	logrus.Debugf("下载到【%v】完成", filePath)
	SuccessCount++

	return nil
}
