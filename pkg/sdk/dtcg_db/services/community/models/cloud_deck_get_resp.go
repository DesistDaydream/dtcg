package models

type CloudDeckGetResp struct {
	Data    CloudDeckData `json:"data"`
	Message string        `json:"message"`
	Success bool          `json:"success"`
}

type CloudDeckData struct {
	ID        int      `json:"id"`
	Game      string   `json:"game"`
	UserID    int      `json:"user_id"`
	WechatID  int      `json:"wechat_id"`
	PublishID string   `json:"publish_id"`
	Envir     string   `json:"envir"`
	Name      string   `json:"name"`
	Source    string   `json:"source"`
	Desc      string   `json:"desc"`
	Stat      Stat     `json:"stat"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	DeckInfo  DeckInfo `json:"deck_info"`
}
