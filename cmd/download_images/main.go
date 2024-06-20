package main

import (
	"fmt"
	"runtime"

	logging "github.com/DesistDaydream/logging/pkg/logrus_init"

	"github.com/DesistDaydream/dtcg/cmd/download_images/handler"
	"github.com/DesistDaydream/dtcg/cmd/download_images/handler/cn"
	"github.com/DesistDaydream/dtcg/cmd/download_images/handler/en"
	"github.com/DesistDaydream/dtcg/cmd/download_images/handler/jp"
	"github.com/DesistDaydream/dtcg/cmd/download_images/handler/moecard"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Flags struct {
	Lang           string
	DirPrefix      string
	CardUpdateTime string
}

func AddFlags(f *Flags) {
	pflag.StringVarP(&f.Lang, "lang", "l", "cn", "图片的语言")
	pflag.StringVarP(&f.DirPrefix, "dir-prefix", "d", "/mnt/d/Projects/dtcg/images", "保存目录的前缀")
	pflag.StringVar(&f.CardUpdateTime, "card-update-time", "2024-05-13", "卡牌更新时间。主要针对中文宣传卡，只有指定的时间之后的卡图才会下载。因为中文官网所有宣传卡都放在一起没有分类")
}

func main() {
	var (
		flags    Flags
		logFlags logging.LogrusFlags
	)
	AddFlags(&flags)
	logging.AddFlags(&logFlags)
	// 指定要下载的图片的语言

	pflag.Parse()

	// 初始化日志
	if err := logging.LogrusInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	// 判断当前系统是 Win 还是 Linux
	if runtime.GOOS == "windows" {
		flags.DirPrefix = "D:\\Projects\\dtcg\\images"
	}

	var imageHandler handler.ImageHandler

	switch flags.Lang {
	case "cn":
		imageHandler = cn.NewImageHandler(flags.DirPrefix, flags.CardUpdateTime)
	case "en":
		imageHandler = en.NewImageHandler(flags.DirPrefix)
	case "jp":
		imageHandler = jp.NewImageHandler(flags.DirPrefix)
	case "moecard":
		imageHandler = moecard.NewImageHandler(flags.DirPrefix)
	default:
		logrus.Fatalln("不支持的语言")
	}

	imageHandler.GetLang(flags.Lang)
	// 获取卡包列表
	cardSetInfo := imageHandler.GetCardSets()
	// 确认是否要下载
	for _, p := range cardSetInfo {
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
	imageHandler.DownloadCardImage(cardSetInfo)

	logrus.WithFields(logrus.Fields{
		"总数": handler.Total,
		"成功": handler.SuccessCount,
		"失败": handler.FailCount,
	}).Infof("统计下载结果")
}
