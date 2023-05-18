package jp_hk

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/DesistDaydream/dtcg/cmd/download_images/handler"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

var first_url = "https://hk.digimoncard.com/cardlist/?search=true&category=507018"
var base_url = "https://hk.digimoncard.com/cardlist"

type ImageHandler struct {
	Lang      string
	DirPrefix string
}

func NewImageHandler(dirPrefix string) handler.ImageHandler {
	return &ImageHandler{
		Lang:      "",
		DirPrefix: dirPrefix,
	}
}

// 获取卡集列表
func (i *ImageHandler) GetCardSets() []*handler.CardSetInfo {
	var allCardPackageInfo []*handler.CardSetInfo

	client := &http.Client{}
	req, _ := http.NewRequest("GET", first_url, nil)
	resp, _ := client.Do(req)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logrus.Fatal(err)
	}

	// 选择 id 为 snaviList 的 div 元素下的所有 li 元素
	doc.Find("#snaviList li").Each(func(i int, s *goquery.Selection) {
		// 从 li 元素中获取 a 元素的 href 属性值
		href, _ := s.Find("a").Attr("href")
		// 从 li 元素中获取 span 元素的文本内容
		text := s.Find("span.title").Text()
		// 使用字符串切片操作提取出 = 后面的数字
		lastIndex := strings.LastIndex(href, "=")
		number := href[lastIndex+1:]

		logrus.WithFields(logrus.Fields{
			"名称": text,
			"ID": number,
		}).Infof("")

		allCardPackageInfo = append(allCardPackageInfo, &handler.CardSetInfo{
			Name:  text,
			ID:    number,
			State: "",
		})
	})

	// 获取需要下载图片的卡包
	needDownloadCardSets := handler.GetNeedDownloadCardPackages(allCardPackageInfo)

	return needDownloadCardSets
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
			logrus.Fatalf("为【%v】卡包创建目录失败: %v", cardSet, err)
		}

		// 获取下载图片的 URL
		urls, err := i.GetImagesURL(cardSet.ID)
		if err != nil {
			logrus.Fatalf("获取下载图片的 URL 失败：%v", err)
		}
		logrus.Infof("准备下载【%v】卡包中的图片，该包中共有 %v 张图片", cardSet.Name, len(urls))

		// 统计需要下载的图片总量
		handler.Total = handler.Total + len(urls)

		var wg sync.WaitGroup
		defer wg.Wait()

		// 下载图片
		for _, url := range urls {
			// 从 URL 中提取文件名
			fileName := i.GenFileName(url)
			// 生成保存图片的绝对路径
			filePath := filepath.Join(dir, fileName)

			wg.Add(1)
			go func(url string) {
				defer wg.Done()

				err := handler.DownloadImage(url, filePath)
				if err != nil {
					handler.FailCount++
					logrus.Errorf("下载图片失败: %v", err)
				} else {
					logrus.Debugf("下载到【%v】完成", filePath)
					handler.SuccessCount++
				}
			}(url)
		}
	}
}

// 从卡片详情中获取下载图片所需的 URL
func (i *ImageHandler) GetImagesURL(cardSetID string) ([]string, error) {
	var urls []string

	url := fmt.Sprintf("%v?search=true&category=%v", base_url, cardSetID)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(req)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	doc.Find("a.card_img").Each(func(i int, s *goquery.Selection) {
		urlSuffix, _ := s.Find("img").Attr("src")
		// 使用字符串切片操作提取出 / 后面的所有字符
		lastIndex := strings.LastIndex(urlSuffix, "/")
		serial := urlSuffix[lastIndex+1:]

		urls = append(urls, fmt.Sprintf("https://hk.digimoncard.com/images/cardlist/card/%v", serial))
	})

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
