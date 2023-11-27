package models

// 数据库模型。卡牌描述信息
type CardsDesc struct {
	Count     int64      `json:"count"`
	PageSize  int        `json:"page_size"`
	PageNum   int        `json:"page_num"`
	PageTotal int        `json:"page_total"`
	Data      []CardDesc `json:"data"`
}

type CardDesc struct {
	ID             int    `json:"id" gorm:"primaryKey"` // ID
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

// 卡牌描述的查询条件
type CardDescQuery struct {
	CardSet    int64     `json:"card_set"`
	ClassInput bool      `json:"class_input"`
	Color      []string  `json:"color"`
	EvoCond    []EvoCond `json:"evo_cond"`
	Keyword    string    `json:"keyword"`
	Language   string    `json:"language"`
	OrderType  string    `json:"order_type"`
	QField     []string  `json:"qField"` // 通过 Keyword 进行查询的字段
	Rarity     []string  `json:"rarity"`
	Tags       []string  `json:"tags"` // 特征
	TagsLogic  string    `json:"tags__logic"`
	Type       string    `json:"type"`
}

// type QueryCardDesc struct {
// 	CardPack   int64     `json:"card_pack"`
// 	ClassInput bool      `json:"class_input"`
// 	Color      []string  `json:"color"`
// 	EvoCond    []EvoCond `json:"evo_cond"`
// 	Keyword    string    `json:"keyword"`
// 	Language   string    `json:"language"`
// 	OrderType  string    `json:"order_type"`
// 	QField     []string  `json:"qField"`
// 	Rarity     []string  `json:"rarity"`
// 	Tags       []string  `json:"tags"` // 特征
// 	TagsLogic  string    `json:"tags__logic"`
// 	Type       string    `json:"type"`
// }

type EvoCond struct {
	Color string `json:"color,omitempty"`
	Cost  string `json:"cost,omitempty"`
	Level string `json:"level,omitempty"`
}
