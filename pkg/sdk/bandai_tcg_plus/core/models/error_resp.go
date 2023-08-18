package models

type ErrorResp struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Msg   string `json:"msg"`
}
