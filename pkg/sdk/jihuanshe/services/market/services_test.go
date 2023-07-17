package market

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DesistDaydream/dtcg/config"
	"github.com/DesistDaydream/dtcg/pkg/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
	"github.com/sirupsen/logrus"
)

var token string = ""

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

// 获取用户订单列表（买入）
func TestOrdersClient_GetBuyerOrders(t *testing.T) {
	initConfig()
	client := NewMarketClient(core.NewClient(token))

	resp, err := client.OrderList("1")
	if err != nil {
		logrus.Errorln(err)
	}
	for _, data := range resp.Data {
		logrus.Infoln(data.OrderID)
	}
}

func TestOrdersClient_GetBuyerOrderProducts(t *testing.T) {
	initConfig()
	client := NewMarketClient(core.NewClient(token))
	resp, err := client.OrderGet(2479030)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(resp)
}

// 获取用户订单列表（卖出）
func TestOrdersClientGetSellerOrders(t *testing.T) {
	initConfig()
	client := NewMarketClient(core.NewClient(token))

	resp, err := client.SellerOrderList("1")
	if err != nil {
		logrus.Errorln(err)
	}
	for _, data := range resp.Data {
		logrus.Infoln(data.OrderID)
	}
}

func TestOrdersClientGetSellerOrderProducts(t *testing.T) {
	initConfig()
	client := NewMarketClient(core.NewClient(token))
	resp, err := client.SellerOrderGet(2475268)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Println(resp)
}

func TestMarketClient_GetProductSellers(t *testing.T) {
	client := NewMarketClient(core.NewClient(""))
	got, err := client.CardVersionsProductsGet("2676", "1")
	if err != nil {
		logrus.Errorln(err)
	}

	for _, data := range got.Data {
		fmt.Println(data.CardVersionImage)
	}
}

func TestMarketClient_OrderList(t *testing.T) {
	type fields struct {
		client *core.Client
	}
	type args struct {
		page string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.BuyerOrdersListResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &MarketClient{
				client: tt.fields.client,
			}
			got, err := o.OrderList(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketClient.OrderList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarketClient.OrderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
