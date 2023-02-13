package products

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
)

var sellerUserID string = "609077"
var token string = ""
var cardVersionID string = "3982"
var productID string = ""

func initConfig() {
	// 初始化配置文件
	c := config.NewConfig("../../../../../config", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	token = c.JHS.Token
}

// 测试将结构体转为 map
func TestStructToMapStr(t *testing.T) {
	obj := models.ProductsGetReqQuery{
		GameKey:       "dgm",
		SellerUserID:  sellerUserID,
		CardVersionID: cardVersionID,
	}

	got := core.StructToMapStr(&obj)

	fmt.Println(len(got))

	gotByte, _ := json.Marshal(got)
	fmt.Println(string(gotByte))
	for k, v := range got {
		fmt.Println(k, v)
	}

}

// 添加商品
func TestProductsClientAdd(t *testing.T) {
	initConfig()

	cardPrice, err := database.GetCardPriceWhereCardVersionID(cardVersionID)
	if err != nil {
		logrus.Errorln("获取卡牌价格详情失败", err)
	}

	client := NewProductsClient(core.NewClient(token))
	resp, err := client.Add(&models.ProductsAddReqBody{
		AuthenticatorID:         "",
		Grading:                 "",
		CardVersionID:           cardVersionID,
		Condition:               "1",
		GameKey:                 "dgm",
		Price:                   "1234.56",
		ProductCardVersionImage: cardPrice.ImageUrl,
		Quantity:                "3",
		Remark:                  "测试卡牌",
	})

	if err != nil {
		logrus.Errorf("%v 上架失败：%v", cardPrice.ScName, err)
	} else {
		logrus.Infof("%v 上架成功：%v", cardPrice.ScName, resp)
	}
}

// 列出商品
func TestProductsClientList(t *testing.T) {
	initConfig()
	currentPage := 1
	client := NewProductsClient(core.NewClient(token))
	products, err := client.List(fmt.Sprint(currentPage), "", "1")
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("共有 %v 个商品", products.Total)

	logrus.WithFields(logrus.Fields{
		"商品ID":   products.Data[0].ProductID,
		"卡牌ID":   products.Data[0].CardVersionID,
		"卡牌名称":   products.Data[0].CardNameCn,
		"售卖价格":   products.Data[0].Price,
		"评级公司ID": products.Data[1].AuthenticatorID,
		"评级公司名称": products.Data[1].AuthenticatorName,
		"评分":     products.Data[1].Grading,
	}).Infof("第一个商品的信息，即刚刚添加的商品信息")

	logrus.Infof("完整信息: %v", products.Data[1])

	productID = fmt.Sprint(products.Data[0].ProductID)
}

// 更新商品
func TestProductsClientUpdate(t *testing.T) {
	initConfig()

	productID = "22075072"
	img := "http://cdn-client.jihuanshe.com/product/2023-02-10-20-25-04-c62Gsu1rOrE9Ea45D1otme3nXxMOEgZbZ1h7PpkD.jpg?imageslim%7CimageMogr2%2Fauto-orient%2Fthumbnail%2F900x%2Fblur%2F1x0%2F%7CimageMogr2%2Fauto-orient%2Fgravity%2FCenter%2Fcrop%2F900x1312%2Fblur%2F1x0%7CimageMogr2%2Fformat%2Fjpg%7Cwatermark%2F2%2Ftext%2F6ZuG5o2i56S-IFVJRDoxMzg1%2Ffont%2F6buR5L2T%2Ffontsize%2F600%2Ffill%2FI0ZGRkZGRg%3D%3D%2Fdissolve%2F90%2Fgravity%2FSouthEast%2Fdx%2F30%2Fdy%2F10"

	client := NewProductsClient(core.NewClient(token))
	resp, err := client.Update(&models.ProductsUpdateReqBody{
		AuthenticatorID: "",
		Grading:         "",
		Condition:       "1",
		// Default:                 "",
		OnSale:                  "1",
		Price:                   "2500.10",
		ProductCardVersionImage: img,
		Quantity:                "9",
		Remark:                  "测试更新卡牌",
	}, productID)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("%v 更新成功：%v", productID, resp)
}

// 删除商品
func TestProductsClientDel(t *testing.T) {
	initConfig()
	client := NewProductsClient(core.NewClient(token))

	ProductIDs := []string{productID}
	for _, ProductID := range ProductIDs {
		resp, err := client.Del(ProductID)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("%v 删除成功：%v", productID, resp)
	}
}

// 获取商品信息
func TestProductsClientGet(t *testing.T) {
	initConfig()
	client := NewProductsClient(core.NewClient(token))

	got, err := client.Get(cardVersionID, sellerUserID)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(got)
}
