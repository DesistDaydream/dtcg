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

func NewServices(token string, retry int) *Services {
	s := new(Services)
	s.init(token, retry)
	return s
}

func (s *Services) init(token string, retry int) {
	s.CoreClient = core.NewClient(token, retry)
	s.Cdb = cdb.NewCdbClient(s.CoreClient)
	s.Community = community.NewCommunityClient(s.CoreClient)
}
