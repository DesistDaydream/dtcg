package products

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

var token string = ""
var cardVersionID string = "2544"

func getToken() {
	file, err := os.ReadFile("token.txt")
	if err != nil {
		logrus.Fatal(err)
	}
	token = string(file)
}

func TestStructToMapStr(t *testing.T) {
	obj := models.ProductsGetRequestQuery{
		GameKey:       "dgm",
		SellerUserID:  "609077",
		CardVersionID: cardVersionID,
		Token:         token,
	}

	got := StructToMapStr(&obj)

	fmt.Println(len(got))

	gotByte, _ := json.Marshal(got)
	fmt.Println(string(gotByte))
	for k, v := range got {
		fmt.Println(k, v)
	}

}

func TestProductsClientAdd(t *testing.T) {
	cardModelToCardVersionID := make(map[string]string)
	filea, err := os.ReadFile("/mnt/d/Projects/DesistDaydream/dtcg/cards/card_version_id_and_card_modle.json")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = json.Unmarshal(filea, &cardModelToCardVersionID)
	if err != nil {
		logrus.Fatalln(err)
	}

	getToken()
	client := NewProductsClient(core.NewClient(token))
	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/价格统计表.xlsx"
	f, err := excelize.OpenFile(file)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			logrus.Errorln(err)
			return
		}
	}()
	sheet := "待上架"
	rows, err := f.GetRows(sheet)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"sheet": sheet,
		}).Fatalf("读取中sheet页异常: %v", err)
	}

	for i := 1; i < len(rows); i++ {
		logrus.Infof("%v 的集换社 ID 为 %v", rows[i][0], cardModelToCardVersionID[rows[i][0]])

		// 开始上架
		resp, err := client.Add(&models.ProductsAddRequestBody{
			CardVersionID:        cardModelToCardVersionID[rows[i][0]],
			Price:                rows[i][1],
			Quantity:             rows[i][2],
			Condition:            "1",
			Remark:               "",
			GameKey:              "dgm",
			UserCardVersionImage: "",
		})

		// resp, err := client.Add(cardModelToCardVersionID[rows[i][0]], rows[i][1], rows[i][2])
		if err != nil {
			logrus.Errorf("%v 上架失败：%v", rows[i][0], err)
		} else {
			logrus.Infof("%v 上架成功：%v", rows[i][0], resp)
		}
	}
}

func TestProductsClientDel(t *testing.T) {
	getToken()
	client := NewProductsClient(core.NewClient(token))

	ProductIDs := []string{"15526516", "15526515", "15526514", "15526513", "15526511", "15526510", "15526509", "15526507", "15526506", "15526505", "15526504", "15525782"}
	for _, ProductID := range ProductIDs {
		got, err := client.Del(ProductID)
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Println(got)
	}
}

func TestProductsClientUpdate(t *testing.T) {
	getToken()
	client := NewProductsClient(core.NewClient(token))
	file := "/mnt/d/Documents/WPS Cloud Files/1054253139/团队文档/东部王国/数码宝贝/我在卖.xlsx"
	f, err := excelize.OpenFile(file)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			logrus.Errorln(err)
			return
		}
	}()
	sheet := "Sheet1"
	rows, err := f.GetRows(sheet)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"sheet": sheet,
		}).Fatalf("读取中sheet页异常: %v", err)
	}

	// for i := 1; i < len(rows); i++ {
	for i := 1; i < 3; i++ {
		resp, err := client.Update(&models.ProductsUpdateRequestBody{
			Condition:            "1",
			OnSale:               "1",
			Price:                rows[i][1],
			Quantity:             rows[i][2],
			Remark:               "",
			UserCardVersionImage: rows[i][13],
		}, rows[i][0])
		if err != nil {
			logrus.Errorf("商品 %v %v 修改失败：%v", rows[i][10], rows[i][8], err)
		} else {
			logrus.Infof("商品 %v %v 修改成功：%v", rows[i][10], rows[i][8], resp)
		}
	}

}

func TestProductsClientGet(t *testing.T) {
	getToken()
	client := NewProductsClient(core.NewClient(token))

	got, err := client.Get(cardVersionID)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(got)
}

func TestProductsClientList(t *testing.T) {
	getToken()
	currentPage := 1
	for {
		client := NewProductsClient(core.NewClient(token))
		products, err := client.List(fmt.Sprint(currentPage))
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Println(products.CurrentPage)

		if products.CurrentPage == products.LastPage {
			break
		}

		currentPage++
	}
}
