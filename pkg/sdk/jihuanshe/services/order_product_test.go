package services

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetOrderProducts(t *testing.T) {
	token := ""
	orderProducts, err := GetOrderProducts(2469672, token)
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, orderProduct := range orderProducts.OrderProducts {
		logrus.WithFields(logrus.Fields{
			"名称":  orderProduct.CardNameCn,
			"版本号": orderProduct.CardVersionNumber,
			"稀有度": orderProduct.CardVersionRarity,
			"价格":  orderProduct.Price,
		}).Infoln("销售商品信息")
	}
	fmt.Println()

}
