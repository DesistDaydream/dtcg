package models

type WishListCreateReqBody struct {
	GameKey    string `json:"game_key"`
	GameSubKey string `json:"game_sub_key"`
	Name       string `json:"name"`
}
