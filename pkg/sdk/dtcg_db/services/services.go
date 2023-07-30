package services

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community"
	"github.com/sirupsen/logrus"
)

type Services struct {
	Cdb       *cdb.CdbClient
	Community *community.CommunityClient
}

func NewServices(isLogin bool, username, password, token string, retry int) *Services {
	if isLogin {
		if core.CheckToken(token) {
			logrus.Infoln("TOKEN 可用，不用重新获取")
		} else {
			logrus.Warnln("TOKEN 不可用，开始重新获取")
			token = core.Login(username, password)
		}
	} else {
		token = ""
	}

	s := new(Services)
	s.init(token, retry)
	return s
}

func (s *Services) init(token string, retry int) {
	coreClient := core.NewClient(token, retry)
	s.Cdb = cdb.NewCdbClient(coreClient)
	s.Community = community.NewCommunityClient(coreClient)
}
