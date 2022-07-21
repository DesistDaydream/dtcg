package handler

import (
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type ImageHandler interface {
	GetCardPackage(string)
}

// 获取需要保存图片的目录
func GetDir(lang string, cardPackageName string) string {
	dir := "./images/" + lang + "/" + cardPackageName
	return dir
}

// 创建图片保存路径
func CreateDir(cardPackage, lang, cardPackageName string, i ImageHandler) error {
	// 获取图片保存目录
	i.GetCardPackage(cardPackage)
	dir := GetDir(lang, cardPackageName)
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
func DownloadImage(url string, filePath string, fileName string) error {
	// 判断目录中是否有这张图片
	if _, err := os.Stat(filePath); err == nil {
		logrus.Errorf("%v 图片已存在", fileName)
		FailCount++
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

	logrus.Debugln("下载完成")
	SuccessCount++

	return nil
}
