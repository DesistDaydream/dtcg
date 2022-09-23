package orders

import (
	"fmt"
	"os"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/sirupsen/logrus"
)

var token string = ""

// var cardVersionID string = "2544"

func getToken() {
	file, err := os.ReadFile("token.txt")
	if err != nil {
		logrus.Fatal(err)
	}
	token = string(file)
}

func TestOrdersClient_GetBuyerOrders(t *testing.T) {
	getToken()
	client := NewOrdersClient(core.NewClient(token))

	resp, err := client.GetBuyerOrders("1")
	if err != nil {
		logrus.Errorln(err)
	}
	for _, data := range resp.Data {
		logrus.Infoln(data.OrderID)
	}
}

func TestOrdersClient_GetBuyerOrderProducts(t *testing.T) {
	getToken()
	client := NewOrdersClient(core.NewClient(token))
	resp, err := client.GetBuyerOrderProducts(2479030)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(resp)
}

func TestOrdersClient_GetSellerOrders(t *testing.T) {
	getToken()
	client := NewOrdersClient(core.NewClient(token))

	resp, err := client.GetSellerOrders("1")
	if err != nil {
		logrus.Errorln(err)
	}
	for _, data := range resp.Data {
		logrus.Infoln(data.OrderID)
	}
}

func TestOrdersClient_GetSellerOrderProducts(t *testing.T) {
	getToken()
	client := NewOrdersClient(core.NewClient(token))
	resp, err := client.GetSellerOrderProducts(2475268)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(resp)
}
