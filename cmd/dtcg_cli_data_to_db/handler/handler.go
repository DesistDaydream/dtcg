package handler

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	"github.com/sirupsen/logrus"
)

var H *Handler

type Handler struct {
	DtcgDBServices *services.Services
}

func NewHandler(username, password, token string, retry int) *Handler {
	user, err := database.GetUser("1")
	if err != nil {
		logrus.Fatalf("获取用户信息失败，原因: %v", err)
	}

	return &Handler{
		DtcgDBServices: services.NewServices(true, username, password, user.MoecardToken, retry),
	}
}
