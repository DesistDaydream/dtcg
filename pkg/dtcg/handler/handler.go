package handler

import "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"

var H *Handler

type Handler struct {
	DtcgDBServices *services.Services
}

func NewHandler(token string, retry int) *Handler {
	return &Handler{
		DtcgDBServices: services.NewServices(token, retry),
	}
}
