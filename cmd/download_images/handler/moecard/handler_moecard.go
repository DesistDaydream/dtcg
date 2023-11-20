package moecard

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/DesistDaydream/dtcg/cmd/download_images/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"

	"github.com/sirupsen/logrus"
)

var client *cdb.CdbClient

type ImageHandler struct {
	Lang        string
	DirPrefix   string
	LangKeyword string
}

func NewImageHandler(dirPrefix string) handler.ImageHandler {
	return &ImageHandler{
		Lang:        "",
		DirPrefix:   dirPrefix,
		LangKeyword: "ja",
		// LangKeyword: "chs",
	}
}

// 获取卡包列表
func (i *ImageHandler) GetCardSets() []*handler.CardSetInfo {
	var allCardPackageInfo []*handler.CardSetInfo

	// 获取所有卡包的名称
	client = cdb.NewCdbClient(core.NewClient(1, "", 10))
	series, err := client.GetSeries()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, serie := range series.Data {
		for _, pack := range serie.SeriesPack {
			if pack.Language == i.LangKeyword {
				logrus.WithFields(logrus.Fields{
					"前缀": pack.PackPrefix,
					"名称": pack.PackName,
					"ID": pack.PackID,
				}).Infof("%v 中的卡包信息", serie.SeriesName)

				packID := fmt.Sprintf("%v", pack.PackID)

				allCardPackageInfo = append(allCardPackageInfo, &handler.CardSetInfo{
					Name: pack.PackPrefix,
					ID:   packID,
				})
			}
		}
	}

	// 排序
	// sort.Slice(cardPackages.Success.CardSetList, func(i, j int) bool {
	// 	return cardPackages.Success.CardSetList[i].UpdatedAt.String() < cardPackages.Success.CardSetList[j].UpdatedAt.String()
	// })

	// 获取需要下载图片的卡包
	needDownloadCardPackages := handler.GetNeedDownloadCardPackages(allCardPackageInfo)

	return needDownloadCardPackages
}

// 下载卡图
func (i *ImageHandler) DownloadCardImage(needDownloadCardSets []*handler.CardSetInfo) {
	// 循环遍历卡包列表，获取卡包中的卡片
	for _, cardSet := range needDownloadCardSets {
		// 生成目录
		dir := handler.GenerateDir(i.DirPrefix, i.Lang, cardSet.Name)
		// 创建目录
		err := handler.CreateDir(dir)
		if err != nil {
			logrus.Fatalf("为【%v】卡包创建目录失败: %v", cardSet.Name, err)
		}

		name, _ := strconv.Atoi(cardSet.ID)
		// 获取下载图片的 URL
		urls, err := i.GetImagesURL(name)
		if err != nil {
			panic(err)
		}
		logrus.Infof("准备下载【%v】卡包中的图片，该包中共有 %v 张图片", cardSet.Name, len(urls))

		// 统计需要下载的图片总量
		handler.Total = handler.Total + len(urls)

		// 下载图片
		for _, url := range urls {
			// 从 URL 中提取文件名
			fileName := i.GenFileName(url.fileName)
			// 生成保存图片的绝对路径
			filePath := filepath.Join(dir, fileName)
			err := handler.DownloadImage(url.url, filePath)
			if err != nil {
				handler.FailCount++
				logrus.Errorf("下载图片失败: %v", err)
			} else {
				logrus.Debugf("下载到【%v】完成", filePath)
				handler.SuccessCount++
			}
		}

	}
}

type cardImgInfo struct {
	url      string
	fileName string
}

// 从卡片详情中获取下载图片所需的 URL
func (i *ImageHandler) GetImagesURL(cardSet int) ([]cardImgInfo, error) {
	var urls []cardImgInfo

	// 根据过滤条件获取卡片详情
	resp, err := client.PostCardSearch(cardSet, "300", i.LangKeyword, "")
	if err != nil {
		return nil, err
	}

	for _, card := range resp.Data.List {
		image := fmt.Sprintf("https://dtcg-pics.moecard.cn/img/%s~thumb.jpg", card.Images[0].ImgPath)

		// logrus.Debugln(mon.ImageCover)
		urls = append(urls, cardImgInfo{
			url:      image,
			fileName: fmt.Sprintf("%v", card.Images[0].ImgPath),
		})
	}

	return urls, nil
}

// 获取图片语言
func (i *ImageHandler) GetLang(lang string) {
	i.Lang = lang
}

// 从 URL 中提取文件名
func (i *ImageHandler) GenFileName(url string) string {
	// 提取 url 中的文件名
	fileName := url[strings.LastIndex(url, "/")+1:]
	return fileName
}
