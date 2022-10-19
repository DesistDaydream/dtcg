package models

type CardSearchReqQuery struct {
	Limit string `query:"limit"`
	Page  string `query:"page"`
}

type CardSearchReqBody struct {
	Keyword    string        `json:"keyword"`
	Language   string        `json:"language"`
	ClassInput bool          `json:"class_input"`
	CardPack   int           `json:"card_pack"`
	Type       string        `json:"type"`
	Color      []interface{} `json:"color"`
	Rarity     []interface{} `json:"rarity"`
	Tags       []interface{} `json:"tags"`
	TagsLogic  string        `json:"tags__logic"`
	OrderType  string        `json:"order_type"`
	EvoCond    []EvoCond     `json:"evo_cond"`
	QField     []interface{} `json:"qField"`
}

type EvoCond struct {
}
