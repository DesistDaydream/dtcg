package services

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/orders"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
)

type Services struct {
	CoreClient *core.Client
	Market     *market.MarketClient
	Orders     *orders.OrdersClient
	Products   *products.ProductsClient
}

func NewServices(token string) *Services {
	s := new(Services)
	s.init(token)
	return s
}

func (s *Services) init(token string) {
	s.CoreClient = core.NewClient(token)
	s.Market = market.NewMarketClient(s.CoreClient)
	s.Orders = orders.NewOrdersClient(s.CoreClient)
	s.Products = products.NewProductsClient(s.CoreClient)
}
