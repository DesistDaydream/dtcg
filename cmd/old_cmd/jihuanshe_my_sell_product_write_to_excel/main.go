package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/xuri/excelize/v2"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"

	logging "github.com/DesistDaydream/logging/pkg/logrus_init"
)

func checkFile(rrFile string) {
	if _, err := os.Stat(rrFile); os.IsNotExist(err) {
		logrus.Fatalf("【%v】文件不存在，请使用 -f 指定文件", rrFile)
	}
}

type Flags struct {
	File  string
	Token string
}

func AddFlsgs(f *Flags) {
	pflag.StringVarP(&f.File, "file", "f", "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/我在卖.xlsx", "指定文件")
	pflag.StringVarP(&f.Token, "token", "t", "", "用户认证信息")
	// pflag.StringVarP(&f.File, "file", "f", "test.xlsx", "指定文件")
}

type Orders struct {
}

// 获取所有我在卖的信息，写入到 Excel 中
func main() {
	var (
		flags    Flags
		logFlags logging.LogrusFlags
	)
	AddFlsgs(&flags)
	logging.AddFlags(&logFlags)
	pflag.Parse()

	if err := logging.LogrusInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	checkFile(flags.File)

	client := products.NewProductsClient(core.NewClient(flags.Token))

	opts := excelize.Options{}
	f, err := excelize.OpenFile(flags.File, opts)
	if err != nil {
		logrus.Errorln(err)
	}
	sheet := "Sheet1"
	f.NewSheet(sheet)

	// 写入第一行数据，设置列名
	var colNames []string
	desc := &models.ProductListData{}
	s := reflect.TypeOf(desc).Elem()
	for i := 0; i < s.NumField(); i++ {
		colNames = append(colNames, s.Field(i).Name)
	}
	for i, colName := range colNames {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, colName)

	}

	// 分页查询
	page := 1 // 从获取到的数据的第一页开始
	row := 2  // 从 Excel 表的第二行开始写入数据
	for {
		products, err := client.List(strconv.Itoa(page), "", "1", "published_at_desc")
		if err != nil || len(products.Data) <= 0 {
			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
		}

		for i, product := range products.Data {
			var values []string
			v := reflect.ValueOf(product)
			for k := 0; k < v.NumField(); k++ {
				v := fmt.Sprintf("%v", v.Field(k).Interface())
				values = append(values, v)
			}

			for j, v := range values {
				// 将要写入的数据的行受上一次写入影响
				// 每个订单的产品写入完成后，下一个订单开始写入的行不能从头开始，所以要加上本次订单产品的数量
				// 所以需要最后的 row = row + len(ops.OrderProducts) 这条逻辑
				cell, _ := excelize.CoordinatesToCellName(j+1, i+row)

				f.SetCellValue(sheet, cell, v)
			}
		}

		logrus.Infof("共 %v 页数据，已写完第 %v 页", products.LastPage, products.CurrentPage)
		// 如果当前处理的页等于最后页，则退出循环
		if products.CurrentPage == products.LastPage {
			logrus.Debugf("退出循环时共 %v 页,处理完 %v 页", products.LastPage, products.CurrentPage)
			break
		}

		// 每页商品写入到 Excel 后，写入下一页商品是，不能从第二行开始了，要接着上一页写完的行开始
		row = row + len(products.Data)
		// 每处理完一页，下一个循环需要处理的页+1
		page++
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}
