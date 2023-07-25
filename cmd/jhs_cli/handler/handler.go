package handler

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	ms "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	js "github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services"
	"github.com/sirupsen/logrus"
)

var H *Handler

type Handler struct {
	MoecardServices *ms.Services
	JhsServices     *js.Services
}

func NewHandler(moecardIsLogin bool, jhsUserID, moecardUsername, moecardPwd string) *Handler {
	user, err := database.GetUser("1")
	if err != nil {
		logrus.Fatalf("获取用户信息失败，原因: %v", err)
	}

	return &Handler{
		MoecardServices: ms.NewServices(moecardIsLogin, moecardUsername, moecardPwd, user.MoecardToken, 1),
		JhsServices:     js.NewServices(user.JhsToken),
	}
}
