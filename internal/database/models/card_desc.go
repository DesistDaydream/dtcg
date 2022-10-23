package models

type CardsDesc struct {
	Count       int64      `json:"count"`
	PageSize    int        `json:"page_size"`
	PageCurrent int        `json:"page_current"`
	PageTotal   int        `json:"page_total"`
	Data        []CardDesc `json:"data"`
}

type CardDesc struct {
	ID             int    `gorm:"primaryKey" json:"id"` // ID
	CardIDFromDB   int    `json:"card_id_from_db"`
	SetID          int    `json:"set_id"`
	SetName        string `json:"set_name"`
	SetPrefix      string `json:"set_prefix"`
	Serial         string `json:"serial"`
	SubSerial      string `json:"sub_serial"`
	JapName        string `json:"jap_name"`
	ScName         string `json:"sc_name"`
	AlternativeArt string `json:"alternative_art"`
	Rarity         string `json:"rarity"`
	Type           string `json:"type"`
	Color          string `json:"color"`
	Level          string `json:"level"`
	Cost           string `json:"cost"`
	Cost1          string `json:"cost_1"`
	EvoCond        string `json:"evo_cond"`
	DP             string `json:"dp"`
	Grade          string `json:"grade"`
	Attribute      string `json:"attribute"`
	Class          string `json:"class"`
	Illustrator    string `json:"illustrator"`
	Effect         string `json:"effect"`
	EvoCoverEffect string `json:"evo_cover_effect"`
	SecurityEffect string `json:"security_effect"`
	IncludeInfo    string `json:"include_info"`
	RaritySC       string `json:"rarity_sc"`
	Image          string `json:"image"`
}
