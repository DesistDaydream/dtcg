package handler

import "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"

var H *Handler

type Handler struct {
	MoecardServices *services.Services
}

func NewHandler(username, password, token string, retry int) *Handler {
	return &Handler{
		MoecardServices: services.NewServices(true, username, password, token, retry),
	}
}
