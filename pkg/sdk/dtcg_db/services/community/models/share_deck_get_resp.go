package models

type ShareDeckGetResp struct {
	Data    ShareDeckData `json:"data"`
	Message string        `json:"message"`
	Success bool          `json:"success"`
}

type ShareDeckData struct {
	DeckInfo DeckInfo `json:"deck_info"`
}
