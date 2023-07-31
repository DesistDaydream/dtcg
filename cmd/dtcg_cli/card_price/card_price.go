package cardprice

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cardPriceCmd := &cobra.Command{
		Use:              "card-price",
		Short:            "控制卡片价格信息",
		PersistentPreRun: cardPricePersistentPreRun,
	}

	cardPriceCmd.AddCommand(
		AddCardPriceCommand(),
		UpdateCardPriceCommand(),
	)

	return cardPriceCmd
}

func cardPricePersistentPreRun(cmd *cobra.Command, args []string) {
	// 执行根命令的初始化操作
	parent := cmd.Parent()
	if parent.PersistentPreRun != nil {
		parent.PersistentPreRun(parent, args)
	}
}

// 从 DtcgDB 获取卡牌价格
func GetPriceFromDtcgdb(cardDesc *models.CardDesc) (int, float64, float64) {
	var (
		minPrice      float64
		avgPrice      float64
		cardVersionID int
	)

	cardPrice, err := handler.H.MoecardServices.Cdb.GetCardPrice(fmt.Sprint(cardDesc.CardIDFromDB))
	if err != nil {
		logrus.Errorf("获取卡牌 %v 价格失败: %v", cardDesc.CardIDFromDB, err)
		return 0, 0, 0
	}

	avgPrice, _ = strconv.ParseFloat(cardPrice.Data.AvgPrice, 64)
	cardVersionID = cardPrice.Data.CardVersionID

	if cardPrice.Data.Total == 0 {
		minPrice = 0
	} else {
		minPrice, _ = strconv.ParseFloat(cardPrice.Data.Data[0].MinPrice, 64)
	}

	if avgPrice == 0 {
		avgPrice = minPrice
	}

	return cardVersionID, minPrice, avgPrice
}

// 从 集换社 获取卡牌价格
func GetPriceFromJhs(cardDesc *models.CardDesc) (int, float64, float64) {
	var (
		minPrice      float64
		avgPrice      float64
		cardVersionID int
	)

	// 获取 cardVersionID
	// TODO: 获取失败时，记录错误次数
	cardPrice, err := database.GetCardPrice(fmt.Sprint(cardDesc.CardIDFromDB))
	if err != nil {
		logrus.Errorf("获取 card_version_id 失败: %v", err)
		return 0, 0, 0
	}

	cardVersionID = cardPrice.CardVersionID

	productInfo, err := handler.H.JhsServices.Products.Get(fmt.Sprint(cardPrice.CardVersionID), "1")
	if err != nil {
		logrus.Errorf("获取卡牌 %v 价格失败: %v", cardDesc.CardIDFromDB, err)
		return cardVersionID, 0, 0
	}

	minPrice, _ = strconv.ParseFloat(productInfo.MinPrice, 64)
	avgPrice, _ = strconv.ParseFloat(productInfo.AvgPrice, 64)

	if avgPrice == 0 {
		avgPrice = minPrice
	}

	// 防止请求太快，等待 0.5 秒
	time.Sleep(time.Millisecond * 500)

	return cardVersionID, minPrice, avgPrice
}

// 获取集换社卡牌的图片
func GetImageURL(cardVersionID int) string {
	page := 1
	// 分页
	for {
		productSellers, err := handler.H.JhsServices.Market.CardVersionsProductsGet(fmt.Sprint(cardVersionID), fmt.Sprint(page))
		if err != nil || productSellers.Total < 1 {
			logrus.Errorf("获取商品 %v 在售清单异常: %v", cardVersionID, err)
			return ""
		}

		for _, d := range productSellers.Data {
			if strings.Contains(d.CardVersionImage, "cdn-client") {
				logrus.Debugf("获取卡图成功")
				return d.CardVersionImage
			}
		}

		logrus.Debugf("商品在售清单共 %v 页，已处理完第 %v 页", productSellers.LastPage, productSellers.CurrentPage)
		if productSellers.CurrentPage == productSellers.LastPage {
			logrus.Debugf("%v/%v 已处理完成，退出循环", productSellers.CurrentPage, productSellers.LastPage)
			return ""
		}

		page = productSellers.CurrentPage + 1
	}
}
