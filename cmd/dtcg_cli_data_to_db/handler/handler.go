package handler

import "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"

var H *Handler

type Handler struct {
	DtcgDBServices *services.Services
}

func NewHandler(username, password string, retry int) *Handler {
	return &Handler{
		DtcgDBServices: services.NewServices(true, username, password, retry),
	}
}