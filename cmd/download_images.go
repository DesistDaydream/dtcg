package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DesistDaydream/dtcg/cmd/handler"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/DesistDaydream/dtcg/pkg/subset"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// 从卡片详情中获取下载图片所需的 URL
func GetImagesURL(c *models.FilterConditionReq) ([]string, error) {
	var urls []string

	// 根据过滤条件获取卡片详情
	cardDescs, err := services.GetCardDescs(c)
	if err != nil {
		return nil, err
	}

	for _, cardDesc := range cardDescs.Page.List {
		// logrus.Debugln(mon.ImageCover)
		urls = append(urls, cardDesc.ImageCover)
	}

	return urls, nil
}

// 获取需要下载图片的卡包
func GetNeedDownloadCardPackages(cardPackages *models.CardPackage) []string {
	var allCardPackageNames []string
	for _, cardPackage := range cardPackages.List {
		logrus.WithFields(logrus.Fields{
			"名称": cardPackage.Name,
			"状态": cardPackage.State,
		}).Infof("卡包信息")

		allCardPackageNames = append(allCardPackageNames, cardPackage.Name)
	}
	fmt.Printf("请选择需要下载图片的卡包，多个卡包用逗号分隔(使用 all 下载所有): ")

	// 读取用户输入
	var cardPackagesName string
	fmt.Scanln(&cardPackagesName)

	// 判断用户输入的卡包名称是否存在
	for {
		if cardPackagesName == "all" {
			break
		}
		if !subset.IsSubset(strings.Split(cardPackagesName, ","), allCardPackageNames) {
			fmt.Printf("卡包名称不存在，请重新输入: ")
			fmt.Scanln(&cardPackagesName)
		} else {
			break
		}
	}

	switch cardPackagesName {
	case "all":
		return allCardPackageNames
	default:
		return strings.Split(cardPackagesName, ",")
	}
}

func main() {
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()
	// 指定要下载的图片的语言
	lang := pflag.String("lang", "cn", "language of images")
	pflag.Parse()

	// 初始化日志
	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	// 获取 cardGroup 列表。即获取所有卡包的名称
	cardPackages, err := services.GetCardPackage()
	if err != nil {
		logrus.Errorf("GetGameCard error: %v", err)
	}

	// 获取需要下载图片的卡包
	needDownloadCardPackages := GetNeedDownloadCardPackages(cardPackages)

	// 确认是否要下载
	for _, p := range needDownloadCardPackages {
		logrus.WithField("卡包名称", p).Infof("待下载卡包名称")
	}
	fmt.Printf("需要下载上述卡包，是否继续？(y/n) ")
	var confirm string
	fmt.Scanln(&confirm)
	if confirm != "y" {
		logrus.Infof("取消下载")
		return
	}

	// 循环遍历卡包列表，获取卡包中的卡片
	for _, cardPackage := range needDownloadCardPackages {
		imageHandler := handler.NewImageHandler()

		// 获取图片保存目录
		dir := imageHandler.GetLang(*lang).GetCardPackage(cardPackage).GetDir().Dir
		// 如果目录不存在则创建
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			// 递归创建目录
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				logrus.Fatalf("创建【%v】 目录失败: %v", dir, err)
			}
		}

		// 设定过滤条件以获取指定卡片的详情
		c := &models.FilterConditionReq{
			Page:             "",
			Limit:            "400",
			Name:             "",
			State:            "0",
			CardGroup:        cardPackage,
			RareDegree:       "",
			BelongsType:      "",
			CardLevel:        "",
			Form:             "",
			Attribute:        "",
			Type:             "",
			Color:            "",
			EnvolutionEffect: "",
			SafeEffect:       "",
			ParallCard:       "",
			KeyEffect:        "",
		}

		// 获取下载图片的 URL
		urls, err := GetImagesURL(c)
		if err != nil {
			panic(err)
		}
		logrus.Infof("准备下载【%v】卡包中的图片，该包中共有 %v 张图片", cardPackage, len(urls))

		// 统计需要下载的图片总量
		handler.Total = handler.Total + len(urls)

		// 下载图片
		for _, url := range urls {
			// 前面已经获取到了图片保存的目录，使用 imageHandler.Dir 即可生成图标保存路径
			imageHandler.
				GetFileName(url).
				GetFilePath().
				DownloadImage(url)
		}

	}

	logrus.WithFields(logrus.Fields{
		"总数": handler.Total,
		"成功": handler.SuccessCount,
		"失败": handler.FailCount,
	}).Infof("统计下载结果")
}
