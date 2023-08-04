package services

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/wishes"
)

type Services struct {
	Market   *market.MarketClient
	Products *products.ProductsClient
	Wishes   *wishes.WishesClient
}

func NewServices(user *models.User) *Services {
	s := new(Services)
	s.init(user.JhsToken)
	return s
}

func (s *Services) init(token string) {
	coreClient := core.NewClient(token)
	s.Market = market.NewMarketClient(coreClient)
	s.Products = products.NewProductsClient(coreClient)
	s.Wishes = wishes.NewWishesClient(coreClient)
}
