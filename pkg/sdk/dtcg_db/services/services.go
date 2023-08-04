package services

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/community"
	"github.com/sirupsen/logrus"
)

type Services struct {
	Cdb       *cdb.CdbClient
	Community *community.CommunityClient
}

func NewServices(user *models.User, isLogin bool, retry int) *Services {
	var token string
	if isLogin {
		if core.CheckToken(user.MoecardToken) {
			logrus.Infoln("Moecard TOKEN 可用，不用重新获取")
		} else {
			logrus.Warnln("Moecard TOKEN 不可用，开始重新获取")
			token = core.Login(user.ID, user.MoecardUsername, user.MoecardPassword)
		}
	} else {
		token = ""
	}

	s := new(Services)
	s.init(user.ID, token, retry)
	return s
}

func (s *Services) init(userID int, token string, retry int) {
	coreClient := core.NewClient(userID, token, retry)
	s.Cdb = cdb.NewCdbClient(coreClient)
	s.Community = community.NewCommunityClient(coreClient)
}
