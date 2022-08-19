package main

import (
	"fmt"
	"runtime"

	"github.com/DesistDaydream/dtcg/cmd/download/handler"
	"github.com/DesistDaydream/dtcg/cmd/download/handler/cn"
	"github.com/DesistDaydream/dtcg/cmd/download/handler/en"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()
	// 指定要下载的图片的语言
	lang := pflag.StringP("lang", "l", "cn", "图片的语言")
	dirPrefix := pflag.StringP("dir-prefix", "d", "/mnt/d/Projects/dtcg/images", "保存目录的前缀")
	pflag.Parse()

	// 初始化日志
	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	// 判断当前系统是 Win 还是 Linux
	if runtime.GOOS == "windows" {
		*dirPrefix = "D:\\Projects\\dtcg\\images"
	}

	var imageHandler handler.ImageHandler

	switch *lang {
	case "cn":
		imageHandler = cn.NewImageHandler(*dirPrefix)
	case "en":
		imageHandler = en.NewImageHandler(*dirPrefix)
	default:
		logrus.Fatalln("不支持的语言")
	}

	imageHandler.GetLang(*lang)
	// 获取卡包列表
	cardInfo := imageHandler.GetCardPackageList()
	// 确认是否要下载
	for _, p := range cardInfo {
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
	imageHandler.DownloadCardImage(cardInfo)

	logrus.WithFields(logrus.Fields{
		"总数": handler.Total,
		"成功": handler.SuccessCount,
		"失败": handler.FailCount,
	}).Infof("统计下载结果")
}
