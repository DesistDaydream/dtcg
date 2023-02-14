package models

type WishesAddReqBody struct {
	CardVersionID     string `json:"card_version_id"`
	GameKey           string `json:"game_key"`
	IgnoreCardVersion string `json:"ignore_card_version"`
	Quantity          string `json:"quantity"`
	Remark            string `json:"remark"`
	WishListID        string `json:"wish_list_id"`
}
