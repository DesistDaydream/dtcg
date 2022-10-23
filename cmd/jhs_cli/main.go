package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
)

var token string

func getToken() {
	file, err := os.ReadFile("pkg/sdk/jihuanshe/services/token.txt")
	if err != nil {
		logrus.Fatal(err)
	}
	token = string(file)
}

// 更新我在卖卡片的卡图
func UpdateImage(productsClient *products.ProductsClient) {
	page := 1 // 从获取到的数据的第一页开始
	for {
		products, err := productsClient.List(strconv.Itoa(page))
		if err != nil || len(products.Data) <= 0 {
			logrus.Fatalf("获取第 %v 页商品失败，列表为空或发生错误：%v", page, err)
		}

		for _, product := range products.Data {
			if !strings.Contains(product.CardVersionImage, "cdn-client") {
				cardPrice, err := database.GetCardPriceWhereCardVersionID(fmt.Sprint(product.CardVersionID))
				if err != nil {
					logrus.Errorln("获取卡牌价格详情失败", err)
				}

				resp, err := productsClient.Update(&models.ProductsUpdateReqBody{
					Condition:            fmt.Sprint(product.Condition),
					OnSale:               fmt.Sprint(product.OnSale),
					Price:                product.Price,
					Quantity:             fmt.Sprint(product.Quantity),
					Remark:               "",
					UserCardVersionImage: cardPrice.ImageUrl,
				}, fmt.Sprint(product.ProductID))
				if err != nil {
					logrus.Errorf("商品 %v %v 修改失败：%v", product.ProductID, product.CardNameCn, err)
				} else {
					logrus.Infof("商品 %v %v 修改成功：%v", product.ProductID, product.CardNameCn, resp)
				}
			}
		}

		logrus.Infof("共 %v 页数据，已处理第 %v 页", products.LastPage, products.CurrentPage)
		// 如果当前处理的页等于最后页，则退出循环
		if products.CurrentPage == products.LastPage {
			logrus.Debugf("退出循环时共 %v 页,处理完 %v 页", products.LastPage, products.CurrentPage)
			break
		}

		// 每处理完一页，下一个循环需要处理的页+1
		page++
	}
}

// 添加商品
func AddProducts(productsClient *products.ProductsClient) {
	cards := []string{"2954", "2955", "2956", "2957"}

	for _, cardVersionID := range cards {
		var price string

		cardPrice, err := database.GetCardPriceWhereCardVersionID(cardVersionID)
		if err != nil {
			logrus.Errorln("获取卡牌价格详情失败", err)
		}

		if cardPrice.AvgPrice == 0 {
			price = "10"
		} else {
			price = fmt.Sprint(cardPrice.AvgPrice + float64(2))
		}

		// 开始上架
		resp, err := productsClient.Add(&models.ProductsAddReqBody{
			CardVersionID:        cardVersionID,
			Price:                price,
			Quantity:             "4",
			Condition:            "1",
			Remark:               "",
			GameKey:              "dgm",
			UserCardVersionImage: cardPrice.ImageUrl,
		})

		// resp, err := client.Add(cardModelToCardVersionID[rows[i][0]], rows[i][1], rows[i][2])
		if err != nil {
			logrus.Errorf("%v 上架失败：%v", cardPrice.ScName, err)
		} else {
			logrus.Infof("%v 上架成功：%v", cardPrice.ScName, resp)
		}
	}
}

func main() {
	dbInfo := &database.DBInfo{
		FilePath: "internal/database/my_dtcg.db",
		Server:   "122.9.154.106:3306",
		Password: "lch1382121",
	}
	database.InitDB(dbInfo)

	getToken()
	productsClient := products.NewProductsClient(core.NewClient(token))

	AddProducts(productsClient)
	// UpdateImage(productsClient)
}
