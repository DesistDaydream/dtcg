package main

import (
	"encoding/json"
	"os"
	"reflect"

	"github.com/DesistDaydream/dtcg/cmd/carddesc_json_to_excel/fileparse"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/DesistDaydream/dtcg/pkg/sdk/cn/services/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func checkFile(rrFile string) {
	if _, err := os.Stat(rrFile); os.IsNotExist(err) {
		logrus.Fatalf("【%v】文件不存在，请使用 -f 指定域名的记录规则文件", rrFile)
	}
}

type Flags struct {
	File        string
	DownloadImg bool
}

func AddFlsgs(f *Flags) {
	// pflag.StringVarP(&f.File, "file", "f", "/mnt/e/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/实卡统计.xlsx", "指定文件")
	pflag.StringVarP(&f.File, "file", "f", "/mnt/e/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/统计表.xlsx", "指定文件")
	// pflag.StringVarP(&f.File, "file", "f", "test.xlsx", "指定文件")
	pflag.BoolVarP(&f.DownloadImg, "downloadImg", "d", false, "是否下载图片")
}

func main() {
	var flags Flags
	AddFlsgs(&flags)
	pflag.Parse()

	checkFile(flags.File)

	// 设定初始化过滤条件
	c := &models.FilterConditionReq{
		Page:             "",
		Limit:            "3",
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

	file, err := os.ReadFile("cards/card_package.json")
	if err != nil {
		logrus.Fatalln(err)
	}

	var cardGroups *models.CacheListResp

	err = json.Unmarshal(file, &cardGroups)
	if err != nil {
		logrus.Fatalln(err)
	}

	// 当想要只获取一个或部分卡盒中的信息时，取消注释
	// cardGroups = &models.CacheListResp{
	// 	Msg:  "",
	// 	Code: 0,
	// 	List: []models.CacheList{
	// 		{Name: "STC-01"},
	// 	},
	// }

	for _, cardGroup := range cardGroups.List {
		// 若要获取卡盒所有卡，需要将限制扩大
		// c.Limit = "300"
		c.CardGroup = cardGroup.Name

		// 根据过滤条件获取卡片详情
		cardDescs, err := services.GetCardsDesc(c)
		if err != nil {
			panic(err)
		}

		// 将 JSON 信息中的一部分写入到 Excel 中，可以包含图片
		// fileparse.WriteExcelData(flags.File, cardDescs, cardGroup.Name, flags.DownloadImg)

		// 将 JSON 信息全部写入到 Excel 中
		var colNames []string
		desc := &models.CardDesc{}
		s := reflect.TypeOf(desc).Elem()
		for i := 0; i < s.NumField(); i++ {
			colNames = append(colNames, s.Field(i).Name)
		}
		fileparse.JsonToExcel(flags.File, cardDescs, cardGroup.Name, colNames)
	}

}
