package handler

import (
	"os"

	ds "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services"
	js "github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services"
	"github.com/sirupsen/logrus"
)

var H *Handler

type Handler struct {
	DtcgDBServices *ds.Services
	JhsServices    *js.Services
}

func NewHandler() *Handler {
	file, err := os.ReadFile("pkg/sdk/jihuanshe/services/token.txt")
	if err != nil {
		logrus.Fatal(err)
	}
	token := string(file)

	return &Handler{
		DtcgDBServices: ds.NewServices(false, "", "", 1),
		JhsServices:    js.NewServices(token),
	}
}
