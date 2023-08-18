package services

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/bandai_tcg_plus/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/bandai_tcg_plus/services/card"
)

type Services struct {
	Card *card.CardClient
}

func NewServices(user *models.User) *Services {
	s := new(Services)
	s.init(user.JhsToken)
	return s
}

func (s *Services) init(token string) {
	coreClient := core.NewClient(token)
	card.NewCardClient(coreClient)
}
