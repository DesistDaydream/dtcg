package models

type WishListGetResp struct {
	CurrentPage int64             `json:"current_page"`
	Data        []WishListGetData `json:"data"`
	From        int64             `json:"from"`
	LastPage    int64             `json:"last_page"`
	NextPageURL string            `json:"next_page_url"`
	Path        string            `json:"path"`
	PerPage     int64             `json:"per_page"`
	PrevPageURL interface{}       `json:"prev_page_url"`
	To          int64             `json:"to"`
	Total       int64             `json:"total"`
}

type WishListGetData struct {
	AvgPrice          string      `json:"avg_price"`
	CardID            int64       `json:"card_id"`
	CardVersionID     int64       `json:"card_version_id"`
	GameKey           string      `json:"game_key"`
	IgnoreCardVersion int64       `json:"ignore_card_version"`
	ImageURL          string      `json:"image_url"`
	Language          interface{} `json:"language"`
	LanguageText      string      `json:"language_text"`
	MinPrice          string      `json:"min_price"`
	NameCN            string      `json:"name_cn"`
	NameCNCnocg       interface{} `json:"name_cn_cnocg"`
	Number            string      `json:"number"`
	NumberAlias       string      `json:"number_alias"`
	Quantity          int64       `json:"quantity"`
	Rarity            string      `json:"rarity"`
	Remark            string      `json:"remark"`
	UserID            int64       `json:"user_id"`
	WishID            int64       `json:"wish_id"`
	WishPrice         interface{} `json:"wish_price"`
}
