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

func NewHandler(isLoginMoecard bool, userID string, moecardUsername string, moecardPwd string, retry int) *Handler {
	user, err := database.GetUser(userID)
	if err != nil {
		logrus.Fatalf("获取用户信息失败，原因: %v", err)
	}

	return &Handler{
		MoecardServices: ms.NewServices(isLoginMoecard, moecardUsername, moecardPwd, user.MoecardToken, retry),
		JhsServices:     js.NewServices(user.JhsToken),
	}
}
