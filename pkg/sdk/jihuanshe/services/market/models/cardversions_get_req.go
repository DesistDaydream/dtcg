package models

type CardVersionGetReqBody struct {
	GameKey    string `json:"game_key"`
	GameSubKey string `json:"game_sub_key"`
}
