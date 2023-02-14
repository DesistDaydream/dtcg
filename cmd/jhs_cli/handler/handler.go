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

func NewHandler(isLogin bool, jhsToken, dtcgdbUsername, dtcgdbPwd string) *Handler {
	return &Handler{
		DtcgDBServices: ds.NewServices(isLogin, dtcgdbUsername, dtcgdbPwd, 1),
		JhsServices:    js.NewServices(jhsToken),
	}
}
