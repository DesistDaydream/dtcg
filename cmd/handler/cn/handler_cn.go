package cn

import (
	"fmt"
	"strings"

	"github.com/DesistDaydream/dtcg/cmd/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/DesistDaydream/dtcg/pkg/subset"
	"github.com/sirupsen/logrus"
)

type CNImageHandler struct {
	Lang            string
	CardPackageName string
}

func NewCNImageHandler() *CNImageHandler {
	return &CNImageHandler{
		Lang:            "",
		CardPackageName: "",
	}
}

// 获取卡包列表
func (i *CNImageHandler) GetCardPackageList() []string {
	// 获取 cardGroup 列表。即获取所有卡包的名称
	cardPackages, err := services.GetCardPackage()
	if err != nil {
		logrus.Errorf("GetGameCard error: %v", err)
	}

	// 获取需要下载图片的卡包
	needDownloadCardPackages := i.GetNeedDownloadCardPackages(cardPackages)

	return needDownloadCardPackages
}

// 获取需要下载图片的卡包
func (i *CNImageHandler) GetNeedDownloadCardPackages(cardPackages *models.CardPackage) []string {
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

// 下载卡图
func (i *CNImageHandler) DownloadCardImage(needDownloadCardPackages []string) {
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
		dir := handler.GenerateDir(i.Lang, cardPackageName)
		// 创建目录
		err := handler.CreateDir(dir)
		if err != nil {
			logrus.Fatalf("为【%v】卡包创建目录失败: %v", cardPackageName, err)
		}

		// 设置卡包名称，以过滤条件获取卡片详情
		c.CardGroup = cardPackageName

		// 获取下载图片的 URL
		urls, err := i.GetImagesURL(c)
		if err != nil {
			panic(err)
		}
		logrus.Infof("准备下载【%v】卡包中的图片，该包中共有 %v 张图片", cardPackageName, len(urls))

		// 统计需要下载的图片总量
		handler.Total = handler.Total + len(urls)

		// 下载图片
		for _, url := range urls {
			// 从 URL 中提取文件名
			fileName := i.GenFileName(url)
			// 生成保存图片的绝对路径
			filePath := i.GenFilePath(cardPackageName, fileName)
			err := handler.DownloadImage(url, filePath)
			if err != nil {
				logrus.Fatalf("下载图片失败: %v", err)
			}
		}

	}
}

// 从卡片详情中获取下载图片所需的 URL
func (i *CNImageHandler) GetImagesURL(c *models.FilterConditionReq) ([]string, error) {
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
func (i *CNImageHandler) GetLang(lang string) {
	i.Lang = lang
}

// 2.从 URL 中提取文件名
func (i *CNImageHandler) GenFileName(url string) string {
	// 提取 url 中的文件名
	fileName := url[strings.LastIndex(url, "/")+1:]
	return fileName
}

// 3.生成图片保存路径
func (i *CNImageHandler) GenFilePath(cardPackageName, fileName string) string {
	dir := handler.GenerateDir(i.Lang, cardPackageName) + "/" + fileName
	return dir
}
