package models

type DeckSearchReq struct {
	Body  DeckSearchReqBody  `json:"body"`
	Query DeckSearchReqQuery `json:"query"`
}

type DeckSearchReqBody struct {
	Tags  []string `json:"tags"`
	Kw    string   `json:"kw"`
	Envir string   `json:"envir"`
}

type DeckSearchReqQuery struct {
	Limit string `json:"limit"`
	Page  string `json:"page"`
}
