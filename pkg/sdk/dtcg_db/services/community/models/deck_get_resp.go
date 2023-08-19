package models

type DeckGetResp struct {
	Data    DeckData `json:"data"`
	Message string   `json:"message"`
	Success bool     `json:"success"`
}

type DeckData struct {
	Author    string   `json:"author"`
	CopyCount int64    `json:"copy_count"`
	CreatedAt string   `json:"created_at"`
	DeckCode  string   `json:"deck_code"`
	DeckInfo  DeckInfo `json:"deck_info"`
	Desc      string   `json:"desc"`
	Envir     string   `json:"envir"`
	HID       string   `json:"hid"`
	LikeCount int64    `json:"like_count"`
	Name      string   `json:"name"`
	Scene     string   `json:"scene"`
	Source    string   `json:"source"`
	Stat      Stat     `json:"stat"`
	Status    int64    `json:"status"`
	UpdatedAt string   `json:"updated_at"`
	ViewCount int64    `json:"view_count"`
}
