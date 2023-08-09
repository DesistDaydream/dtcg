package models

type PacksGetReq struct {
	GameKey    string `json:"game_key"`
	GameSubKey string `json:"game_sub_key"`
	Page       string `json:"page"`
}
