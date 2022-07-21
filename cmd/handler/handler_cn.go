package handler

import (
	"fmt"
	"strings"

	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/DesistDaydream/dtcg/pkg/subset"
	"github.com/sirupsen/logrus"
)

type CNImageHandler struct {
	Lang            string
	CardPackageName string

	// 文件保存路径相关信息
	Dir      string
	FileName string
	FilePath string
}

func NewCNImageHandler() *CNImageHandler {
	return &CNImageHandler{
		Lang:            "",
		CardPackageName: "",
		FileName:        "",
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
	needDownloadCardPackages := GetNeedDownloadCardPackages(cardPackages)

	return needDownloadCardPackages
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

// 下载卡图
func (i *CNImageHandler) DownloadCardImage(needDownloadCardPackages []string) {
	// 循环遍历卡包列表，获取卡包中的卡片
	for _, cardPackage := range needDownloadCardPackages {
		// 创建目录
		err := CreateDir(cardPackage, i.Lang, i.CardPackageName, i)
		if err != nil {
			logrus.Fatalf("为【%v】卡包创建目录失败: %v", cardPackage, err)
		}

		// 设定过滤条件以获取指定卡片的详情
		c := &models.FilterConditionReq{
			Page:             "",
			Limit:            "400",
			Name:             "",
			State:            "0",
			CardGroup:        cardPackage,
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

		// 获取下载图片的 URL
		urls, err := i.GetImagesURL(c)
		if err != nil {
			panic(err)
		}
		logrus.Infof("准备下载【%v】卡包中的图片，该包中共有 %v 张图片", cardPackage, len(urls))

		// 统计需要下载的图片总量
		Total = Total + len(urls)

		// 下载图片
		for _, url := range urls {
			// 前面已经获取到了图片保存的目录，使用 imageHandler.Dir 即可生成图标保存路径
			fileName := i.GetFileName(url)
			err := DownloadImage(url, i.GetFilePath(fileName), fileName)
			if err != nil {
				logrus.Fatalf("下载图片失败: %v", err)
			}
		}

	}
}

// 获取图片保存路径
// 1.获取图片语言
func (i *CNImageHandler) GetLang(lang string) {
	i.Lang = lang
}

// 2.获取卡包名称
func (i *CNImageHandler) GetCardPackage(cardPackageName string) {
	i.CardPackageName = cardPackageName
}

// 3.获取图片名称
func (i *CNImageHandler) GetFileName(url string) string {
	// 提取 url 中的文件名
	fileName := url[strings.LastIndex(url, "/")+1:]
	return fileName
}

// 4.获取图片保存路径
func (i *CNImageHandler) GetFilePath(fileName string) string {
	dir := GetDir(i.Lang, i.CardPackageName) + "/" + fileName
	return dir
}

// 获取需要下载图片的卡包
func GetNeedDownloadCardPackages(cardPackages *models.CardPackage) []string {
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
