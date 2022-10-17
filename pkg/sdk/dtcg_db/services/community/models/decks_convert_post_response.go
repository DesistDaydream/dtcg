package models

type DecksConvertPostResponse struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Data struct {
	DeckInfo DeckInfo `json:"deck_info"`
}

type DeckInfo struct {
	Eggs []Egg  `json:"eggs"`
	Main []Main `json:"main"`
}

type Egg struct {
	Card   EggCard `json:"card"`
	Number int64   `json:"number"`
}

type EggCard struct {
	Attribute      string        `json:"attribute"`
	CardID         int64         `json:"card_id"`
	CardPack       int64         `json:"card_pack"`
	Class          []string      `json:"class"`
	Color          []string      `json:"color"`
	Cost           string        `json:"cost"`
	Cost1          string        `json:"cost_1"`
	DP             string        `json:"DP"`
	Effect         string        `json:"effect"`
	EvoCond        string        `json:"evo_cond"`
	EvoCoverEffect string        `json:"evo_cover_effect"`
	Grade          string        `json:"grade"`
	Illustrator    string        `json:"illustrator"`
	Images         []PurpleImage `json:"images"`
	IncludeInfo    string        `json:"include_info"`
	JapName        string        `json:"japName"`
	Level          string        `json:"level"`
	Package        PurplePackage `json:"package"`
	Rarity         string        `json:"rarity"`
	RaritySC       string        `json:"rarity$SC"`
	ScName         string        `json:"scName"`
	SecurityEffect string        `json:"security_effect"`
	Serial         string        `json:"serial"`
	SubSerial      string        `json:"sub_serial"`
	Type           string        `json:"type"`
}

type PurpleImage struct {
	CardID    int64  `json:"card_id"`
	ID        int64  `json:"id"`
	ImgPath   string `json:"img_path"`
	ThumbPath string `json:"thumb_path"`
}

type PurplePackage struct {
	Language        string `json:"language"`
	PackID          int64  `json:"pack_id"`
	PackPrefix      string `json:"pack_prefix"`
	PackReleaseDate string `json:"pack_releaseDate"`
}

type Main struct {
	Card   MainCard `json:"card"`
	Number int64    `json:"number"`
}

type MainCard struct {
	Attribute      string        `json:"attribute"`
	CardID         int64         `json:"card_id"`
	CardPack       int64         `json:"card_pack"`
	Class          []string      `json:"class"`
	Color          []string      `json:"color"`
	Cost           string        `json:"cost"`
	Cost1          string        `json:"cost_1"`
	DP             string        `json:"DP"`
	Effect         string        `json:"effect"`
	EvoCond        string        `json:"evo_cond"`
	EvoCoverEffect string        `json:"evo_cover_effect"`
	Grade          string        `json:"grade"`
	Illustrator    string        `json:"illustrator"`
	Images         []FluffyImage `json:"images"`
	IncludeInfo    string        `json:"include_info"`
	JapName        string        `json:"japName"`
	Level          string        `json:"level"`
	Package        FluffyPackage `json:"package"`
	Rarity         string        `json:"rarity"`
	RaritySC       string        `json:"rarity$SC"`
	ScName         string        `json:"scName"`
	SecurityEffect string        `json:"security_effect"`
	Serial         string        `json:"serial"`
	SubSerial      string        `json:"sub_serial"`
	Type           string        `json:"type"`
}

type FluffyImage struct {
	CardID    int64  `json:"card_id"`
	ID        int64  `json:"id"`
	ImgPath   string `json:"img_path"`
	ThumbPath string `json:"thumb_path"`
}

type FluffyPackage struct {
	Language        string `json:"language"`
	PackID          int64  `json:"pack_id"`
	PackPrefix      string `json:"pack_prefix"`
	PackReleaseDate string `json:"pack_releaseDate"`
}
