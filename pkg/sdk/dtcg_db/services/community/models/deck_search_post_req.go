package models

type DeckSearchReqQuery struct {
	Limit string `query:"limit"`
	Page  string `query:"page"`
}

type DeckSearchReqBody struct {
	Tags  []string `json:"tags"`
	Kw    string   `json:"kw"`
	Envir string   `json:"envir"`
}
