package main

import (
	"fmt"

	"github.com/DesistDaydream/dtcg/cmd/handler"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

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

	switch *lang {
	case "cn":
		imageHandler := handler.NewCNImageHandler()
		imageHandler.GetLang(*lang)
		// 获取卡包列表
		needDownloadCardPackages := imageHandler.GetCardPackageList()
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

		// 下载
		imageHandler.DownloadCardImage(needDownloadCardPackages)
	case "en":
		imageHandler := handler.NewENImageHandler()
		imageHandler.GetLang(*lang)
		// 获取卡包列表
		needDownloadCardPackages := imageHandler.GetCardPackageList()
		// 确认是否要下载
		for _, p := range needDownloadCardPackages {
			logrus.WithField("卡包名称", p.Name).Infof("待下载卡包名称")
		}
		fmt.Printf("需要下载上述卡包，是否继续？(y/n) ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "y" {
			logrus.Infof("取消下载")
			return
		}

		// 下载
		imageHandler.DownloadCardImage(needDownloadCardPackages)
	default:
		logrus.Fatalln("不支持的语言")
	}

	logrus.WithFields(logrus.Fields{
		"总数": handler.Total,
		"成功": handler.SuccessCount,
		"失败": handler.FailCount,
	}).Infof("统计下载结果")
}
