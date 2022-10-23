package services

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community"
)

type Services struct {
	CoreClient *core.Client
	Cdb        *cdb.CdbClient
	Community  *community.CommunityClient
}

func NewServices(token string) *Services {
	s := new(Services)
	s.init(token)
	return s
}

func (s *Services) init(token string) {
	s.CoreClient = core.NewClient(token)
	s.Cdb = cdb.NewCdbClient(s.CoreClient)
	s.Community = community.NewCommunityClient(s.CoreClient)
}
