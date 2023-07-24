package models

type LoginPostResp struct {
	Data    LoginPostRespData `json:"data"`
	Message string            `json:"message"`
	Success bool              `json:"success"`
}

type LoginPostRespData struct {
	Token string `json:"token"`
}
