package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type ImageHandler interface {
	GetLang(string)
	GetCardPackageList() []*CardPackageInfo
	DownloadCardImage([]*CardPackageInfo)
}

type CardPackageInfo struct {
	Lang  string
	Name  string
	ID    string
	State string
}

// 生成需要保存图片的目录
func GenerateDir(dirPrefix string, lang string, cardPackageName string) string {
	// 生成目录
	dir := filepath.Join(dirPrefix, lang, cardPackageName)
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

// 根据用户输入获取需要下载图片的卡包
func GetNeedDownloadCardPackages(allCardPackageInfo []*CardPackageInfo) []*CardPackageInfo {
	fmt.Printf("请选择需要下载图片的卡包，多个卡包用逗号分隔(使用 all 下载所有): ")

	// 读取用户输入
	var userInputCardPackage string
	var someCardPackageInfo []*CardPackageInfo

	fmt.Scanln(&userInputCardPackage)
	userInputCardPackageSlice := strings.Split(userInputCardPackage, ",")
	for _, name := range userInputCardPackageSlice {
		for _, cardInfo := range allCardPackageInfo {
			if cardInfo.Name == name {
				someCardPackageInfo = append(someCardPackageInfo, cardInfo)
			}
		}
	}

	switch userInputCardPackage {
	case "all":
		return allCardPackageInfo
	default:
		return someCardPackageInfo
	}
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
