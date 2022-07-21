package handler

import (
	"fmt"
	"strings"

	"github.com/DesistDaydream/dtcg/pkg/sdk/en/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/en/services"
	"github.com/DesistDaydream/dtcg/pkg/subset"
	"github.com/sirupsen/logrus"
)

type ENImageHandler struct {
	Lang            string
	CardPackageName string
}

func NewENImageHandler() *ENImageHandler {
	return &ENImageHandler{
		Lang:            "",
		CardPackageName: "",
	}
}

// 获取卡包列表
func (i *ENImageHandler) GetCardPackageList() []string {
	// 获取所有卡包的名称
	cardPackages, err := services.GetCardFilterInfo(&models.CardFilterInfoReq{
		GameTitleID:  "2",
		LanguageCode: i.Lang,
	})

	if err != nil {
		logrus.Errorf("GetGameCard error: %v", err)
	}

	// 获取需要下载图片的卡包
	needDownloadCardPackages := i.GetNeedDownloadCardPackages(cardPackages.Success.CardSetList)

	return needDownloadCardPackages
}

// 获取需要下载图片的卡包
func (i *ENImageHandler) GetNeedDownloadCardPackages(cardPackages []models.CardSetList) []string {
	var allCardPackageNames []string
	for _, cardPackage := range cardPackages {
		logrus.WithFields(logrus.Fields{
			"名称": cardPackage.Name,
			"ID": cardPackage.ID,
			"编号": cardPackage.Number,
		}).Infof("卡包信息")

		// ID 转为 string
		cardPackageID := fmt.Sprintf("%v", cardPackage.ID)

		allCardPackageNames = append(allCardPackageNames, cardPackageID)
	}
	fmt.Printf("请选择需要下载图片的卡包，多个卡包用逗号分隔(使用 all 下载所有，输入ID): ")

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
func (i *ENImageHandler) DownloadCardImage(needDownloadCardPackages []string) {
	// 设定过滤条件以获取指定卡片的详情
	c := &models.CardListReq{
		CardSet:     "",
		GameTitleID: "2",
		Limit:       "400",
		Offset:      "0",
	}

	// 循环遍历卡包列表，获取卡包中的卡片
	for _, cardPackage := range needDownloadCardPackages {
		// TODO: 自动维护这些对应关系
		var cardPackageName string
		switch cardPackage {
		case "23":
			cardPackageName = "BT05"
		case "31":
			cardPackageName = "BT04"
		case "35":
			cardPackageName = "BT01"
		case "36":
			cardPackageName = "BT01"
		case "39":
			cardPackageName = "P"
		case "38":
			cardPackageName = "ST-6"
		case "33":
			cardPackageName = "ST-5"
		case "30":
			cardPackageName = "ST-4"
		case "32":
			cardPackageName = "ST-3"
		case "25":
			cardPackageName = "ST-2"
		case "28":
			cardPackageName = "ST-1"
		default:
			logrus.Fatalln("卡包名称不存在")
		}

		// 生成目录
		dir := GenerateDir(i.Lang, cardPackageName)
		// 创建目录
		err := CreateDir(dir)
		if err != nil {
			logrus.Fatalf("为【%v】卡包创建目录失败: %v", cardPackageName, err)
		}

		// 设置卡包名称，以过滤条件获取卡片详情
		c.CardSet = cardPackage

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
			// 从 URL 中提取文件名
			fileName := i.GenFileName(url)
			// 生成保存图片的绝对路径
			filePath := i.GenFilePath(cardPackageName, fileName)
			err := DownloadImage(url, filePath)
			if err != nil {
				logrus.Fatalf("下载图片失败: %v", err)
			}
		}

	}
}

// 从卡片详情中获取下载图片所需的 URL
func (i *ENImageHandler) GetImagesURL(r *models.CardListReq) ([]string, error) {
	var urls []string

	// 根据过滤条件获取卡片详情
	cardDescs, err := services.GetCardList(r)
	if err != nil {
		return nil, err
	}

	for _, cardDesc := range cardDescs.Success.Cards {
		// logrus.Debugln(mon.ImageCover)
		urls = append(urls, cardDesc.ImageURL)
	}

	return urls, nil
}

// 获取图片保存路径
// 1.获取图片语言
func (i *ENImageHandler) GetLang(lang string) {
	i.Lang = lang
}

// 2.从 URL 中提取文件名
func (i *ENImageHandler) GenFileName(url string) string {
	// 提取 url 中的文件名
	fileName := url[strings.LastIndex(url, "/")+1:]
	return fileName
}

// 3.生成图片保存路径
func (i *ENImageHandler) GenFilePath(cardPackageName, fileName string) string {
	dir := GenerateDir(i.Lang, cardPackageName) + "/" + fileName
	return dir
}
