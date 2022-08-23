package models

type DeckReq struct {
	Tags  []string `json:"tags"`
	Kw    string   `json:"kw"`
	Envir string   `json:"envir"`
}
