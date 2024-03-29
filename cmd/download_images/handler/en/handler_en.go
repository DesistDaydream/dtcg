package en

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/DesistDaydream/dtcg/cmd/download_images/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/bandai_tcg_plus/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/bandai_tcg_plus/services/card"
	"github.com/sirupsen/logrus"
)

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

// 获取卡包列表
func (i *ImageHandler) GetCardSets() []*handler.CardSetInfo {
	// 获取所有卡包的名称
	cardPackages, err := card.GetCardFilterInfo(&models.CardFilterInfoReq{
		GameTitleID:  "2",
		LanguageCode: i.Lang,
	})

	if err != nil {
		logrus.Errorf("GetGameCard error: %v", err)
	}

	var allCardPackageInfo []*handler.CardSetInfo

	// 排序
	sort.Slice(cardPackages.Success.CardSetList, func(i, j int) bool {
		return cardPackages.Success.CardSetList[i].UpdatedAt.String() < cardPackages.Success.CardSetList[j].UpdatedAt.String()
	})

	for _, cardSet := range cardPackages.Success.CardSetList {
		logrus.WithFields(logrus.Fields{
			"名称": cardSet.Name,
			"ID": cardSet.ID,
			"编号": cardSet.Number,
			// "创建时间": cardSet.CreatedAt.Format("2006-01-02 15:04:05"),
			// "更新时间": cardSet.UpdatedAt.Format("2006-01-02 15:04:05"),
		}).Infof("创建于 %v 更新于 %v 的卡包信息", cardSet.CreatedAt.Format("2006-01-02 15:04:05"), cardSet.UpdatedAt.Format("2006-01-02 15:04:05"))

		// ID 转为 string
		cardPackageID := fmt.Sprintf("%v", cardSet.ID)

		allCardPackageInfo = append(allCardPackageInfo, &handler.CardSetInfo{
			Name: cardSet.Number,
			ID:   cardPackageID,
		})
	}

	// 获取需要下载图片的卡包
	needDownloadCardPackages := handler.GetNeedDownloadCardPackages(allCardPackageInfo)

	return needDownloadCardPackages
}

// 下载卡图
func (i *ImageHandler) DownloadCardImage(needDownloadCardPackages []*handler.CardSetInfo) {
	// 设定过滤条件以获取指定卡片的详情
	c := &models.CardListReqQuery{
		CardSet:     "",
		GameTitleID: "2",
		Limit:       "400",
		Offset:      "0",
	}

	// 循环遍历卡包列表，获取卡包中的卡片
	for _, cardPackage := range needDownloadCardPackages {
		// 生成目录
		dir := handler.GenerateDir(i.DirPrefix, i.Lang, cardPackage.Name)
		// 创建目录
		err := handler.CreateDir(dir)
		if err != nil {
			logrus.Fatalf("为【%v】卡包创建目录失败: %v", cardPackage.Name, err)
		}

		// 设置卡包名称，以过滤条件获取卡片详情
		c.CardSet = cardPackage.ID

		// 获取下载图片的 URL
		urls, err := i.GetImagesURL(c)
		if err != nil {
			panic(err)
		}
		logrus.Infof("准备下载【%v】卡包中的图片，该包中共有 %v 张图片", cardPackage.Name, len(urls))

		// 统计需要下载的图片总量
		handler.Total = handler.Total + len(urls)

		// 下载图片
		for _, url := range urls {
			// 从 URL 中提取文件名
			fileName := i.GenFileName(url)
			// 生成保存图片的绝对路径
			filePath := filepath.Join(dir, fileName)
			err := handler.DownloadImage(url, filePath)
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

// 从卡片详情中获取下载图片所需的 URL
func (i *ImageHandler) GetImagesURL(r *models.CardListReqQuery) ([]string, error) {
	var urls []string

	// 根据过滤条件获取卡片详情
	cardDescs, err := card.GetCardList(r)
	if err != nil {
		return nil, err
	}

	for _, cardDesc := range cardDescs.Success.Cards {
		// logrus.Debugln(mon.ImageCover)
		urls = append(urls, cardDesc.ImageURL)
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
