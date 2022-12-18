package handler

import (
	ds "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	js "github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services"
)

var H *Handler

type Handler struct {
	DtcgDBServices *ds.Services
	JhsServices    *js.Services
}

func NewHandler(jhsToken string) *Handler {
	return &Handler{
		DtcgDBServices: ds.NewServices(false, "", "", 1),
		JhsServices:    js.NewServices(jhsToken),
	}
}
