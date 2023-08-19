package models

type DeckConvertPostResp struct {
	Data    DeckConvertData `json:"data"`
	Message string          `json:"message"`
	Success bool            `json:"success"`
}

type DeckConvertData struct {
	DeckInfo DeckInfo `json:"deck_info"`
}

type DeckInfo struct {
	Main []CardInfo `json:"main"`
	Eggs []CardInfo `json:"eggs"`
}

type CardInfo struct {
	Number int  `json:"number"`
	Cards  Card `json:"card"`
}

type Card struct {
	CardID         int      `json:"card_id"`
	CardPack       int      `json:"card_pack"`
	Serial         string   `json:"serial"`
	SubSerial      string   `json:"sub_serial"`
	JapName        string   `json:"japName"`
	ScName         string   `json:"scName"`
	Rarity         string   `json:"rarity"`
	Type           string   `json:"type"`
	Color          []string `json:"color"`
	Level          string   `json:"level"`
	Cost           string   `json:"cost"`
	Cost1          string   `json:"cost_1"`
	EvoCond        string   `json:"evo_cond"`
	DP             string   `json:"DP"`
	Grade          string   `json:"grade"`
	Attribute      string   `json:"attribute"`
	Class          []string `json:"class"`
	Illustrator    string   `json:"illustrator"`
	Effect         string   `json:"effect"`
	EvoCoverEffect string   `json:"evo_cover_effect"`
	SecurityEffect string   `json:"security_effect"`
	IncludeInfo    string   `json:"include_info"`
	RaritySC       string   `json:"rarity$SC"`
	Package        Package  `json:"package"`
	Images         []Images `json:"images"`
}

type Package struct {
	PackID          int    `json:"pack_id"`
	PackPrefix      string `json:"pack_prefix"`
	PackReleaseDate string `json:"pack_releaseDate"`
	Language        string `json:"language"`
}
