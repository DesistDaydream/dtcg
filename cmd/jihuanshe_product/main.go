package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/xuri/excelize/v2"

	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
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

func main() {
	var flags Flags
	AddFlsgs(&flags)
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()
	pflag.Parse()

	if err := logging.LogInit(logFlags.LogLevel, logFlags.LogOutput, logFlags.LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	checkFile(flags.File)

	client := products.NewProductsClient(core.NewClient(flags.Token))

	products, err := client.List("1")
	if err != nil || len(products.Data) <= 0 {
		logrus.Fatalf("获取商品列表失败，列表为空或发生错误：%v", err)
	}

	opts := excelize.Options{}
	f, err := excelize.OpenFile(flags.File, opts)
	if err != nil {
		logrus.Errorln(err)
	}
	sheet := "Sheet1"
	f.NewSheet(sheet)

	// 写入第一行数据，设置列名
	var colNames []string
	desc := &models.ProductList{}
	s := reflect.TypeOf(desc).Elem()
	for i := 0; i < s.NumField(); i++ {
		colNames = append(colNames, s.Field(i).Name)
	}
	for i, colName := range colNames {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, colName)

	}

	// 如果查询到的记录条数大于 pageSize 的值，那么需要分页查询。
	if products.LastPage > 1 {
		for i, row := 1, 2; i <= products.LastPage; i++ {
			products, err = client.List(strconv.Itoa(i))
			if err != nil {
				logrus.Fatalln(err)
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

			row = row + len(products.Data)
		}
	}

	err = f.Save()
	if err != nil {
		logrus.Errorln(err)
	}
}
