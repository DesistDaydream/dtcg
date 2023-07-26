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

func NewHandler(username, password string, retry int) *Handler {
	user, err := database.GetUser("1")
	if err != nil {
		logrus.Fatalf("获取用户信息异常，原因: %v", err)
	}

	return &Handler{
		MoecardServices: ms.NewServices(true, username, password, user.MoecardToken, retry),
		JhsServices:     js.NewServices(user.JhsToken),
	}
}
