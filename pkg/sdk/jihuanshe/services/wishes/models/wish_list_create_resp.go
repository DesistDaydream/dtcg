package models

type WishListCreateResp struct {
	GameKey    string `json:"game_key"`
	GameSubKey string `json:"game_sub_key"`
	Name       string `json:"name"`
	WishListID int64  `json:"wish_list_id"`
}
