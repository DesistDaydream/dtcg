package cn

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/DesistDaydream/dtcg/cmd/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/sirupsen/logrus"
)

type ImageHandler struct {
	Lang string
}

func NewImageHandler() handler.ImageHandler {
	return &ImageHandler{
		Lang: "",
	}
}

// 获取卡包列表
func (i *ImageHandler) GetCardPackageList() []*handler.CardPackageInfo {
	// 获取 cardGroup 列表。即获取所有卡包的名称
	cardPackages, err := services.GetCardPackage()
	if err != nil {
		logrus.Errorf("GetGameCard error: %v", err)
	}

	var allCardPackageInfo []*handler.CardPackageInfo

	for _, cardPackage := range cardPackages.List {
		logrus.WithFields(logrus.Fields{
			"名称": cardPackage.Name,
			"状态": cardPackage.State,
		}).Infof("卡包信息")

		allCardPackageInfo = append(allCardPackageInfo, &handler.CardPackageInfo{
			Name:  cardPackage.Name,
			State: cardPackage.State,
		})
	}

	// 获取需要下载图片的卡包
	cardPackageInfo := handler.GetNeedDownloadCardPackages(allCardPackageInfo)

	return cardPackageInfo
}

// 下载卡图
func (i *ImageHandler) DownloadCardImage(needDownloadCardPackages []*handler.CardPackageInfo) {
	// 设定过滤条件以获取指定卡片的详情
	c := &models.FilterConditionReq{
		Page:             "",
		Limit:            "400",
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
		dir := handler.GenerateDir(i.Lang, cardPackageName.Name)
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
			panic(err)
		}
		logrus.Infof("准备下载【%v】卡包中的图片，该包中共有 %v 张图片", cardPackageName.Name, len(urls))

		// 统计需要下载的图片总量
		handler.Total = handler.Total + len(urls)

		// 下载图片
		for _, url := range urls {
			// 从 URL 中提取文件名
			fileName := i.GenFileName(url)
			// 生成保存图片的绝对路径
			filePath := i.GenFilePath(cardPackageName.Name, fileName)
			err := handler.DownloadImage(url, filePath)
			if err != nil {
				logrus.Fatalf("下载图片失败: %v", err)
			}
		}

	}
}

// 从卡片详情中获取下载图片所需的 URL
func (i *ImageHandler) GetImagesURL(c *models.FilterConditionReq) ([]string, error) {
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

// 获取图片保存路径
// 1.获取图片语言
func (i *ImageHandler) GetLang(lang string) {
	i.Lang = lang
}

// 2.从 URL 中提取文件名
func (i *ImageHandler) GenFileName(urlStr string) string {
	// 提取 url 中的文件名
	fileName := urlStr[strings.LastIndex(urlStr, "/")+1:]

	// 将文件名中的汉字解码
	fileName, err := url.QueryUnescape(fileName)
	if err != nil {
		logrus.Errorf("%v 解码文件名失败: %v", fileName, err)
	}

	// 将文件名中的非中文字符替换为空
	match := "[!^\u4e00-\u9fa5]"
	reg := regexp.MustCompile(match)
	newFileName := reg.ReplaceAllString(fileName, "")

	return newFileName
}

// 3.生成图片保存路径
func (i *ImageHandler) GenFilePath(cardPackageName, fileName string) string {
	dir := handler.GenerateDir(i.Lang, cardPackageName) + "/" + fileName
	return dir
}
