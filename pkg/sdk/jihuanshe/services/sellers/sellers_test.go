package sellers

import (
	"fmt"
	"os"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/sellers/models"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

var (
	token  string = ""
	client *SellersClient
	table  *tablewriter.Table
	// sellerUserID  string = "609077"
	cardVersionID string = "4282"
	productID     string = "33597830"
)

func initConfig() {
	// 初始化配置文件
	c, _ := config.NewConfig("../../../../../config", "")

	// 初始化数据库
	dbInfo := &database.DBInfo{
		FilePath: c.SQLite.FilePath,
		Server:   c.Mysql.Server,
		Password: c.Mysql.Password,
	}

	database.InitDB(dbInfo)

	token = c.JHS.Token
}

func init() {
	initConfig()
	client = NewSellersClient(core.NewClient(token))
	table = tablewriter.NewWriter(os.Stdout)
}

// 添加商品
func TestProductsClientAdd(t *testing.T) {
	initConfig()

	cardPrice, err := database.GetCardPriceWhereCardVersionID(cardVersionID)
	if err != nil {
		logrus.Errorln("获取卡牌价格详情失败", err)
	}

	client := NewSellersClient(core.NewClient(token))
	resp, err := client.ProductAdd(&models.ProductsAddReqBody{
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
	client := NewSellersClient(core.NewClient(token))
	products, err := client.ProductList(fmt.Sprint(currentPage), "", "1", "published_at_desc")
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("共有 %v 个商品", products.Total)

	if products.Total != 0 {
		logrus.WithFields(logrus.Fields{
			"商品ID":   products.Data[0].ProductID,
			"卡牌ID":   products.Data[0].CardVersionID,
			"卡牌名称":   products.Data[0].CardNameCN,
			"售卖价格":   products.Data[0].Price,
			"评级公司ID": products.Data[0].AuthenticatorID,
			"评级公司名称": products.Data[0].AuthenticatorName,
			"评分":     products.Data[0].Grading,
		}).Infof("第一个商品的信息，即刚刚添加的商品信息")

		logrus.Infof("完整信息: %v", products.Data[0])
	}
}

// 更新商品
func TestProductsClientUpdate(t *testing.T) {
	initConfig()

	img := "http://cdn-client.jihuanshe.com/product/2023-02-10-20-25-04-c62Gsu1rOrE9Ea45D1otme3nXxMOEgZbZ1h7PpkD.jpg?imageslim%7CimageMogr2%2Fauto-orient%2Fthumbnail%2F900x%2Fblur%2F1x0%2F%7CimageMogr2%2Fauto-orient%2Fgravity%2FCenter%2Fcrop%2F900x1312%2Fblur%2F1x0%7CimageMogr2%2Fformat%2Fjpg%7Cwatermark%2F2%2Ftext%2F6ZuG5o2i56S-IFVJRDoxMzg1%2Ffont%2F6buR5L2T%2Ffontsize%2F600%2Ffill%2FI0ZGRkZGRg%3D%3D%2Fdissolve%2F90%2Fgravity%2FSouthEast%2Fdx%2F30%2Fdy%2F10"

	client := NewSellersClient(core.NewClient(token))
	resp, err := client.ProductUpdate(&models.ProductsUpdateReqBody{
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
	client := NewSellersClient(core.NewClient(token))

	products, err := client.ProductList("1", "", "1", "published_at_desc")
	if err != nil {
		logrus.Fatal(err)
	}

	productID = fmt.Sprint(products.Data[0].ProductID)

	resp, err := client.ProductDel(productID)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("%v 删除成功：%v", productID, resp)

}

// 获取提现信息
// func TestSellersClientWithdraw(t *testing.T) {
// 	var totalBuyerPrcie float64
// 	page := 1

// 	for {
// 		withdraws, err := client.Withdraw(strconv.Itoa(page))
// 		if err != nil {
// 			logrus.Error(err)
// 		}

// 		for _, withdraw := range withdraws.Data {
// 			money, _ := strconv.ParseFloat(withdraw.Money, 64)
// 			totalBuyerPrcie = totalBuyerPrcie + money
// 		}

// 		logrus.Infof("买入订单共 %v 页，已处理完第 %v 页", withdraws.LastPage, withdraws.CurrentPage)
// 		if withdraws.CurrentPage == withdraws.LastPage {
// 			logrus.Debugf("%v/%v 已处理完成，退出循环", withdraws.CurrentPage, withdraws.LastPage)
// 			break
// 		}

// 		page = withdraws.CurrentPage + 1
// 	}

// 	logrus.Infof("当前已提现总额：%v", totalBuyerPrcie)
// }
