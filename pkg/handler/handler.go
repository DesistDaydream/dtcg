package handler

import (
	"github.com/DesistDaydream/dtcg/internal/database/models"
	ms "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	js "github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services"
)

var H *Handler

type Handler struct {
	UserID          int
	MoecardServices *ms.Services
	JhsServices     *js.Services
}

func NewHandler(user *models.User, isLoginMoecard bool, retry int) *Handler {
	return &Handler{
		UserID:          user.ID,
		MoecardServices: ms.NewServices(user, isLoginMoecard, retry),
		JhsServices:     js.NewServices(user),
	}
}
