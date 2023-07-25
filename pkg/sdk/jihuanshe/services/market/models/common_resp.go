package models

type CommonSuccessResp struct {
	Message string `json:"message"`
}

type CommonFailureResp struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Msg   string `json:"msg"`
}
