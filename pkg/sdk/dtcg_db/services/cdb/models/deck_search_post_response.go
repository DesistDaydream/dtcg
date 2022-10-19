package models

type DeckSearchPostResponse struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Data struct {
	Meta  Meta  `json:"meta"`
	Decks Decks `json:"decks"`
}

type Meta struct {
	Tags  []int  `json:"tags"`
	Envir string `json:"envir"`
}

type Decks struct {
	Count int    `json:"count"`
	List  []List `json:"list"`
}

type List struct {
	Hid         string        `json:"hid"`
	Envir       string        `json:"envir"`
	Name        string        `json:"name"`
	Author      string        `json:"author"`
	Source      string        `json:"source"`
	Desc        string        `json:"desc"`
	Stat        Stat          `json:"stat"`
	ViewCount   int           `json:"view_count"`
	CopyCount   int           `json:"copy_count"`
	LikeCount   int           `json:"like_count"`
	Status      int           `json:"status"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
	Scene       string        `json:"scene"`
	PickupCards []PickupCards `json:"pickup_cards"`
}

type Stat struct {
	Main map[string]interface{} `json:"main"` // 主卡组
	Eggs map[string]interface{} `json:"eggs"` // 蛋卡组
}

type Main struct {
	Type  map[string]interface{} `json:"type"`
	Color map[string]interface{} `json:"color"`
	Level map[string]interface{} `json:"level"`
}

type Eggs struct {
	Type  map[string]interface{} `json:"type"`
	Color map[string]interface{} `json:"color"`
	Level map[string]interface{} `json:"level"`
}

type PickupCards struct {
	Serial string   `json:"serial"`
	Images []Images `json:"images"`
}

type Images struct {
	ID        int    `json:"id"`
	CardID    int    `json:"card_id"`
	ImgPath   string `json:"img_path"`
	ThumbPath string `json:"thumb_path"`
}
