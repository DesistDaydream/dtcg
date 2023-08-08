package models

type CardVersionsListReq struct {
	CategoryID       string `json:"categoryId"`
	GameKey          string `json:"game_key"`
	GameSubKey       string `json:"game_sub_key"`
	Page             string `json:"page"`
	Rarity           string `json:"rarity"`
	Sorting          string `json:"sorting"`
	SortingPriceType string `json:"sorting_price_type"`
}
