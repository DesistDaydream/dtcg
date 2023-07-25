package models

type UpdateTokenPostReqBody struct {
	PushDevice string `json:"push_device"`
	Token      string `json:"token"`
}
