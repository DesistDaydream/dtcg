package cn

import (
	"net/url"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/DesistDaydream/dtcg/cmd/download_images/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
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
	// 获取 cardGroup 列表。即获取所有卡包的名称
	cardPackages, err := services.GetCardGroups()
	if err != nil {
		logrus.Errorf("获取卡牌集合失败: %v", err)
	}

	var allCardPackageInfo []*handler.CardSetInfo

	// 排序
	sort.Slice(cardPackages.List, func(i, j int) bool {
		return cardPackages.List[i].UpdateTime < cardPackages.List[j].UpdateTime
	})

	for _, cardPackage := range cardPackages.List {
		logrus.WithFields(logrus.Fields{
			"名称": cardPackage.Name,
			"状态": cardPackage.State,
		}).Infof("创建于 %v 更新于 %v 的卡包信息", cardPackage.CreateTime, cardPackage.UpdateTime)

		allCardPackageInfo = append(allCardPackageInfo, &handler.CardSetInfo{
			Name:  cardPackage.Name,
			State: cardPackage.State,
		})
	}

	// 获取需要下载图片的卡包
	cardPackageInfo := handler.GetNeedDownloadCardPackages(allCardPackageInfo)

	return cardPackageInfo
}

// 下载卡图
func (i *ImageHandler) DownloadCardImage(needDownloadCardPackages []*handler.CardSetInfo) {
	// 设定过滤条件以获取指定卡片的详情
	c := &models.FilterConditionReq{
		Page:             "",
		Limit:            "300",
		Name:             "",
		State:            "0",
		CardGroup:        "",
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

	// 循环遍历卡包列表，获取卡包中的卡片
	for _, cardPackageName := range needDownloadCardPackages {
		// 生成目录
		dir := handler.GenerateDir(i.DirPrefix, i.Lang, cardPackageName.Name)
		// 创建目录
		err := handler.CreateDir(dir)
		if err != nil {
			logrus.Fatalf("为【%v】卡包创建目录失败: %v", cardPackageName, err)
		}

		// 设置卡包名称，以过滤条件获取卡片详情
		c.CardGroup = cardPackageName.Name

		// 获取下载图片的 URL
		urls, err := i.GetImagesURL(c)
		if err != nil {
			logrus.Fatalf("获取下载图片的 URL 失败：%v", err)
		}
		logrus.Infof("准备下载【%v】卡包中的图片，该包中共有 %v 张图片", cardPackageName.Name, len(urls))

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
func (i *ImageHandler) GetImagesURL(c *models.FilterConditionReq) ([]string, error) {
	var urls []string

	// 根据过滤条件获取卡片详情
	cardDescs, err := services.GetCardsDesc(c)
	if err != nil {
		return nil, err
	}

	for _, cardDesc := range cardDescs.Page.CardsDesc {
		// logrus.Debugln(mon.ImageCover)
		urls = append(urls, cardDesc.ImageCover)
	}

	return urls, nil
}

// 获取图片语言
func (i *ImageHandler) GetLang(lang string) {
	i.Lang = lang
}

// 从 URL 中提取文件名
func (i *ImageHandler) GenFileName(urlStr string) string {
	// 提取 url 中的文件名
	fileName := urlStr[strings.LastIndex(urlStr, "/")+1:]

	// 将文件名中的汉字解码
	fileName, err := url.QueryUnescape(fileName)
	if err != nil {
		logrus.Errorf("%v 解码文件名失败: %v", fileName, err)
	}

	// 将文件名中的中文字符替换为空
	expr := "[\u4e00-\u9fa5 | \u00a0]"
	reg := regexp.MustCompile(expr)
	newFileName := reg.ReplaceAllString(fileName, "")

	// 将文件名中的时间戳去掉
	timestampExpr := "[0-9]{13}"
	timestampReg := regexp.MustCompile(timestampExpr)
	newFileName2 := timestampReg.ReplaceAllString(newFileName, "")

	return newFileName2
}
