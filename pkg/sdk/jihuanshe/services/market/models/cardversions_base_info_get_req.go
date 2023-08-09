package models

type CardVersionsBaseInfoGetReqBody struct {
	CardVersionID string `json:"card_version_id"`
	GameKey       string `json:"game_key"`
	GameSubKey    string `json:"game_sub_key"`
}
