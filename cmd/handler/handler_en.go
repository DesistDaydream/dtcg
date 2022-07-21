package handler

import (
	"strings"
)

type ENImageHandler struct {
	Lang        string
	CardPackage string

	// 文件保存路径相关信息
	Dir      string
	FileName string
	FilePath string
}

func NewENImageHandler() *ENImageHandler {
	return &ENImageHandler{
		Lang:        "",
		CardPackage: "",
		FileName:    "",
	}
}

// 获取图片保存路径
// 1.获取图片语言
func (i *ENImageHandler) GetLang(lang string) *ENImageHandler {
	i.Lang = lang
	return i
}

// 2.获取卡包名称
func (i *ENImageHandler) GetCardPackage(cardPackageName string) *ENImageHandler {
	i.CardPackage = cardPackageName
	return i
}

// 获取需要保存图片的目录
func (i *ENImageHandler) GetDir() *ENImageHandler {
	i.Dir = "./images/" + i.Lang + "/" + i.CardPackage
	return i
}

// 3.获取图片名称
func (i *ENImageHandler) GetFileName(url string) *ENImageHandler {
	// 提取 url 中的文件名
	i.FileName = url[strings.LastIndex(url, "/")+1:]
	return i
}

// 4.获取图片保存路径
func (i *ENImageHandler) GetFilePath() *ENImageHandler {
	i.FilePath = i.Dir + "/" + i.FileName
	return i
}
