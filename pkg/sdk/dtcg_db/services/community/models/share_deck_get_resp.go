package models

type ShareDeckGetResp struct {
	Data    ShareDeckGetRespData `json:"data"`
	Message string               `json:"message"`
	Success bool                 `json:"success"`
}

type ShareDeckGetRespData struct {
	DeckInfo DeckInfo `json:"deck_info"`
}
