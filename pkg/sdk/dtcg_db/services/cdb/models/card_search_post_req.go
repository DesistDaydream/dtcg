package models

type CardSearchReqQuery struct {
	Limit string `form:"limit"`
	Page  string `form:"page"`
}

type CardSearchReqBody struct {
	Keyword    string    `json:"keyword"`
	Language   string    `json:"language"` // 卡牌语言：chs,ja
	ClassInput bool      `json:"class_input"`
	CardPack   int       `json:"card_pack"`
	Type       string    `json:"type"`
	Color      []string  `json:"color"`
	Rarity     []string  `json:"rarity"`
	Tags       []string  `json:"tags"` // 卡牌特征。比如：贯通，抽卡，进击，联展，合体 等等
	TagsLogic  string    `json:"tags__logic"`
	OrderType  string    `json:"order_type"`
	EvoCond    []EvoCond `json:"evo_cond"`
	// 搜索范围。可以是：serial(编号)，scName(中文卡名)，japName(日文卡名)，effect(效果)，evo_cover_effect(进化源效果)，security_effect(安防效果)
	// 若为空则从所有范围搜索
	QField []string `json:"qField"`
	// 是否是异画。有任意字符时表示只筛选异画卡，为空时则表示所有。并没有只获取原画卡的逻辑。
	// 想要只获取原画卡，需要使用 Rarity 字段指定全部非异画的稀有度，且该字段为空。
	IsParallel string `json:"is_parallel"`
}
type EvoCond struct {
}
