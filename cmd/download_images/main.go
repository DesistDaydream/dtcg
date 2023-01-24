package main

import (
	"fmt"
	"runtime"

	"github.com/DesistDaydream/dtcg/cmd/download_images/handler"
	"github.com/DesistDaydream/dtcg/cmd/download_images/handler/cn"
	"github.com/DesistDaydream/dtcg/cmd/download_images/handler/en"
	"github.com/DesistDaydream/dtcg/cmd/download_images/handler/jp"
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	Lang      string
	DirPrefix string
}

func AddFlags(f *Flags) {
	pflag.StringVarP(&f.Lang, "lang", "l", "cn", "图片的语言")
	pflag.StringVarP(&f.DirPrefix, "dir-prefix", "d", "/mnt/d/Projects/dtcg/images", "保存目录的前缀")
}

func main() {
	var (
		flags    Flags
		logFlags logging.LoggingFlags
	)
	AddFlags(&flags)
	logging.AddFlags(&logFlags)
	// 指定要下载的图片的语言

	pflag.Parse()

	// 初始化日志
	if err := logging.LogInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	// 判断当前系统是 Win 还是 Linux
	if runtime.GOOS == "windows" {
		flags.DirPrefix = "D:\\Projects\\dtcg\\images"
	}

	var imageHandler handler.ImageHandler

	switch flags.Lang {
	case "cn":
		imageHandler = cn.NewImageHandler(flags.DirPrefix)
	case "en":
		imageHandler = en.NewImageHandler(flags.DirPrefix)
	case "jp":
		imageHandler = jp.NewImageHandler(flags.DirPrefix)
	default:
		logrus.Fatalln("不支持的语言")
	}

	imageHandler.GetLang(flags.Lang)
	// 获取卡包列表
	cardInfo := imageHandler.GetCardGroups()
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
